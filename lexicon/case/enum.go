package _case

type Case int

const (
	_ Case = iota
	NOMINATIVE
	GENITIVE
	PARTITIVE
	ACCUSATIVE
	DATIVE
	INSTRUMENTAL
	PREPOSITIONAL
	LOCATIVE
	VOCATIVE
)

func ToString(caseVal Case) string {
	switch caseVal {
	case NOMINATIVE:
		return "NOM"
	case GENITIVE:
		return "GEN"
	case DATIVE:
		return "DAT"
	case ACCUSATIVE:
		return "ACC"
	case INSTRUMENTAL:
		return "INS"
	case PARTITIVE:
		return "PAR"
	case PREPOSITIONAL:
		return "PRE"
	case LOCATIVE:
		return "LOC"
	case VOCATIVE:
		return "VOC"
	default:
		return "unknown"
	}
}

