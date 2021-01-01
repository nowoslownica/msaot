package syntax

import (
	"fmt"
	"github.com/eakarpov/msaot/graphematical"
	"github.com/eakarpov/msaot/lexicon/dictionary"
	"github.com/eakarpov/msaot/morphological"
	"testing"
)

func TestBuildSyntaxTree(t *testing.T) {
	_, close, err := dictionary.InitMain()
	if err != nil {
		t.Error(err)
	}
	defer close()
	tokenizer := graphematical.Tokenizer{}
	tokenizer.Parse("Я есмь сказал, что я есмь хотел.")
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
		fmt.Println(len(tree.ComplexSentence)) // 2
		fmt.Println(tree.ComplexSentence[0].Predicate.Value) // есмь сказал
		fmt.Println(tree.ComplexSentence[0].Joiner) // empty
		fmt.Println(tree.ComplexSentence[0].Subject.Value) // я +
	}
}
