# Specs

From [pandocs](https://gbdev.io/pandocs/Specifications.html).

<style>
td {
    text-align: center;
}
td:first-child {
    text-align: left;
}
</style>

<table>
  <thead>
    <tr>
      <th></th><th>Game Boy (DMG)</th><th>Game Boy Pocket (MGB)</th><th>Super Game Boy (SGB)</th><th>Game Boy Color (CGB)</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td>CPU</td><td colspan="4">8-bit 8080-like Sharp CPU (speculated to be a SM83 core)</td>
    </tr>
    <tr>
      <td>CPU freq</td><td colspan="2">4.194304&nbsp;MHz<sup class="footnote-reference"><a href="#dmg_clk">1</a></sup></td><td>Depends on revision<sup class="footnote-reference"><a href="#sgb_clk">2</a></sup></td><td>Up to 8.388608&nbsp;MHz</td>
    </tr>
    <tr>
        <td>Work RAM</td><td colspan="3">8&nbsp;KiB</td><td>32&nbsp;KiB<sup class="footnote-reference"><a href="#compat">3</a></sup> (4&nbsp;+&nbsp;7&nbsp;×&nbsp;4&nbsp;KiB)</td>
    </tr>
    <tr>
        <td>Video RAM</td><td colspan="3">8&nbsp;KiB</td><td>16&nbsp;KiB<sup class="footnote-reference"><a href="#compat">3</a></sup> (2&nbsp;×&nbsp;8&nbsp;KiB)</td>
    </tr>
    <tr>
        <td>Screen</td><td>LCD 4.7&nbsp;×&nbsp;4.3&nbsp;cm</td><td>LCD 4.8&nbsp;×&nbsp;4.4&nbsp;cm</td><td>CRT TV</td><td>TFT 4.4&nbsp;×&nbsp;4&nbsp;cm</td>
    </tr>
    <tr>
        <td>Resolution</td><td colspan="2">160&nbsp;×&nbsp;144</td><td>160&nbsp;×&nbsp;144 within 256&nbsp;×&nbsp;224 border</td><td>160&nbsp;×&nbsp;144</td>
    </tr>
    <tr>
        <td>OBJ ("sprites")</td><td colspan="4">8&nbsp;×&nbsp;8 or 8&nbsp;×&nbsp;16 ; max 40 per screen, 10 per line</td>
    </tr>
    <tr>
        <td>Palettes</td><td colspan="2">BG: 1&nbsp;×&nbsp;4, OBJ: 2&nbsp;×&nbsp;3</td><td>BG/OBJ: 1&nbsp;+&nbsp;4&nbsp;×&nbsp;3, border: 4&nbsp;×&nbsp;15</td><td>BG: 8&nbsp;×&nbsp;4, OBJ: 8&nbsp;×&nbsp;3<sup class="footnote-reference"><a href="#compat">3</a></sup></td>
    </tr>
    <tr>
        <td>Colors</td><td>4 shades of green</td><td>4 shades of gray</td><td colspan="2">32768 colors (15-bit RGB)</td>
    </tr>
    <tr>
        <td>Horizontal sync</td><td colspan="2">9.198&nbsp;KHz</td><td>Complicated<sup class="footnote-reference"><a href="#sgb_vid">4</a></sup></td><td>9.198&nbsp;KHz</td>
    </tr>
    <tr>
        <td>Vertical sync</td><td colspan="2">59.73&nbsp;Hz</td><td>Complicated<sup class="footnote-reference"><a href="#sgb_vid">4</a></sup></td><td>59.73&nbsp;Hz</td>
    </tr>
    <tr>
        <td>Sound</td><td colspan="2">4 channels with stereo output</td><td>4 GB channels + SNES audio</td><td>4 channels with stereo output</td>
    </tr>
    <tr>
        <td>Power</td><td>DC 6V, 0.7&nbsp;W</td><td>DC 3V, 0.7&nbsp;W</td><td>Powered by SNES</td><td>DC 3V, 0.6&nbsp;W</td>
    </tr>
  </tbody>
</table>

# CPU

DMG-CPU, which is the CPU in the Gameboy, is a modified version of the ZiLog Z80.

- 4.1Mhz.
- 8 bit processor
- 16 bit address space
  - It has 64Kb of memory space, but the data in it is 8 bit (nums from 0 - 255)

## Model

The CPU is a Sharp SM83. You can find the detailed model in gekkio's reference manual, but it's defined by:
- Interacts with the rest of the SoC using **interrupts**
- 8 bit bidirectional data bus
- 16 bit address bus; single direction, controlled by the CPU
- Control Unit (CU), which **decodes** the executed instructions and checks + dispatches interrupts
- Register file; (NOTE: Register chapter references this)
  - Program Counter (PC) 16 bit
  - Stack Pointer (SP) 16 bit
  - Accumulator (A) 8 bit
  - Flags (F) 8 bit
  - BC, DE, HL: general purpose register pairs of two 8 bit halves (QUESTION: Does this mean that it's a single 16 bit which is then split?)
  - Instruction Register (IR) 8 bit; special purpose
  - Interrupt Enable (IE) 8 bit; special purpose
- 8 bit Arithmentic Logic Unit (ALU); two 8 bit input ports; outputs to the CPU data bus or the registers file
- 16 bit Increment/Decrement Unit (IDU); responsible for incrementing/decrementing 16 bit address bus value; outputs to the register file

## Memory Mapped I/O

Although the CPU (? not sure if it's just in the CPU or in total on the GameBoy) has 16 bit address space, meaning we have 65536 memory positions, not all of it is actual memory.

So the CPU handles I/O through memory addresses, from the various components, like the APU or PPU. This means that some addresses are not meant for either the 8Kb of work RAM or the 8Kb of video RAM, but to these various components, which total to the 64Kb mentioned above.

This also means that the CPU doesn't need extra instructions to talk to the other devices! (pretty cool!)

The whole span of memory is from 0 -> #FFFF. Breaking this into what each section/bank does we have:
- 0 -> 7FFF (32Kb): Returns whatever data the cartridge has stored on **its** memory chip
- 8000 -> 9FFF (8Kb): Graphics RAM, which holds tiles and data that tells the GameBoy how to compose the tiles into a background layer for the game
- A000 -> BFFF (8Kb): External RAM. In practice this is stored in the cartridge and holds things like the game save
- C000 -> DFFF (8Kb): Work RAM.
- E000 -> FFFF (8Kb): Contains an identical copy of the work RAM. And it applies to reads and writes, meaning that if we write something to the address E000 it'll appear in C000 as well, and the opposite as well.
  - This memory range also has 512 bytes that have another function.
    - NOTE: I haven't registered what each subsection does

## Registers

It has 10 registers:

[A|F]
[B|C] [SP]
[D|E] [PC]
[H|L]

> What are registers?
> <br>
> Super fast memory cells located near the CPU. 
> In my own terms, it's something in the CPU which is small and fast to access.

All of those registers are 8 bit, minus **SP** and **PC** which are a special case, using 16 bits.

- [A]ccumulator:
  - Most mathematical operations can be performed on this register 

## Instruction set

**TODO**:
Jot down all teh instruction set, to have a quick ref guide.

- HALT
  - Pauses CPU until a new screen refresh is refreshed or other timers wake it up
  - This is good to save battery life

# Memory

## Work RAM

As we can't do much with just 10 registers, the CPU is connected to this 8 bit 8Kb storage bank.

But this isn't as fast as the registers, and so we're still required to move data into the registers to be able to use it (factually correct?).

## Video RAM

# PPU (Picture Processing Unit)

# APU (Audio Processing Unit)