package lemmas

import (
	_case "github.com/eakarpov/msaot/lexicon/case"
	"github.com/eakarpov/msaot/lexicon/gender"
	"github.com/eakarpov/msaot/lexicon/number"
	"github.com/eakarpov/msaot/lexicon/pos"
	"github.com/eakarpov/msaot/lexicon/tense"
	"strings"
)

type NCase struct {
	Case   _case.Case
	Number number.Number
	Person *int
	Gender gender.Gender
}

type VCase struct {
	Person int
	Gender gender.Gender
	Number number.Number
	Tense  tense.Tense
}

type Lemma struct {
	Normal       string
	Value        string
	Pos          pos.POS
	Indeclinable bool
	NCase        *NCase
	VCase        *VCase
	Comma        bool // ???
	Animate      *bool
}

type GrammarConfig struct {
	POS    string
	Gender *string
	Type   string
}

type FlexyConfig struct {
	POS     pos.POS
	Type    string
	NConfig NounFlexyConfig
}

type NounFlexyConfig struct {
	Case   _case.Case
	Number int
	Person string
}

func GetLemmaConfig(grammarString string) *Lemma {
	posValue := pos.POS("")
	genderValue := gender.Gender(0)
	numberValue := number.Number(0)
	indeclinable := false
	var animate *bool

	switch {
	case grammarString == "adj.":
	case grammarString == "adv.":
	case grammarString == "prep.":
		{
			posValue = pos.POS(strings.TrimRight(grammarString, "."))
		}
	case strings.Contains(grammarString, "v."):
		{
			posValue = pos.VERB
		}
	case strings.Contains(grammarString, "indecl."): {
		indeclinable = true
	}
	case strings.Contains(grammarString, "pron."): {
		posValue = pos.PRONOUN
	}
	case strings.Contains(grammarString, "f."):
		{
			posValue = pos.NOUN
			genderValue = gender.FEMININE
		}
	case strings.Contains(grammarString, "m."):
		{
			posValue = pos.NOUN
			genderValue = gender.MASCULINE
		}
	case strings.Contains(grammarString, "n."):
		{
			posValue = pos.NOUN
			genderValue = gender.NEUTRAL
		}
	case strings.Contains(grammarString, "anim."):
		{
			animateValue := true
			animate = &animateValue
		}
	case strings.Contains(grammarString, "sg."):
		{
			numberValue = number.SINGULAR
		}
	case strings.Contains(grammarString, "pl."):
		{
			numberValue = number.PLURAL
		}
	default:
	}
	var nCase *NCase
	if genderValue != 0 {
		nCase = &NCase{
			Number: numberValue,
			Gender: genderValue,
		}
	}
	return &Lemma{
		Pos:          posValue,
		NCase:        nCase,
		VCase:        nil,
		Animate:      animate,
		Indeclinable: indeclinable,
	}
}

