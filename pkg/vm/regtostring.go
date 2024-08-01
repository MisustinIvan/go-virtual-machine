package vm

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
