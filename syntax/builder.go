package syntax

import (
	"github.com/eakarpov/msaot/lexicon/pos"
	"github.com/eakarpov/msaot/morphological"
)

func BuildSyntaxTree(sentence []*morphological.Lemma) (*SyntaxTree, error) {
	var complexSentence = make([]*Sentence, 0)
	// zero step - find single sentences

	// for each single sentence
	var singleSentence Sentence

	complexSentence = append(complexSentence, &singleSentence)

	// first step - find subject
	s, err := findSubject(sentence)
	if err != nil {
		return nil, err
	}
	singleSentence.Subject = s

	// second step - find predicate
	pr, err := findPredicate(sentence)
	if err != nil {
		return nil, err
	}
	singleSentence.Predicate = pr

	return &SyntaxTree{ ComplexSentence: complexSentence }, nil
}

func findPredicate(sentence []*morphological.Lemma) (*Predicate, error) {
	var pr *Predicate

	for _, word := range sentence {
		if word.Pos == pos.VERB && word.Value != word.Normal {

		}
	}

	return pr, nil
}

func findSubject(sentence []*morphological.Lemma) (*Subject, error) {
	var s *Subject
	for _, word := range sentence {
		if word.Value == word.Normal &&
			(word.Pos == pos.NOUN || word.Pos == pos.PRONOUN) {
			// without conjunction group
			s = &Subject{
				Value:     word,
				Dependent: nil,
			}
		}
	}
	return s, nil
}
