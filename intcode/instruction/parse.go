package instruction

// Parse takes an incode instruction and returns mode3, mode2, mode1, and the opcode
func Parse(op int) (int, int, int, int) {
	return op / 10000, (op / 1000) % 10, (op / 100) % 10, op % 100
}
