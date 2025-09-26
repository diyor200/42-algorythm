package main

type Solution struct{}

var set = map[rune]byte{
	'q': 1,
	'w': 1,
	'e': 1,
	'r': 1,
	't': 1,
	'y': 1,
	'u': 1,
	'i': 1,
	'o': 1,
	'p': 1,
	'a': 2,
	's': 2,
	'd': 2,
	'f': 2,
	'g': 2,
	'h': 2,
	'j': 2,
	'k': 2,
	'l': 2,
	'z': 3,
	'x': 3,
	'c': 3,
	'v': 3,
	'b': 3,
	'n': 3,
	'm': 3,
}

func (Solution) FindWords(words []string) []string {
	var res = []string{}
	for i := range words {
		ok := true
		var row byte
		for j := range words[i] {
			if row == 0 {
				row = set[toLower(rune(words[i][j]))]
			} else {
				ok = row == set[toLower(rune(words[i][j]))]
			}

			if !ok {
				break
			}
		}

		if ok {
			res = append(res, words[i])
		}
	}

	return res
}

func toLower(r rune) rune {
	if r < 97 {
		return r + 32
	}

	return r
}
