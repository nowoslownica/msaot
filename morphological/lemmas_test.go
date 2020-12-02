package morphological

import (
	"fmt"
	"github.com/eakarpov/msaot/lexicon"
	"testing"
)

func TestNormalize(t *testing.T) {
	_, closeF, err := lexicon.InitMain()
	if err != nil {
		return
	}
	lemma, err := Normalize("мамы", FlexyConfig{
		POS: "n",
		Type: "1",
		Case: "Genitive",
		Number: "singular",
	})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(lemma)
	defer func() {
		_ = closeF()
	}()
}
