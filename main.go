package main

import "strconv"

type Solution struct{}

func (Solution) GroupAnagrams(words []string) int {
	var groups = make(map[byte]int8)

	for i := range words {
		var key byte
		for j := range words[i] {
			key += words[i][j]
		}

		groups[key] += 1
	}

	var res int
	for range groups {
		res++
	}

	return res
}

// for leetcode: https://leetcode.com/problems/group-anagrams/
func groupAnagrams(strs []string) [][]string {
	if len(strs) == 1 {
		return [][]string{strs}
	}

	var groups = make(map[string][]string)
	for i := range strs {
		letters := [26]int16{}
		for j := range strs[i] {
			letters[strs[i][j]-'a']++
		}

		var key string
		for l := range letters {
			if letters[l] > 0 {
				key += string('a'+l) + strconv.Itoa(int(letters[l]))
			}
		}

		groups[key] = append(groups[key], strs[i])
	}

	var res = make([][]string, 0, len(groups))
	for k := range groups {
		res = append(res, groups[k])
	}

	return res
}
