package morphological

import (
	"fmt"
	_case "github.com/eakarpov/msaot/lexicon/case"
	"github.com/eakarpov/msaot/lexicon/dictionary"
	"github.com/eakarpov/msaot/lexicon/pos"
	"testing"
)

func TestNormalize(t *testing.T) {
	_, closeF, err := dictionary.InitMain()
	if err != nil {
		return
	}
	lemma, err := Normalize("мамы", FlexyConfig{
		POS: pos.NOUN,
		//Type: "1",
		nConfig: NounFlexyConfig{
			Case: _case.GENITIVE,
			Number: 1,
		},
	})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(lemma)
	defer func() {
		_ = closeF()
	}()
}
