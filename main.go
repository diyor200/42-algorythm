package main

func longestPalindrome(s string) int {
	if len(s) == 1 {
		return 1
	}

	letters := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		letters[s[i]]++
	}

	res := 0
	for k := range letters {
		if letters[k] > 1 {
			res += letters[k] / 2
			delete(letters, k)
		}
	}

	res *= 2
	if len(letters) > 0 {
		res += 1
	}

	return res
}
