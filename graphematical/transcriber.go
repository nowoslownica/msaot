package graphematical

var importedLat2SimpleCyrMap = map[rune]interface{}{
	'a': 'а',
	'å': 'а',
	'b': 'б',
	'c': 'ц',
	'ć': "чі",
	'č': 'ч',
	'd': 'д',
	'đ': "джі",
	'e': 'е',
	'ė': 'е',
	'ę': 'е',
	'ě': 'ѣ',
	'f': 'ф',
	'g': 'г',
	'h': 'х',
	'i': 'и',
	'j': 'і',
	'k': 'к',
	'l': 'л',
	'ľ': "лі",
	'm': 'м',
	'n': 'н',
	'ń': "ні",
	'o': 'о',
	'ȯ': 'о',
	'p': 'п',
	'r': 'р',
	'ř': "рі",
	's': 'с',
	'ś': "сі",
	'š': 'ш',
	't': 'т',
	'u': 'у',
	'ų': 'у',
	'v': 'в',
	'y': 'ы',
	'z': 'з',
	'ź': "зі",
	'ž': 'ж',
}

func ImportedLat2SimpleCyr(latWord string) string {
	res := ""

	for _, letter := range latWord {
		substr := importedLat2SimpleCyrMap[letter]

		if i, ok := substr.(rune); ok {
			res += string(i)
		} else if i, ok := substr.(string); ok {
			res += i
		}
	}

	return res
}
