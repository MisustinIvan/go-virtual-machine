# 16 bit virtual machine
## memory
 - size: 64 kilobytes (2^16 bits)
 - little endian

## registers
 - size: 16 bits to hold a memory address
 - `pc`: program counter
 - `bp`: base pointer
 - `sp`: stack pointer
 - `ra`: general purpose (access to higher and lower parts)
 - `rb`: general purpose (access to higher and lower parts)
 - `rc`: general purpose (access to higher and lower parts)
 - `rd`: general purpose (access to higher and lower parts)
 - `rs`: special register used as result of operations (cmp), overflow or division by zero. not directly accessible

## instructions
### special
  - special instructions are not real in the cpu, but are made up by the assembler
  - [x] `cal`: addr(uint16) - call addr
    - pushes all modified registers from stack frame onto the stack, including the base pointer and jumps to specified address
  - [x] `ret`: - return
    - pops the base pointer and then pops all of the registers pushed previously to restore previous stack frame

 - if an instruction was to generate an error(divzero, overflow), the operation is left to proceed and the program is informed via the special register. in case of division by zero, the result is zero
 - [x] `nop`: no operation
 - [x] `add`: reg0(uint8) reg1(uint8) - add reg0 and reg1 to reg0
 - [x] `adi`: reg0(uint8) val0(uint16) - add reg0 and val0 to reg0
 - [x] `sub`: reg0(uint8) reg1(uint8) - subtract reg1 from reg0 to reg0
 - [x] `sbi`: reg0(uint8) val0(uint16) - subtract val0 from reg0 to reg0
 - [x] `mul`: reg0(uint8) reg1(uint8) - multiply reg0 with reg1 to reg0
 - [x] `mli`: reg0(uint8) val0(uint16) - multiply reg0 with val0 to reg0
 - [x] `div`: reg0(uint8) reg1(uint8) - divide reg0 by reg1 to reg0
 - [x] `dvi`: reg0(uint8) val0(uint16) - divide reg0 by val0 to reg0
 - [x] `mov`: reg0(uint8) reg1(uint8) - move reg1 to reg0
 - [x] `mvi`: reg0(uint8) val0(uint16) - move val0 to reg0
 - [x] `pus`: reg0(uint8) - push reg0
 - [x] `pop`: reg0(uint8) - pop reg0
 - [x] `jmp`: addr0(uint16) - jump to a constant addr0
 - [x] `jpi`: reg0(uint8) - jump to an address in reg0
 - [x] `inc`: reg0(uint8) - increment value in reg0
 - [x] `dec`: reg0(uint8) - decrement value in reg0
 - [x] `lod`: reg0(uint8) addr0(uint16) - load value from addr0 into reg0
 - [x] `ld8`: reg0(uint8) addr0(uint16) - load 8 bits from addr0 into lower part of reg0
 - [x] `str`: addr0(uint16) reg0(uint8) - store value from reg0 into addr0
 - [x] `st8`: addr0(uint16) reg0(uint8) - store lower 8 bits of reg0 into addr0
 - [x] `prt`: reg0(uint8) - print reg0 as number(whole register)
 - [x] `pr8`: reg0(uint8) - print reg0 as char(lower 8 bits)
 - [x] `hlt`: halt
 - [x] `cmp`: reg0(uint8) reg1(uint8) - compare values in reg0 and reg1, result is stored in rs
 - [x] `jif`: addr0(uint16) - jump if overflow
 - [x] `jid`: addr0(uint16) - jump if division by zero
 - [x] `jiz`: addr0(uint16) - jump if zero
 - [x] `jie`: addr0(uint16) - jump if equals
 - [x] `jne`: addr0(uint16) - jump if not equals
 - [x] `jig`: addr0(uint16) - jump if greater
 - [x] `jis`: addr0(uint16) - jump if smaller
 - [x] `jeg`: addr0(uint16) - jump if greater or equal
 - [x] `jes`: addr0(uint16) - jump if smaller or equal
 - [ ] `and`: reg0(uint8) reg1(uint8) ands reg0 and reg1 to reg0
 - [ ] `not`: reg0(uint8) reg1(uint8) nots reg0 and reg1 to reg0
 - [ ] `or`:  reg0(uint8) reg1(uint8) ors reg0 and reg1 to reg0
 - [ ] `xor`: reg0(uint8) reg1(uint8) - xor reg0 with reg1 to reg0
    - do i really need xor?
