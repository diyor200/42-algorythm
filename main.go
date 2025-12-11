package main

import "fmt"

func guessNumber(n int) int {
	middle := int(n / 2)
	res := guess(middle)
	left, right := 0, n
	for res != 0 {
		switch res {
		case 1:
			left = middle + 1
		case -1:
			right = middle
		case 0:
			break
		}

		middle = (left + right) / 2
		res = guess(middle)
	}

	return middle
}

func main() {
	n := 30

	fmt.Println(guessNumber(n))
}

func guess(x int) int {
	guessed := 7
	if x > guessed {
		return -1
	} else if x < guessed {
		return 1
	}

	return 0
}
