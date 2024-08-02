package vm

const NREGS = 8

type register struct {
	val  uint16
	kind reg
}

type reg uint8

const (
	PC reg = iota
	BP
	SP
	RA
	RB
	RC
	RD
	RS
)

func regtostring(r reg) string {
	switch r {
	case PC:
		return "PC"
	case BP:
		return "BP"
	case SP:
		return "SP"
	case RA:
		return "RA"
	case RB:
		return "RB"
	case RC:
		return "RC"
	case RD:
		return "RD"
	case RS:
		return "RS"
	default:
		return "WHAT THE FUCK ARE YOU DOING DUDE"
	}
}
