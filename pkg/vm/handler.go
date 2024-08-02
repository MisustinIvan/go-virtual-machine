package vm

var handler map[opcode]operation = map[opcode]operation{
	ADD: add_op{},
	SUB: sub_op{},
	MUL: mul_op{},
	DIV: div_op{},
	INC: inc_op{},
	DEC: dec_op{},
	MOV: mov_op{},
	LOD: lod_op{},
	STR: str_op{},
	HLT: hlt_op{},
	NOP: nop_op{},
	PRT: prt_op{},
	JMP: jmp_op{},
}
