package morphological

import (
	"fmt"
	"github.com/eakarpov/msaot/graphematical"
	"github.com/eakarpov/msaot/lexicon"
	"testing"
)

func TestBuildLemmaChains(t *testing.T) {
	_, closeF, err := lexicon.InitMain()
	if err != nil {
		return
	}
	defer func() {
		_ = closeF()
	}()

	empty := make([][]*Lemma, 0)
	testCases := []struct{
		Data []*graphematical.Token
		Expected [][]*Lemma
	}{
		{[]*graphematical.Token{
			{Value: "Я", Position: 0},
			{Value: "есмь", Position: 1},
			{Value: "пошёл", Position: 2},
			{Value: "гулять", Position: 3},
			{Value: "по", Position: 4},
			{Value: "красивому", Position: 5},
			{Value: "городу", Position: 6},
			{Value: ".", Position: 7},
		}, empty},
	}
	for _, test := range testCases {
		lemmasArrs, err := BuildLemmaChains(test.Data)
		if err != nil {
			t.Error("Error while building lemma chains", err)
		}
		for _, lemmas := range lemmasArrs {
			fmt.Println(lemmas)
		}
	}
}
