package main

type Solution struct{}

func (Solution) CanConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}

	var letters = make(map[byte]int8, len(magazine))

	for i := range magazine {
		letters[magazine[i]] += 1
	}

	for i := range ransomNote {
		if _, ok := letters[ransomNote[i]]; ok {
			letters[ransomNote[i]] -= 1
			if letters[ransomNote[i]] < 0 {
				return false
			}
		} else {
			return false
		}
	}

	return true
}
