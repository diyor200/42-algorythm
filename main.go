package main

type Solution struct{}

func (_ Solution) HammingWeight(n int) int {
	count := 0

	for n != 0 {
		count += n & 1
		n = n >> 1
	}

	return count
}
