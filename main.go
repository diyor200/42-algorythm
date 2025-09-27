package main

type Solution struct{}

// slice version: T: O(n), S: O(1)
func (Solution) IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var letters [26]int16
	for i := range s {
		letters[s[i]-'a'] += 1
		letters[t[i]-'a'] -= 1
	}

	for k := range letters {
		if letters[k] != 0 {
			return false
		}
	}

	return true
}

// map based vesrion: T: O(n), S: O(k)(worst: O(n))
// func (Solution) IsAnagram(s string, t string) bool {
// 	if len(s) != len(t) {
// 		return false
// 	}

// 	var letters = make(map[byte]int, len(s))
// 	for i := range s {
// 		letters[s[i]] += 1
// 	}

// 	for i := range t {
// 		letters[t[i]] -= 1
// 		if letters[t[i]] < 0 {
// 			return false
// 		}
// 	}

// 	for k := range letters {
// 		if letters[k] != 0 {
// 			return false
// 		}
// 	}

// 	return true
// }
