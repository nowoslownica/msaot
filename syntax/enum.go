package syntax

import (
	"github.com/eakarpov/msaot/lexicon/lemmas"
)

type Noun struct {
	Value *lemmas.Lemma
}

type Verb struct {
	Value *lemmas.Lemma
}

type Adjective struct {
	Value *lemmas.Lemma
}

type POS struct {
	Noun   *Noun
	Verb   *Verb
	Joiner *lemmas.Lemma
}

type Collocation struct {
	Type        string
	Preposition *lemmas.Lemma
	Value       *POS
	Dependent   []*POS
	Joiner      *lemmas.Lemma
}

type Subject struct {
	Value     *lemmas.Lemma
	Dependent []*Collocation
}

type Predicate struct {
	Value *lemmas.Lemma
}

type Sentence struct {
	Joiner    *lemmas.Lemma // before the item
	Subject   *Subject
	Predicate *Predicate
}

type SyntaxTree struct {
	ComplexSentence []*Sentence
}
