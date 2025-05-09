# Resources

## CPU

- https://en.wikipedia.org/wiki/CPU
- https://en.wikipedia.org/wiki/opcodes
- https://en.wikipedia.org/wiki/Opcode_table
- https://en.wikipedia.org/wiki/Instruction_set_architecture#Instructions
- https://en.wikipedia.org/wiki/Machine_code
- https://www.youtube.com/watch?v=ZyPewh04OzQ
- https://www.youtube.com/watch?v=cNN_tTXABUA
- https://www.youtube.com/watch?v=8XmxKPJDGU0&list=PLrOv9FMX8xJHqMvSGB_9G9nZZ_4IgteYf&index=3
- https://www.youtube.com/watch?v=LnzuMJLZRdU&list=PLowKtXNTBypFbtuVMUVXNR0z1mu7dp7eH&index=1
- https://www.youtube.com/watch?v=Z5JC9Ve1sfI
- https://www.youtube.com/watch?v=ctjnYgCo8Bc
- https://www.youtube.com/watch?v=UrLHufXCRGc&t=220s
- https://www.reddit.com/r/EmuDev/comments/h01yvg/best_way_to_handle_opcodes_of_processor/
- https://www.reddit.com/r/EmuDev/comments/gmslx7/how_do_you_prefer_to_handle_dispatch/
- https://www.youtube.com/watch?v=B7djs4zSbuU
- https://www.quora.com/How-can-overflow-be-detected-while-adding-unsigned-or-twos-complement-binary-number
- https://teaching.idallen.com/dat2343/10f/notes/040_overflow.txt
- https://youtu.be/1I5ZMmrOfnA?si=W3A47jlV_LFaar-w

## GB

- https://www.youtube.com/watch?v=3yXJYvavj6E
- https://www.youtube.com/watch?v=7afPrCee5Ek
- https://gb-archive.github.io/salvage/decoding_gbz80_opcodes/Decoding%20Gamboy%20Z80%20Opcodes.html - Decoding Z80's opcodes; not used in this emulator just to keep things simpler for now
- https://izik1.github.io/gbops/
- https://floooh.github.io/2021/12/06/z80-instruction-timing.html
- https://gbdev.gg8.se/wiki/articles/Gameboy_Bootstrap_ROM
- https://github.com/ISSOtm/gb-bootroms/
- https://gbdev.io/pandocs
- https://github.com/Baekalfen/PyBoy/blob/master/extras/PyBoy.pdf
- https://github.com/Gekkio/gb-ctr
- https://www.youtube.com/watch?v=TRja-uIWq0c
- http://gameboy.mongenel.com/dmg/asmmemmap.html
- https://www.youtube.com/watch?v=s0Kcg5_z0rI
- https://stackoverflow.com/questions/3434654/immediate-data-in-assembly
- https://rylev.github.io/DMG-01/public/book/introduction.html
- https://realboyemulator.wordpress.com/2013/01/03/a-look-at-the-game-boy-bootstrap-let-the-fun-begin/ - Really in depth guide of the boot ROM instructions
- https://jnz.dk/z80/opref.html
- https://stackoverflow.com/questions/8034566/overflow-and-carry-flags-on-z80
- https://stackoverflow.com/questions/3434654/immediate-data-in-assembly
- https://retrocomputing.stackexchange.com/questions/1626/why-did-pok%C3%A9mon-red-have-so-many-overflow-bugs

## Generic Debug Tools

- https://mh-nexus.de/en/hxd/ - Hex Editor

## Binary

- https://www.youtube.com/watch?v=4qH4unVtJkE
- https://www.youtube.com/watch?v=ydboHy_yNts
- https://courses.cs.washington.edu/courses/cse390b/21sp/readings/negative_binary.html

## Go

- https://google.github.io/styleguide/go/decisions#receiver-type
- https://go.dev/blog/slices-intro
 
## Compilers/Interpreters

I'm leaving this here because I thought of an idea that we might be able to write games in other languages, like Go, and then compile it to assembly?

Wait that doesn't make sense, as Go is already compile to assembly. But! it isn't compile to the GameBoy CPU, so not sure if there's something that could be done there or if I got it all wrong in my understanding of this.

The idea came from jdh's [video](https://www.youtube.com/watch?v=7A1SzIIKMho)

- https://norasandler.com/2017/11/29/Write-a-Compiler.html
- https://interpreterbook.com/
- https://www.youtube.com/watch?v=7A1SzIIKMho