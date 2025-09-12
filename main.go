package main

type Solution struct{}

func (Solution) ReverseString(s []string) []string {
	i := 0
	j := len(s) - 1

	for i < j {
		s[i], s[j] = s[j], s[i]

		i++
		j--
	}

	return s
}
