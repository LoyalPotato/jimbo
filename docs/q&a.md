# Q&A

Collection of questions and answers for any topic around emulation.

## What does it mean for a CPU to be 8 bit?

I think jdh's video, where he makes his own CPU, might shed some light into what this means.

If you recall, we also have 32-bit and 64-bit architectures (which is what I'm using currently!).

But what this means is that the **word size**.

## What's a word?

It's the size of the register.

But to note, that it can still be bigger than registers... I haven't fully understood and it seems like a complicated matter to get a concise answer, so for simplification, I'll just take it as word = size of a register.

> word is the unit that a machine uses when working with memory

I think this statement above means that we can load from a memory address, at most 8 bits, since that's the size of the register


## (TODO: Go deeper in understanding) What's the computer data bus?

In the same family as the mem. address bus, it's used to carry data from the CPU to other components.

> https://www.bbc.co.uk/bitesize/guides/zhppfcw/revision/2
> https://www.bbc.co.uk/bitesize/guides/zr8kt39/revision/3

## (TODO: Go deeper in understanding) What's a memory address bus and its width (like 16 bit for GB CPU) ?

Carries memory addresses from the processor to other components like storage or i/o devices. 

> https://www.bbc.co.uk/bitesize/guides/zhppfcw/revision/2
> https://www.bbc.co.uk/bitesize/guides/zr8kt39/revision/3 

The width of the address bus is the number of connections that the CPU has to the main memory.

> When you access a memory address in software, you are referring to an address within the addressable memory space defined by the address bus width. For example, in a system with a 16-bit address bus, the address range is from 0x0000 to 0xFFFF (64 KB).
> https://simple.wikipedia.org/wiki/Address_bus

> https://en.wikipedia.org/wiki/Bus_(computing)#Address_bus
> https://superuser.com/questions/1588358/how-can-i-find-my-computers-address-bus-width-and-data-bus-size
> https://www.youtube.com/watch?v=s0Kcg5_z0rI

## What does it mean for a CPU to have 16 bit address bus?

1. It has 2^16 = 65536 KiB of addressable space. This means that the memory addresses can go from 0000 -> FFFF.

## How do I know how many bits a memory address can hold?

> Each address in RAM is usually byte-addressable
> https://forum.allaboutcircuits.com/threads/how-many-bits-can-be-stored-in-each-address-in-a-ram.151977/

## Connection between CPU address bus and max amount of RAM?

> https://cs.stackexchange.com/questions/106276/connection-between-cpu-address-bus-and-max-amount-of-ram

## What's a CPU (instruction) cycle?

> also known as the fetch–decode–execute cycle, or simply the fetch–execute cycle
> It is composed of three main stages: the fetch stage, the decode stage, and the execute stage.

The stages of one cycle

1. Fetch stage: The fetch stage initiates the instruction cycle by retrieving the next instruction from memory. At the end of this stage, the PC points to the next instruction that will be read at the next cycle.

2. Decode stage: During this stage, the encoded instruction in the CIR is interpreted by the control unit (CU). It determines what operations and additional operands are required for execution and sends respective signals to respective components within the CPU, such as the ALU or FPU, to prepare for the execution of the instruction.

3. Execute stage: This is the stage where the actual operation specified by the instruction is carried out by the relevant functional units of the CPU. Logical or arithmetic operations may be run by the ALU, data may be read from or written to memory, and the results are stored in registers or memory as required by the instruction. Based on output from the ALU, the PC might branch.

4. Repeat cycle

In most processors, **interrupts/halts** can occur. This is where the CPU stops execution and remains in an interrup service routine, and when the interrupt stops, it returns to the instruction that it was meant to execute, sometimes interrupting mid-executing and then continue where it left of.

> https://en.wikipedia.org/wiki/Instruction_cycle

### Initiation

This is where the first cycle begins, which occurs as soon as the CPU has power.

The initial PC value is predefined by the system's architecture.

## What's a CPU tick?

It's on the OS, more related to scheduling.

I saw it in a [video](https://youtu.be/e87qKixKFME?si=MPWchEk8D9ht7YC7), being used instead of a cycle, but I assume it was just a misnomer.

> https://stackoverflow.com/questions/43651954/what-is-a-clock-cycle-and-clock-speed

## (TODO: Go deeper in its workings) What's the Program Counter for?

Register that holds the memory address of the **next** instruction to be executed.
After handling the instruction, the PC will be moved to a specific pointer or branch to a specific pointer, which will be handled in the next cycle.

## What's an opcode?

I think it's just a byte that represents what operation/instruction that we need to execute.

> For an x86 CPU, the two bytes B0 61 in the instruction stream cause register AL to be filled with the value 61h.
> B0 is the opcode and 61 is the parameter.
> https://en.wikipedia.org/wiki/Assembly_language#Assembly_language


## How can I reproduce the same cycle tick? For example 4Mhz CPU, but I'm running mine on a Ghz CPU, so do I just need to impose some kind of waiting?

## What're M-cycles?

Machine cycles.

> https://www.reddit.com/r/EmuDev/comments/utyx2g/looking_for_help_understanding_gameboy_clock/

## What're T-cycles?

System clock ticks.

> https://gbdev.io/gb-opcodes/optables/

## What's the difference between machine cycles & system clock ticks? When do I use one or the other?

## In which units are the instruction sets/opcodes?

In the control unit.

## What's an assembler?

Software that turns assembly code into machine/binary code.

## What's an disassembler?

It does the opposite of the assembler; it turns the machine binary code into assembly language.

## What's the CPU clock?

A device which produces an oscillating signal (high/low).

## TODO What's the stack pointer for?

## Do all CPU have a control unit?

## TODO How do 8-bit CPUs handle multi-byte instructions?

> https://stackoverflow.com/questions/19320201/z80-instruction-register-size
> https://www.reddit.com/r/learnprogramming/comments/coynix/how_do_8bit_computers_deal_with_8bit_opcodes_and/

## Is the assumption that the program counter is incremented when doing multi byte fetches correct?

Yes

## In the Z80, when doing relative jumps, it can take 2-3 machine codes depending on the outcome. Why do we need 2 cycles if we can access the flags before fetching the operand (offset)?

It seems like it's just a simple design decision, that stems from the fact that it doesn't take much compute to discard it as it does to use it. This means that we can just fetch it, use it if needed or not if it's not.

> https://retrocomputing.stackexchange.com/questions/22023/why-does-the-z80-jp-absolute-instruction-always-take-10-states-to-execute
> LLM prompt: why does JR cc,n16 in z80 take 3 machine cycles -> why does it need to fetch the operand if we can check the flags before

# Q&A - GB

## How can I find the opcode representations for GB instruction set?

## Do I need to know about GB clocks for emulation?

## Are the pair registers, like BC, 16 bit or just two 8 bit registers joined into one?

They are two 8 bit registers, which can be used as a single 16 bit one.

> https://stackoverflow.com/questions/18687655/is-the-z80-game-boy-cpu-8-or-16-bits

## What's immediate in the context of the instruction set?

## Does the GB always start at BootROM location?

## In the tech ref. for the GB, there's a PHI 1 MiHz. What does it mean?

It's a clock input.

## On the tetris disassembly with mgbdis tool, there's a .map file and there's content before the 0x1000 address. What's its purpose? And when does it run? What's the point of the data before the cartridge header? If the PC starts at 0x0100, then this would never run no?

## How're negative numbers handled? I read in the ref inside this question that it uses two's complement, but how do we know when we want a negative number or not?

> https://stackoverflow.com/questions/3434654/immediate-data-in-assembly

## What's a memory bank controller?

## Does the ALU use some temporary storage for some internal operations?

I had some doubts as to what holds the values of intermediate storage, like in relative jump instructions, that need to add an offset to the PC and if I used the A register, then in the boot rom, it'd override the previously calculated storage.

Looking around I found some references to WZ, but specific to the Z80, so I still had some doubts about its existance in the GB DMG CPU. I finally found some concrete evidence from gekkio's [tweet](https://x.com/gekkio/status/1584236637932126208) (and [here](https://gekkio.fi/blog/2022/8-years-of-game-boy-tweets/)) where he shows the registers in the CPU, WZ included! This makes sense because in his reference, on the JR pseudo code, he makes use of Z & W variable names, but I assumed that it was just a generic temporary value.

> https://gist.github.com/SonoSooS/c0055300670d678b5ae8433e20bea595