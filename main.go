package main

type Solution struct{}

func (Solution) DetectCapitalUse(word string) bool {
	if isUpper(word) || isLower(word) {
		return true
	}

	return isUpper(word[:1]) && isLower(word[1:])
}

func isUpper(s string) bool {
	return inRange(s, 'A', 'Z')
}
func isLower(s string) bool { return inRange(s, 'a', 'z') }

func inRange(s string, start, end rune) bool {
	for _, ch := range s {
		if !(ch >= start && ch <= end) {
			return false
		}
	}

	return true
}
