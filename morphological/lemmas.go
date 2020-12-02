package morphological

import (
	"fmt"
	"github.com/eakarpov/msaot/lexicon"
)

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

type GrammarConfig struct {
	POS string
	Gender *string
	Type string
}

type FlexyConfig struct {
	POS string
	Type string
	Case string
	Number string
	Person string
}

func Build(normal string, config GrammarConfig) *Lemma {
	return nil
}

func Normalize(word string, config FlexyConfig) (*Lemma, error) {
	flexies, err := lexicon.GetFlexForm(word)
	if err != nil {
		return nil, err
	}
	fmt.Println(flexies)
	return nil, nil
}
