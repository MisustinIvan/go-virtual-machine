package vm

func (v *VM) readu16(address uint16) (uint16, bool) {
	if address > MEMSIZE-2 {
		return 0, false
	}

	var lower uint8 = v.Memory[address]
	var higher uint8 = v.Memory[address+1]

	return uint16(lower) | (uint16(higher) << 8), true
}

func (v *VM) readu8(address uint16) (uint16, bool) {
	if address > MEMSIZE-1 {
		return 0, false
	}

	var lower uint8 = v.Memory[address]

	return uint16(lower), true
}

func (v *VM) writeu16(address uint16, value uint16) bool {
	if address > MEMSIZE-2 {
		return false
	}

	var lower uint8 = uint8(value >> 8)
	var higher uint8 = uint8(value)
	v.Memory[address] = lower
	v.Memory[address+1] = higher
	return true
}

func (v *VM) writeu8(address uint16, value uint16) bool {
	if address > MEMSIZE-1 {
		return false
	}

	var lower uint8 = uint8(value)

	v.Memory[address] = lower

	return true
}
