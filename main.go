package main

import "fmt"

// brute-force
func isPalindromeSlow(s string) bool {
	// A= 65, a= 97
	var res string
	for i := range s {
		switch {
		case s[i] > 96 && s[i] < 123:
			res += string(s[i])
		case s[i] < 91 && s[i] > 64:
			fmt.Println(s[i])
			res += string(s[i] + 32)
		case s[i] > 47 && s[i] < 58:
			res += string(s[i])
		default:
			continue
		}

	}

	if res != "" {
		j := len(res) - 1
		i := 0
		for i < j {
			if res[i] != res[j] {
				return false
			}
			i++
			j--
		}
	}

	return true
}

// two pointer
func isPalindrome(s string) bool {
	i := 0
	j := len(s) - 1
	for i < j {
		var a, b string
		switch {
		case s[i] > 96 && s[i] < 123:
			a = string(s[i])
		case s[i] < 91 && s[i] > 64:
			a = string(s[i] + 32)
		case s[i] > 47 && s[i] < 58:
			a = string(s[i])
		default:
			i++
			continue
		}

		switch {
		case s[j] > 96 && s[j] < 123:
			b = string(s[j])
		case s[j] < 91 && s[j] > 64:
			b = string(s[j] + 32)
		case s[j] > 47 && s[j] < 58:
			b = string(s[j])
		default:
			j--
			continue
		}

		if a != b {
			return false
		}
		i++
		j--
	}

	return true
}
