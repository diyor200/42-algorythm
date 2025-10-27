package main

type Solution struct{}

func findComplement(num int) int {
	if num == 0 {
		return 1
	} else if num == 1 {
		return 0
	}

	// find MSB(most significant bit)
	pos := 0
	for i := 0; i < 32; i++ {
		if (num>>i)&1 == 1 {
			pos = i
		}
	}

	mask := 1<<(pos+1) - 1
	return num ^ mask
}
