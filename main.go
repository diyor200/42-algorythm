package main

type Solution struct{}

func (Solution) MaxArea(nums []int) int {
	i, j := 0, len(nums)-1
	result := 0

	for i < j {
		h := 0
		if nums[i] < nums[j] {
			h = nums[i]
			i++
		} else {
			h = nums[j]
			j--
		}
		area := (j - i + 1) * h // +1 bo‘lishi shart emas, lekin to‘g‘ri joyda kamaytirilganiga qarab ishlatiladi
		if area > result {
			result = area
		}
	}

	return result
}
