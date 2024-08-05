_start:
	MVI RA 0x0
	MVI RB 0x1
	MVI RC 0x14
	PRT RA

loop:
	MVI RD 0x0
	CMP RC RD
	JIE exit
	DEC RC
	CAL fib
	PRT RA
	JMP loop

exit:
	HLT

fib:
	PUS RA
	ADD RA RB
	POP RB
	RET
