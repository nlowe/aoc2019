package day23

import (
	"fmt"
	"reflect"
)

const (
	addressAnswer    = 255
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

	bufs []chan packet
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
	}
}

func (r *router) RouteTraffic() <-chan int {
	answer := make(chan int)
	go r.runPacketAggregator(answer)
	go r.runTransmitter()

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

		if dst == addressAnswer {
			answer <- y
			return
		}

		fmt.Printf("[%2d] send {%d,%d} to %d\n", sender, x, y, dst)
		r.bufs[dst] <- packet{sender, dst, x, y}
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
			fmt.Printf("[%2d] got {%d,%d} from %d\n", id, p.x, p.y, p.from)
			in <- p.x
			in <- p.y
		}
	}
}
