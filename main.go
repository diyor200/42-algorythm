package main

import (
	"fmt"
	"math"
)

func main() {
	b := byte(48)
	fmt.Println(int(b))
	fmt.Println(tenPow(-1))
	s := Solution{}
	fmt.Println(s.MyAtoi("9223372036854775808"))
}

type Solution struct{}

func (_ Solution) MyAtoi(s string) int {
	zero := byte(48)
	minus := byte(45)
	negative := false

	nums := []int{}
	startIndex := 0

	for i := range s {
		if s[i] == ' ' || s[i] == '-' || s[i] == '+' {
			if len(nums) != 0 {
				break
			}

			if i > 0 && (s[i-1] == '-' || s[i-1] == '+') {
				return 0
			}

			continue
		} else if s[i]-zero < 0 || s[i]-zero > 9 {
			if len(nums) >= 0 {
				break
			}
		}

		if s[i]-zero > 9 && len(nums) == 0 {
			return 0
		}

		if s[i]-zero >= 0 {
			if len(nums) == 0 {
				startIndex = i
			}

			nums = append(nums, int(s[i]-zero))
		}
	}

	if startIndex != 0 {
		if s[startIndex-1] == minus {
			negative = true
		}
	}

	var res int64
	n := len(nums)
	for i := range nums {
		res += int64(nums[i] * tenPow(n-i-1))
	}

	if negative {
		res = -res
	} else if res < 0 {
		if math.MinInt32 > (-1 * res) {
			res = math.MaxInt32
		}
	}

	if math.MaxInt32 < res {
		res = math.MaxInt32
	}

	if math.MinInt32 > res {
		res = math.MinInt32
	}

	return int(res)
}

func tenPow(n int) int {
	res := 1
	for i := 0; i < n; i++ {
		res *= 10
	}

	return res
}
