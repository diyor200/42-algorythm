package main

var vowels = map[rune]bool{
	'a': true,
	'e': true,
	'i': true,
	'o': true,
	'u': true,
	'A': true,
	'E': true,
	'I': true,
	'O': true,
	'U': true,
}

func reverseVowels(s string) string {
	n := len(s)
	var res = make([]byte, n)
	i := 0
	j := n - 1
	for i <= j {
		if !vowels[rune(s[i])] {
			res[i] = s[i]
			i++
			continue
		}

		if !vowels[rune(s[j])] {
			res[j] = s[j]
			j--
			continue
		}

		if vowels[rune(s[i])] && vowels[rune(s[j])] {
			res[i], res[j] = s[j], s[i]
		}

		i++
		j--
	}

	return string(res)
}
