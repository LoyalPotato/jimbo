package cpu

/*
QUESTION:
I'm still a bit confused as to what part of the cpu holds what (registers, handles the decoding/choosing, executing, etc)

I did see, that we get the current address from the PC, which then sets the instruction byte into the IR register

QUESTION:
- Which instructions are handled by the ALU? Only arithmetic & jmp (logic?)
*/

import (
	"errors"
)

// TODO
type iControlUnit interface {
	init()
	fetch()
	decode(uint8) interface{}
	ld(dst, src uint8) // TODO: How can I make it receive only from the type of registers?
}

type IR uint8 // Instruction Register
type IE uint8 // Interrupt Enable

type A uint8 // Accumulator
type F uint8 // Flags
type B uint8 // General Purpose
type C uint8 // General Purpose
type D uint8 // General Purpose
type E uint8 // General Purpose
type H uint8 // General Purpose
type L uint8 // General Purpose

type SP uint16 // Stack Pointer
type PC uint16 // Program Counter

type registers struct {
	IR // Instruction Register
	IE // Interrupt Enable

	A // Accumulator
	F // Flags
	B // General Purpose
	C // General Purpose
	D // General Purpose
	E // General Purpose
	H // General Purpose
	L // General Purpose

	SP // Stack Pointer
	PC // Program Counter
}

// Decodes executed instruction
// QUESTION: In the model on the ref. guide, it doesn't have any connection to the register file
// so how does it get the address of the instruction?
type controlUnit struct {
	registers registers
}

func newCU() controlUnit {
	return controlUnit{}
}

func (cu *controlUnit) init() {
	// QUESTION: Does the GB always start at this location?

	// Boot ROM start addr
	cu.registers.PC = 0x0000
	// Init all/some of the registers?
	panic("to implement")
}

func (cu *controlUnit) fetch() {
	// MAIN TODO: Figure out if what's on the PC is the actual instruction
	// or the address from where we need to fetch the instruction in memory.
	// I think it's the latter, since from what I remember seeing, the PC points to an address hmm

	memVal := cu.fetchPCMemValue()
	// TODO: need to have memory to fetch the value in this address?
	// TODO: need to put the instruction in the IR
	cu.decode(memVal)

	// TODO: Check if this is the rigth place for it
	cu.registers.PC++
}

// QUESTION: This will need to be changed when I figure out what size the values can be in memory
// The whole, if it can store addr, then it should be uint16?
func (cu *controlUnit) fetchPCMemValue() uint8 {
	// memAddr := uint16(cu.registers.PC)
	panic("to implement/fix")
	// TODO: I think I need to implement a bus (data and/or addr) to be able to interact
	// with memory from the control unit
	// return emulator.GlobalCtx.Memory.FetchValue(memAddr)
}

func (cu *controlUnit) decode(instr uint8) {
	panic("to implement")
}

func (cu *controlUnit) ld(dst, src uint8) error {
	if ok := checkR8(src); !ok {
		return errors.New("invalid src register being used")
	}

	if ok := checkR8(dst); !ok {
		return errors.New("invalid dst register being used")
	}

	panic("to implement")
}

// TEST: Will this work if I pass any uint8?
// I want it to only return true if the type is one of the r8 registers but I don't want
// interoperability; that is, if I pass a random uint8, will it work?
func checkR8(val interface{}) bool {
	if _, ok := val.(A); ok {
		return ok
	}

	if _, ok := val.(B); ok {
		return ok
	}

	if _, ok := val.(C); ok {
		return ok
	}

	if _, ok := val.(D); ok {
		return ok
	}

	if _, ok := val.(E); ok {
		return ok
	}

	if _, ok := val.(H); ok {
		return ok
	}

	if _, ok := val.(L); ok {
		return ok
	}

	return false
}
