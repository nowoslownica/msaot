package morphological

import (
	"github.com/eakarpov/msaot/lexicon"
)

type NCase struct {
	Case   lexicon.Case
	Number int
	Person *int
}

type VCase struct {
	Person string
	Number string
	Tense  string
}

type Lemma struct {
	Normal string
	Value  string
	Pos   string
	nCase  *NCase
	vCase  *VCase
}

type GrammarConfig struct {
	POS string
	Gender *string
	Type string
}

type FlexyConfig struct {
	POS lexicon.POS
	Type string
	nConfig NounFlexyConfig
}

type NounFlexyConfig struct {
	Case lexicon.Case
	Number int
	Person string
}

func Build(normal string, config GrammarConfig) *Lemma {
	return nil
}

func Normalize(word string, config FlexyConfig) (*Lemma, error) {
	switch config.POS {
	case lexicon.NOUN:
		return NormalizeNoun(word, config.nConfig)
	default:
		return nil, nil
	}
}

func NormalizeNoun(word string, config NounFlexyConfig) (*Lemma, error) {
	flexies, err := lexicon.GetFlexForm(word)
	if err != nil {
		return nil, err
	}
	var lemma *Lemma
	for _, form := range flexies {
		if form.R.GPositionGrammarPosition != nil {
			position := form.R.GPositionGrammarPosition
			gCase := position.GCase
			gNumber := position.GNumber
			if gNumber.Valid &&
				gCase.Valid &&
				lexicon.Case(gCase.Int64) == config.Case &&
				int(gNumber.Int64) == config.Number {
				var person *int
				if position.GPerson.Valid {
					value := int(position.GPerson.Int64)
					person = &value
				}
				lemma = &Lemma{
					Normal: form.Normal,
					Value:  form.Value,
					Pos:    form.Pos,
					nCase:  &NCase{
						Case:   lexicon.Case(gCase.Int64),
						Number: int(gNumber.Int64),
						Person: person,
					},
				}
			}
		}
	}
	return lemma, nil
}
