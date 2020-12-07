package morphological

import (
	"github.com/eakarpov/msaot/graphematical"
	"strings"
)

func BuildLemmaChains(tokens []*graphematical.Token) ([][]*Lemma, error) {
	res := make([][]*Lemma, 0)
	res = append(res, make([]*Lemma, 0))
	for _, token := range tokens {
		lemmas, err := NormalizeAuto(strings.ToLower(token.Value))
		if err != nil {
			return nil, err
		}
		if len(lemmas) == 1 {
			for i := range res {
				res[i] = append(res[i], lemmas[0])
			}
		}
		// TODO: if more than 1 - create a doubled slice and copy previous values
	}
	return res, nil
}
