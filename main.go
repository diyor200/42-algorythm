package main

func containsNearbyDuplicate(nums []int, k int) bool {
	for i := 0; k > 0 && i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				return true
			}
			if k == (j - i) {
				break
			}
		}
	}

	return false
}

func containsNearbyDuplicate(nums []int, k int) bool {
	if k == 0 {
		return false
	}

	n := len(nums)
	if k > n {
		k = n
	}

	seen := make(map[int]bool, k)
	for i := 0; i < k; i++ {
		if _, ok := seen[nums[i]]; ok {
			return true
		}
		seen[nums[i]] = true
	}

	for i := k; i < len(nums); i++ {
		if _, ok := seen[nums[i]]; ok {
			return true
		}
		delete(seen, nums[i-k])
		seen[nums[i]] = true
	}

	return false
}
