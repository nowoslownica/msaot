package graphematical

var vowels = []rune{'a', 'o', 'u', 'i', 'y', 'ė', 'å', 'ę', 'ȯ', 'ų'}

func IsVowel(r rune) bool {
	for _, v := range vowels {
		if v == r {
			return true
		}
	}
	return false
}
