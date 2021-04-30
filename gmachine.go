// Package gmachine implements a simple virtual CPU, known as the G-machine.
package gmachine

// DefaultMemSize is the number of 64-bit words of memory which will be
// allocated to a new G-machine by default.
const (
	DefaultMemSize        = 1024
	OpHALT         uint64 = 0
	OpNOOP         uint64 = 1
	OpINCA         uint64 = 2
	OpDECA         uint64 = 3
	OpSETA         uint64 = 4
)

type gMachine struct {
	P      uint64
	A      uint64
	Memory []uint64
}

func New() *gMachine {
	return &gMachine{P: 0, A: 0, Memory: make([]uint64, DefaultMemSize)}
}

func (g *gMachine) Run() {
	for {
		op := g.Memory[g.P]
		switch op {
		case OpHALT:
			g.P += 1
			return
		case OpNOOP:
			g.P += 1
		case OpINCA:
			g.A += 1
			g.P += 1
		case OpDECA:
			g.A -= 1
			g.P += 1
		case OpSETA:
			g.P += 1
			g.A = g.Memory[g.P]
			g.P += 1
		}
	}
}

func (g *gMachine) RunProgram(instructions []uint64) {
	for index, instruction := range instructions {
		g.Memory[index] = instruction
	}
	g.Run()
}
