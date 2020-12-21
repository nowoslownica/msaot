package synthesizer

import (
	_case "github.com/eakarpov/msaot/lexicon/case"
	"github.com/eakarpov/msaot/lexicon/gender"
	"github.com/eakarpov/msaot/lexicon/number"
)

type DeclensionForm struct {
	Case   _case.Case
	Number number.Number
	Ending string
}

type GrammarDeclensionType struct {
	Ending string
	Gender gender.Gender
	Forms  []DeclensionForm
}

var firstDeclensionForms = []DeclensionForm{
	{ Case: _case.NOMINATIVE, Ending: "а", Number: number.SINGULAR},
	{ Case: _case.GENITIVE, Ending: "y",  Number: number.SINGULAR },
	{ Case: _case.DATIVE, Ending: "ě",  Number: number.SINGULAR },
	{ Case: _case.ACCUSATIVE, Ending: "u",  Number: number.SINGULAR },
	{ Case: _case.INSTRUMENTAL, Ending: "ojų",  Number: number.SINGULAR },
	{ Case: _case.LOCATIVE, Ending: "ě",  Number: number.SINGULAR },
	{ Case: _case.VOCATIVE, Ending: "o",  Number: number.SINGULAR },

	{ Case: _case.NOMINATIVE, Number: number.PLURAL, Ending: "y" },
	{ Case: _case.GENITIVE, Number: number.PLURAL, Ending: "" },
	{ Case: _case.DATIVE, Number: number.PLURAL, Ending: "am" },
	{ Case: _case.ACCUSATIVE, Number: number.PLURAL, Ending: "" },
	{ Case: _case.INSTRUMENTAL, Number: number.PLURAL, Ending: "ami" },
	{ Case: _case.LOCATIVE, Number: number.PLURAL, Ending: "ah" },
	{ Case: _case.VOCATIVE, Number: number.PLURAL, Ending: "y" },
}

var firstDeclensionA = GrammarDeclensionType{
	Ending: "a",
	Gender: 1,
	Forms:  firstDeclensionForms,
}

var firstDeclensionB = GrammarDeclensionType{
	Ending: "a",
	Gender: 2,
	Forms:  firstDeclensionForms,
}

var secondDeclensionA = GrammarDeclensionType{
	Ending: "",
	Gender: 1,
	Forms: []DeclensionForm{
		{ Case: _case.NOMINATIVE, Number: number.SINGULAR, Ending: "" },
		{ Case: _case.GENITIVE, Number: number.SINGULAR, Ending: "a" },
		{ Case: _case.DATIVE, Number: number.SINGULAR, Ending: "u" },
		{ Case: _case.ACCUSATIVE, Number: number.SINGULAR, Ending: "" },
		{ Case: _case.INSTRUMENTAL, Number: number.SINGULAR, Ending: "om" },
		{ Case: _case.LOCATIVE, Number: number.SINGULAR, Ending: "u" },
		{ Case: _case.VOCATIVE, Number: number.SINGULAR, Ending: "e" },

		{ Case: _case.NOMINATIVE, Number: number.PLURAL, Ending: "y" },
		{ Case: _case.GENITIVE, Number: number.PLURAL, Ending: "ov" },
		{ Case: _case.DATIVE, Number: number.PLURAL, Ending: "am" },
		{ Case: _case.ACCUSATIVE, Number: number.PLURAL, Ending: "y" },
		{ Case: _case.INSTRUMENTAL, Number: number.PLURAL, Ending: "ami" },
		{ Case: _case.LOCATIVE, Number: number.PLURAL, Ending: "ah" },
		{ Case: _case.VOCATIVE, Number: number.PLURAL, Ending: "je" },
	},
}

var declensions = []GrammarDeclensionType{
	firstDeclensionA,
	firstDeclensionB,
	secondDeclensionA,
}
