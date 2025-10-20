package main

type Solution struct{}

// Brian Kernighan algorythm
func (_ Solution) HammingDistance(x int, y int) int {
	count := 0
	n := x ^ y

	for n != 0 {
		count++
		n &= n - 1
	}

	return count
}
