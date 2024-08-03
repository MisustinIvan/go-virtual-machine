package vm

type operation interface {
	do(v *VM) bool
	size() uint8
}

func op_ok(o operation, v VM) bool {
	// lets test -1, makes sense to me
	return v.get_pc()+uint16(o.size())-1 < MEMSIZE-1
}
