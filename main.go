package main

func isPerfectSquare(num int) bool {
	if num == 1 {
		return true
	}

	middle := int(num / 2)
	left, right := 0, num

	for left < right {
		sqr := middle * middle
		if sqr > num {
			right = middle
		} else if sqr < num {
			left = middle + 1
		} else {
			return true
		}

		middle = (left + right) / 2
	}

	return false
}
