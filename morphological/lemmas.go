package morphological

import (
	_case "github.com/eakarpov/msaot/lexicon/case"
	"github.com/eakarpov/msaot/lexicon/dictionary"
	"github.com/eakarpov/msaot/lexicon/gender"
	"github.com/eakarpov/msaot/lexicon/lemmas"
	"github.com/eakarpov/msaot/lexicon/number"
	"github.com/eakarpov/msaot/lexicon/pos"
	"github.com/eakarpov/msaot/lexicon/tense"
)

func Build(normal string, config lemmas.GrammarConfig) *lemmas.Lemma {
	return nil
}

func NormalizeAuto(word string) ([]*lemmas.Lemma, error) {
	flexies, err := dictionary.GetFlexForm(word)
	if err != nil {
		return nil, err
	}
	res := make([]*lemmas.Lemma, 0)
	for _, flexy := range flexies {
		var nCase *lemmas.NCase
		var vCase *lemmas.VCase
		position := flexy.R.GPositionIdGrammarPosition
		if position.GCase.Valid {
			var person *int
			if position.GPerson.Valid {
				personValue := int(position.GPerson.Int64)
				person = &personValue
			}
			nCase = &lemmas.NCase{
				Case:   _case.Case(position.GCase.Int64),
				Number: number.Number(position.GNumber.Int64),
				Person: person,
				Gender: gender.Gender(position.GGender.Int64), // check valid?
			}
		}
		if position.GTense.Valid {
			vCase = &lemmas.VCase{
				Person: int(position.GPerson.Int64),
				Number: number.Number(position.GNumber.Int64),
				Tense:  tense.Tense(position.GTense.Int64),
				Gender: gender.Gender(position.GGender.Int64), // check valid?
			}
		}
		res = append(res, &lemmas.Lemma{
			Normal: flexy.R.LemmaIdLemmas.Value,
			Value:  flexy.Value,
			Pos:    pos.POS(flexy.R.LemmaIdLemmas.Pos),
			NCase:  nCase,
			VCase:  vCase,
		})
	}
	return res, nil
}

func Normalize(word string, config lemmas.FlexyConfig) (*lemmas.Lemma, error) {
	switch config.POS {
	case pos.NOUN:
		return NormalizeNoun(word, config.NConfig)
	default:
		return nil, nil
	}
}

func NormalizeNoun(word string, config lemmas.NounFlexyConfig) (*lemmas.Lemma, error) {
	flexies, err := dictionary.GetFlexForm(word)
	if err != nil {
		return nil, err
	}
	var lemma *lemmas.Lemma
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
				lemma = &lemmas.Lemma{
					Normal: lemmaValue.Value,
					Value:  form.Value,
					Pos:    pos.POS(lemmaValue.Pos),
					NCase: &lemmas.NCase{
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
