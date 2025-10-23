package main

type Solution struct{}

// url: Find the Difference - LeetCode

func findTheDifference(s string, t string) byte {
	var res byte
	for i := range s {
		res ^= s[i] ^ t[i]
	}

	return res ^ t[len(t)-1]
}
