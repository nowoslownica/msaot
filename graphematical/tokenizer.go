package graphematical

import (
	"strings"
)

type Token struct {
	Value    string
	Position int
	Quoted   bool
}

type Tokenizer struct {
	Items []*Token
}

func RuneIn(arr []rune, r rune) bool {
	for _, rn := range arr {
		if rn == r {
			return true
		}
	}
	return false
}

func (t *Tokenizer) Parse(sentence string) *Tokenizer {
	tokens := make([]*Token, 0)
	pos := 0
	tmp := ""
	quoted := false
	dequoted := false
	quoteCount := 0
	prev := rune(0)
	for i, l := range sentence {
		isPnkt := RuneIn([]rune{',', '.', ':', '-', '?', '!'}, l)

		if RuneIn([]rune{'"', '«', '»'}, l) {
			if prev == ' ' {
				quoted = true
				quoteCount++
				if quoteCount == 1 { // только открыли группу в кавычках, сохранять не нужно
					prev = l
					continue
				}
			} else {
				dequoted = true
				isLast := false
				j := i + 1
				index := strings.Index(sentence[j:], string(l))
				if index != -1 {
					further := rune(0)
					if l == '"' {
						further = '"'
					} else {
						further = '«'
					}
					fIndex := strings.Index(sentence[j:], string(further))
					if fIndex < index {
						isLast = true
					}
				} else {
					isLast = true
				}
				if isLast {
					quoteCount--
				}
			}
		}

		if l != ' ' && !isPnkt {
			if dequoted {
				dequoted = false
				if quoteCount > 0 {
					quoteCount--
				} else {
					prev = l
					continue
				}
			}
			tmp += string(l)
			prev = l
			continue
		}

		if tmp != "" {
			if quoteCount == 0 {
				tokens = append(tokens, &Token{
					Value:    tmp,
					Position: pos,
					Quoted:   quoted,
				})
				quoted = false
				tmp = ""
				pos++
			} else {
				tmp += " " // here the space case is considered
			}
		}
		if isPnkt {
			tokens = append(tokens, &Token{
				Value:    string(l),
				Position: pos,
			})
		}
		prev = l
	}
	t.Items = tokens
	return t
}
