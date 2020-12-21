package number

type Number int

const (
	_ Number = iota
	SINGULAR
	DUAL
	PLURAL
)

func ToString(num Number) string {
	switch num {
	case SINGULAR:
		return "SG"
	case DUAL:
		return "DU"
	case PLURAL:
		return "PL"
	default:
		return "unknown"
	}
}