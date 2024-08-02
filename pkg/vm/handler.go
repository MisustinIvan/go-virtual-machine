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
	LD8: ld8_op{},
	STR: str_op{},
	ST8: st8_op{},
	HLT: hlt_op{},
	CMP: cmp_op{},
	NOP: nop_op{},
	PRT: prt_op{},
	PR8: pr8_op{},
	JMP: jmp_op{},
	JIF: jif_op{},
	JID: jid_op{},
	JIZ: jiz_op{},
	JIE: jie_op{},
	JNE: jne_op{},
	JIG: jig_op{},
	JIS: jis_op{},
	JEG: jeg_op{},
	JES: jes_op{},
}
