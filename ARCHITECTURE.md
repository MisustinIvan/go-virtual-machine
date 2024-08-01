# 16 bit virtual machine
## memory
 - size: 64 kilobytes (2^16 bits)

## registers
 - size: 16 bits to hold a memory address
 - `bp`: base pointer
 - `sp`: stack pointer
 - `pc`: program counter
 - `ra`: general purpose (access to higher and lower parts)
 - `rb`: general purpose (access to higher and lower parts)
 - `rc`: general purpose (access to higher and lower parts)
 - `rd`: general purpose (access to higher and lower parts)
 - `rs`: special register used as result of operations (cmp), overflow or division by zero. not directly accessible

## instructions
 - `add`: reg0(uint8) reg1(uint8) - add reg0 and reg1 to reg0
 - `sub`: reg0(uint8) reg1(uint8) - subtract reg1 from reg0 to reg0
 - `mul`: reg0(uint8) reg1(uint8) - multiply reg0 with reg1 to reg0
 - `div`: reg0(uint8) reg1(uint8) - divide reg0 by reg1 to reg0
 - `mov`: reg0(uint8) reg1(uint8) - move reg1 to reg0
 - `jmp`: addr0(uint16) - jump to a constant addr0
 - `inc`: reg0(uint8) - increment value in reg0
 - `dec`: reg0(uint8) - decrement value in reg0
 - `lod`: reg0(uint8) addr0(uint16) - load value from addr0 into reg0
 - `str`: addr0(uint16) reg0(uint8) - store value from reg0 into addr0
 - `hlt`: halt
 - `cmp`: reg0(uint8) reg1(uint8) - compare values in reg0 and reg1, result is stored in rs
 - `jio`: addr0(uint16) - jump if one
 - `jiz`: addr0(uint16) - jump if zero
 - `jie`: addr0(uint16) - jump if equals
 - `jne`: addr0(uint16) - jump if not equals
 - `jig`: addr0(uint16) - jump if greater
 - `jis`: addr0(uint16) - jump if smaller
 - `jeg`: addr0(uint16) - jump if greater or equal
 - `jes`: addr0(uint16) - jump if smaller or equal
 - `xor`: reg0(uint8) reg1(uint8) - xor reg0 with reg1 to reg0
