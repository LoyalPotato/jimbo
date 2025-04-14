package cpu

/*
TODO
- Implement the cycle &| tick for the CPU
*/
type ICPU interface {
	Init()
	Boot()
	execute()
}

type CPU struct {
	cu  controlUnit
	alu arithmeticLogicUnit
	idu incdecUnit
}

// NOTE: Do I need this?
// Increment/Decrement
type incdecUnit struct{}

func (cpu *CPU) Init() {
	cpu.cu.init()
}

func (cpu *CPU) Boot() {

}

/*
NOTE:

Uses fetch/execute overlap when possible to do opcode fetches & execute instructions.

*/
