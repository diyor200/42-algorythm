package main

import "fmt"

func convertToTitle(columnNumber int) string {
	if columnNumber < 27 {
		return string(byte('A' - 1 + columnNumber))
	}

	var res []byte

	for columnNumber > 0 {
		columnNumber -= 1
		res = append(res, byte(columnNumber%26+'A'))
		columnNumber /= 26
	}

	j := len(res) - 1
	for i := 0; i < j; i++ {
		res[i], res[j] = res[j], res[i]
		j--
	}

	return string(res)
}

func main() {
	fmt.Println(convertToTitle(25))
}
