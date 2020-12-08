package syntax

import (
	"fmt"
	"github.com/eakarpov/msaot/graphematical"
	"github.com/eakarpov/msaot/lexicon"
	"github.com/eakarpov/msaot/morphological"
	"testing"
)

func TestBuildSyntaxTree(t *testing.T) {
	_, close, err := lexicon.InitMain()
	if err != nil {
		t.Error(err)
	}
	defer close()
	tokenizer := graphematical.Tokenizer{}
	tokenizer.Parse("Я сказал, что я хотел.")
	LoS, err := morphological.BuildLemmaChains(tokenizer.Items)
	if err != nil {
		t.Error(err)
	}
	sentences := morphological.BuildSentenceLemmaChains(LoS)
	for _, sntnce := range sentences {
		tree, err := BuildSyntaxTree(sntnce.Lemmas)
		if err != nil {
			t.Error(err)
		}
		fmt.Println(tree.ComplexSentence[0].Subject.Value)
	}
}
