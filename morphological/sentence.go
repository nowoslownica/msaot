package morphological

import (
	"github.com/eakarpov/msaot/graphematical"
	"strings"
)

type SentenceLemmaChain struct {
	ChainPosition int
	Lemmas        []*Lemma
}

func BuildSentenceLemmaChains(chains [][]*LemmaOrStop) []SentenceLemmaChain {
	res := make([]SentenceLemmaChain, 0)
	for i, chain := range chains {
		temp := make([]*Lemma, 0)
		for _, lemma := range chain {
			if lemma.Stop {
				res = append(res, SentenceLemmaChain{
					ChainPosition: i,
					Lemmas:        temp,
				})
				temp = make([]*Lemma, 0)
			}
			temp = append(temp, lemma.Lemma)
		}
	}
	return res
}

type LemmaOrStop struct {
	Lemma *Lemma
	Stop  bool
}

func BuildLemmaChains(tokens []*graphematical.Token) ([][]*LemmaOrStop, error) {
	res := make([][]*LemmaOrStop, 0)
	res = append(res, make([]*LemmaOrStop, 0))
	for _, token := range tokens {
		if token.Value == "." {
			for i := range res {
				res[i] = append(res[i], &LemmaOrStop{ Stop: true })
			}
			continue
		}
		lemmas, err := NormalizeAuto(strings.ToLower(token.Value))
		if err != nil {
			return nil, err
		}
		if len(lemmas) == 1 {
			for i := range res {
				res[i] = append(res[i], &LemmaOrStop{ Lemma: lemmas[0], Stop: false })
			}
		}
		// TODO: if more than 1 - create a doubled slice and copy previous values
	}
	return res, nil
}
