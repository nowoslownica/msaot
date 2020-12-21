package synthesizer

import (
	"github.com/eakarpov/msaot/graphematical"
	"github.com/eakarpov/msaot/lexicon/gender"
	"github.com/eakarpov/msaot/lexicon/lemmas"
	"github.com/eakarpov/msaot/lexicon/pos"
	"strings"
)

func getNounForms(word string, dType GrammarDeclensionType) []*lemmas.Lemma {
	res := make([]*lemmas.Lemma, 0)
	base := strings.TrimSuffix(word, dType.Ending)
	for _, form := range dType.Forms {
		val := strings.Join([]string{base, form.Ending}, "")
		res = append(res, &lemmas.Lemma{
			Normal:       word,
			Value:        val,
			Pos:          pos.NOUN,
			Indeclinable: false,
			NCase:        &lemmas.NCase{
				Case:   form.Case,
				Number: form.Number,
				Gender: dType.Gender,
			},
		})
	}
	return res
}

func GetNounFormsByEndingAndGender(word string, gend gender.Gender) []*lemmas.Lemma {
	lastLetter := rune(word[len(word) - 1])
	for _, dType := range declensions {
		if dType.Ending != "" && strings.HasSuffix(word, dType.Ending) && gend == dType.Gender {
			return getNounForms(word, dType)
		} else if dType.Ending == "" &&
			!graphematical.IsVowel(lastLetter) &&
			gend == dType.Gender {
			return getNounForms(word, dType)
		}
	}
	return nil
}

func GetAdjectiveFormsByEnding() []lemmas.Lemma {
	return nil
}

func GetWordForms(word string, posValue pos.POS, genderValue gender.Gender, animate bool) []*lemmas.Lemma {
	switch posValue {
	case pos.NOUN:
		return GetNounFormsByEndingAndGender(word, genderValue)
	default:
		return nil
	}
}
