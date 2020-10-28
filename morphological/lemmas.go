package morphological

type NCase struct {
	Case   string
	Number string
	Person string
}

type VCase struct {
	Person string
	Number string
	Tense  string
}

type Lemma struct {
	Normal string
	Value  string
	Post   string
	nCase  *NCase
	vCase  *VCase
}

func Build(word string) *Lemma {
	return nil
}
