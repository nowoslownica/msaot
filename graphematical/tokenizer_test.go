package graphematical

import "testing"

func TokenContains(arr []*Token, item *Token) bool {
	for _, elem := range arr {
		if elem.Value == item.Value && elem.Position == item.Position {
			return true
		}
	}
	return false
}

func TestTokenizer_Parse(t *testing.T) {
	testCases := []struct{
		Data string
		Expected []*Token
	}{
		{"Я пошёл гулять.", []*Token{
			{Value: "Я", Position: 0},
			{Value: "пошёл", Position: 1},
			{Value: "гулять", Position: 2},
			{Value: ".", Position: 3},
		}},
		{ `Я работаю в компании «НПО «Эшелон».`, []*Token{
			{ Value: "Я", Position: 0},
			{ Value: "работаю", Position: 1},
			{ Value: "в", Position: 2},
			{ Value: "компании", Position: 3},
			{ Value: "НПО «Эшелон»", Position: 4, Quoted: true},
			{ Value: ".", Position: 5},
		}},
		{ `Я работаю в компании "НПО "Эшелон".`, []*Token{
			{ Value: "Я", Position: 0},
			{ Value: "работаю", Position: 1},
			{ Value: "в", Position: 2},
			{ Value: "компании", Position: 3},
			{ Value: `НПО "Эшелон"`, Position: 4, Quoted: true},
			{Value: ".", Position: 5},
		}},
		{ `Когда "Компания "веселых" друзей" открылась на побережье?`, []*Token{
			{ Value: "Когда", Position: 0},
			{ Value: `Компания "веселых" друзей`, Position: 1, Quoted: true},
			{ Value: "открылась", Position: 2},
			{ Value: "на", Position: 3},
			{ Value: `побережье`, Position: 4},
			{ Value: "?", Position: 5},
		}},
		{ `Я сказал, что это такое.`, []*Token{
			{ Value: "Я", Position: 0},
			{ Value: "сказал", Position: 1},
			{ Value: ",", Position: 2},
			{ Value: "что", Position: 3},
			{ Value: `это`, Position: 4},
			{ Value: "такое", Position: 5},
			{ Value: ".", Position: 6},
		}},
	}
	for _, test := range testCases {
		tokenizer := Tokenizer{}
		tokenizer.Parse(test.Data)
		for _, item := range tokenizer.Items {
			if !TokenContains(test.Expected, item) {
				t.Errorf("Token not found: %s", item.Value)
				return
			}
		}
	}
}
