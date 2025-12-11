package main

import "fmt"

func mySqrt(x int) int {
	if x == 0 || x == 1 {
		return x
	}

	middle := int(x / 2)
	left, right := 0, x
	for left < right {
		sqr := middle * middle
		if sqr > x {
			right = middle
		} else if sqr < x {
			left = middle + 1
		} else {
			return middle
		}

		middle = (left + right) / 2
	}

	return left - 1
}

func main() {
	x := 25

	fmt.Println(mySqrt(x))
}
