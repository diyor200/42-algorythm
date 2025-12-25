package main

import "fmt"

// url:https://leetcode.com/problems/excel-sheet-column-number/submissions/1865190077/?envType=problem-list-v2&envId=string

func titleToNumber(columnTitle string) int {
	n := len(columnTitle)
	if n == 1 {
		return 1 + int(columnTitle[0]-'A')
	}

	res := int(columnTitle[n-1] - 'A' + 1)
	j := 1

	for i := n - 2; i >= 0; i-- {
		res += int(columnTitle[i]-'A'+1) * pow(26, j)
		j++
	}

	return res
}

func pow(p, n int) int {
	res := 1
	for i := 0; i < n; i++ {
		res *= p
	}

	return res
}

func main() {
	fmt.Println(titleToNumber("AA"))
}
