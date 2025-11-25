package main

func flipAndInvertImage(image [][]int) [][]int {
	for img := range image {
		reverse(image[img])
		for i := range image[img] {
			image[img][i] ^= 1
		}
	}

	return image
}

func reverse(nums []int) {
	n := len(nums) - 1
	for i := 0; i < n; i++ {
		nums[i], nums[n] = nums[n], nums[i]
		n--
	}
}
