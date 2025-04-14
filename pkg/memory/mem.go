package memory

import "fmt"

type iMemory interface {
	Init()
	FetchValue(addr uint16) uint8
	// TODO
	SetValue(addr uint16, value uint8) // QUESTION: If I can set other addresses in mem value, shouldn't it be uint16?
}

/*
Memory Map (https://gbdev.io/pandocs/Memory_Map.html)

0000	3FFF	16 KiB ROM bank 00	From cartridge, usually a fixed bank
4000	7FFF	16 KiB ROM Bank 01–NN	From cartridge, switchable bank via mapper (if any)
8000	9FFF	8 KiB Video RAM (VRAM)	In CGB mode, switchable bank 0/1
A000	BFFF	8 KiB External RAM	From cartridge, switchable bank if any
C000	CFFF	4 KiB Work RAM (WRAM)
D000	DFFF	4 KiB Work RAM (WRAM)	In CGB mode, switchable bank 1–7
E000	FDFF	Echo RAM (mirror of C000–DDFF)	Nintendo says use of this area is prohibited.
FE00	FE9F	Object attribute memory (OAM)
FEA0	FEFF	Not Usable	Nintendo says use of this area is prohibited.
FF00	FF7F	I/O Registers
FF80	FFFE	High RAM (HRAM)
FFFF	FFFF	Interrupt Enable register (IE)
*/
// TODO: Add better descriptions for each range one
const (
	MIN_MEM uint16 = 0x0
	MAX_MEM uint16 = 0x1000

	// Bank 00
	MIN_BK0 uint16 = MIN_MEM
	MAX_BK0 uint16 = 0x3FFF

	// Bank 01-NN
	MIN_BK1N uint16 = MIN_MEM
	MAX_BK1N uint16 = 0x7FFF

	// Video RAM
	// TODO: There's something about switchable bank, so need to look into it
	// first, before setting values

	// External RAM
	// TODO: There's something about switchable bank, so need to look into it
	// first, before setting values

	// W(ork)RAM
	MIN_WRAM = 0xC000
	MAX_WRAM = 0xCFFF

	// W(ork)RAM Switchable
	// TODO: There's something about switchable bank, so need to look into it
	// first, before setting values

	// Object attribute memory
	MIN_OAM = 0xFE00
	MAX_OAM = 0xFE9F

	// Forbidden
	MIN_FORB = 0xFEA0
	MAX_FORB = 0xFEFF

	// IO Registers
	// TODO: This can be further subdivided
	MIN_IOREG = 0xFF00
	MAX_IOREG = 0xFF7F

	// H(igh) RAM
	MIN_HRAM = 0xFF80
	MAX_HRAM = 0xFFFE

	// Interruptable Register
	IE = 0xFFFF
)

type Memory struct {
	memArr [MAX_MEM]uint8
}

func (m *Memory) Init() {
	panic("to implement. put values in the correct spot in memory; for example any boot stuff")
}

func (m *Memory) FetchValue(addr uint16) uint8 {
	if addr > MAX_MEM || addr < MIN_MEM {
		panic(fmt.Sprintf("tried to fetch outside of memory range; addr: %X", addr))
	}

	if addr >= MIN_FORB || addr <= MAX_FORB {
		panic(fmt.Sprintf("tried to fetch in forbidden range; addr: %X", addr))

	}

	return m.memArr[addr]
}

// TODO: Should I have functions to access specific parts of the memory?
// I haven't looked into (and implemented) the whole memory banks, so dunno
// if I have to be careful with some ranges

func (m *Memory) SetValue(addr uint16, value uint8) {
	panic("to implement")
}
