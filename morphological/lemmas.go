package morphological

import (
	"github.com/eakarpov/msaot/lexicon"
	_case "github.com/eakarpov/msaot/lexicon/case"
	"github.com/eakarpov/msaot/lexicon/pos"
)

type NCase struct {
	Case   _case.Case
	Number int
	Person *int
}

type VCase struct {
	Person int
	Number int
	Tense  int
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
	POS pos.POS
	Type string
	nConfig NounFlexyConfig
}

type NounFlexyConfig struct {
	Case _case.Case
	Number int
	Person string
}

func Build(normal string, config GrammarConfig) *Lemma {
	return nil
}

func NormalizeAuto(word string) ([]*Lemma, error) {
	flexies, err := lexicon.GetFlexForm(word)
	if err != nil {
		return nil, err
	}
	res := make([]*Lemma, 0)
	for _, flexy := range flexies {
		var nCase *NCase
		var vCase *VCase
		position := flexy.R.GPositionIdGrammarPosition
		if position.GCase.Valid {
			var person *int
			if position.GPerson.Valid {
				personValue := int(position.GPerson.Int64)
				person = &personValue
			}
			nCase = &NCase{
				Case: _case.Case(position.GCase.Int64),
				Number: int(position.GNumber.Int64),
				Person: person,
			}
		}
		if position.GTense.Valid {
			vCase = &VCase{
				Person: int(position.GPerson.Int64),
				Number: int(position.GNumber.Int64),
				Tense:  int(position.GTense.Int64),
			}
		}
		res = append(res, &Lemma{
			Normal: flexy.R.LemmaIdLemmas.Value,
			Value:  flexy.Value,
			Pos:    flexy.R.LemmaIdLemmas.Pos,
			nCase:  nCase,
			vCase:  vCase,
		})
	}
	return res, nil
}

func Normalize(word string, config FlexyConfig) (*Lemma, error) {
	switch config.POS {
	case pos.NOUN:
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
		if form.R.GPositionIdGrammarPosition != nil && form.R.LemmaIdLemmas != nil {
			position := form.R.GPositionIdGrammarPosition
			gCase := position.GCase
			gNumber := position.GNumber
			if gNumber.Valid &&
				gCase.Valid &&
				_case.Case(gCase.Int64) == config.Case &&
				int(gNumber.Int64) == config.Number {
				var person *int
				if position.GPerson.Valid {
					value := int(position.GPerson.Int64)
					person = &value
				}
				lemmaValue := form.R.LemmaIdLemmas
				lemma = &Lemma{
					Normal: lemmaValue.Value,
					Value:  form.Value,
					Pos:    lemmaValue.Pos,
					nCase:  &NCase{
						Case:   _case.Case(gCase.Int64),
						Number: int(gNumber.Int64),
						Person: person,
					},
				}
			}
		}
	}
	return lemma, nil
}
