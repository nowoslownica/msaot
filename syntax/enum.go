package syntax

import "github.com/eakarpov/msaot/morphological"

type Noun struct {
	Value *morphological.Lemma
}

type Verb struct {
	Value *morphological.Lemma
}

type Adjective struct {
	Value *morphological.Lemma
}

type POS struct {
	Noun   *Noun
	Verb   *Verb
	Joiner *morphological.Lemma
}

type Collocation struct {
	Type        string
	Preposition *morphological.Lemma
	Value       *POS
	Dependent   []*POS
	Joiner      *morphological.Lemma
}

type Subject struct {
	Value     *morphological.Lemma
	Dependent []*Collocation
}

type Predicate struct {
	Value *morphological.Lemma
}

type Sentence struct {
	Joiner    *morphological.Lemma // before the item
	Subject   *Subject
	Predicate *Predicate
}

type SyntaxTree struct {
	ComplexSentence []*Sentence
}
