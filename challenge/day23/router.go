package day23

import (
	"context"
	"fmt"
	"reflect"
	"runtime"
	"time"

	"golang.org/x/sync/semaphore"
)

const (
	addressNAT       = 255
	routerBufferSize = 16
)

type packet struct {
	from int
	to   int
	x    int
	y    int
}

type router struct {
	ins  []chan int
	outs []<-chan int

	bufs        []chan packet
	idleTracker *semaphore.Weighted

	trackNat      bool
	lastNATPacket packet
}

func NewRouter(in []chan int, out []<-chan int) *router {
	bufs := make([]chan packet, len(in))
	for i := range in {
		bufs[i] = make(chan packet, routerBufferSize)
	}

	return &router{
		ins:  in,
		outs: out,
		bufs: bufs,

		idleTracker: semaphore.NewWeighted(int64(len(in))),
	}
}

func (r *router) RouteTraffic() <-chan int {
	answer := make(chan int)
	go r.runPacketAggregator(answer)
	go r.runTransmitter()
	if r.trackNat {
		go r.runNATHandler(answer)
	}

	return answer
}

func (r *router) runPacketAggregator(answer chan int) {
	readers := make([]reflect.SelectCase, len(r.outs))
	for i, ch := range r.outs {
		readers[i] = reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(ch)}
	}

	for {
		sender, value, ok := reflect.Select(readers)
		if !ok {
			return
		}

		dst := int(value.Int())
		x := <-r.outs[sender]
		y := <-r.outs[sender]

		parsed := packet{sender, dst, x, y}

		if dst == addressNAT {
			if !r.trackNat {
				fmt.Printf("[NAT] write answer (non-tracking): %d\n", y)
				answer <- y
				return
			}

			r.lastNATPacket = parsed
			fmt.Printf("[NAT] intercept {%d,%d} from %d\n", x, y, sender)
			continue
		}

		fmt.Printf("[%3d] send {%d,%d} to %d\n", sender, x, y, dst)
		r.bufs[dst] <- parsed
	}
}

func (r *router) runTransmitter() {
	for i := range r.ins {
		id := i
		go r.runTransmitterFor(id)
	}
}

func (r *router) runTransmitterFor(id int) {
	packets := r.bufs[id]
	in := r.ins[id]

	for {
		select {
		case in <- -1:
		case p := <-packets:
			_ = r.idleTracker.Acquire(context.Background(), 1)
			fmt.Printf("[%3d] got {%d,%d} from %d\n", id, p.x, p.y, p.from)
			in <- p.x
			in <- p.y
			r.idleTracker.Release(1)
		}

		// Give other goroutines some time to run
		runtime.Gosched()
	}
}

func (r *router) runNATHandler(answer chan int) {
	last := -1
	for {
		// Give the adapters some time to send some traffic, especially if we just
		// kick-started node 0
		time.Sleep(50 * time.Millisecond)

		if r.lastNATPacket.to != addressNAT || !r.idleTracker.TryAcquire(int64(len(r.ins))) {
			// Yield back to the scheduler
			runtime.Gosched()
			continue
		}

		if r.lastNATPacket.y == last {
			fmt.Printf("[NAT] write answer (tracking) %d\n", last)
			answer <- r.lastNATPacket.y
			return
		}

		fmt.Printf("[NAT] network idle enough, sending {%d,%d} from %d to [  0]\n", r.lastNATPacket.x, r.lastNATPacket.y, r.lastNATPacket.from)
		r.bufs[0] <- r.lastNATPacket
		last = r.lastNATPacket.y

		// Don't try to immediately steal the semaphore back
		r.idleTracker.Release(int64(len(r.ins)))
		runtime.Gosched()
	}
}
