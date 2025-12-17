package main

func nextGreatestLetter(letters []byte, target byte) byte {
	n := len(letters)
	if letters[n-1] <= target || letters[0] > target {
		return letters[0]
	}

	var idx int
	left, right := 0, n
	middle := 0
	for left < right {
		middle = (left + right) / 2
		if letters[middle] == target {
			for {
				middle++
				if letters[middle] > target {
					idx = middle
					break
				}
			}
			break
		} else if letters[middle] > target {
			right = middle
		} else {
			left = middle + 1
		}
	}

	if idx == 0 {
		return letters[right]
	}

	return letters[idx]
}
