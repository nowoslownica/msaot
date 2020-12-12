package morphological

import (
	_case "github.com/eakarpov/msaot/lexicon/case"
	"github.com/eakarpov/msaot/lexicon/dictionary"
	"github.com/eakarpov/msaot/lexicon/gender"
	"github.com/eakarpov/msaot/lexicon/number"
	"github.com/eakarpov/msaot/lexicon/pos"
	"github.com/eakarpov/msaot/lexicon/tense"
)

func Build(normal string, config GrammarConfig) *Lemma {
	return nil
}

func NormalizeAuto(word string) ([]*Lemma, error) {
	flexies, err := dictionary.GetFlexForm(word)
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
				Case:   _case.Case(position.GCase.Int64),
				Number: number.Number(position.GNumber.Int64),
				Person: person,
				Gender: gender.Gender(position.GGender.Int64), // check valid?
			}
		}
		if position.GTense.Valid {
			vCase = &VCase{
				Person: int(position.GPerson.Int64),
				Number: number.Number(position.GNumber.Int64),
				Tense:  tense.Tense(position.GTense.Int64),
				Gender: gender.Gender(position.GGender.Int64), // check valid?
			}
		}
		res = append(res, &Lemma{
			Normal: flexy.R.LemmaIdLemmas.Value,
			Value:  flexy.Value,
			Pos:    pos.POS(flexy.R.LemmaIdLemmas.Pos),
			NCase:  nCase,
			VCase:  vCase,
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
	flexies, err := dictionary.GetFlexForm(word)
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
					Pos:    pos.POS(lemmaValue.Pos),
					NCase: &NCase{
						Case:   _case.Case(gCase.Int64),
						Number: number.Number(gNumber.Int64),
						Person: person,
					},
				}
			}
		}
	}
	return lemma, nil
}
