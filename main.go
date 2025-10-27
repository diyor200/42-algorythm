package main

import "fmt"

type Solution struct{}

func toHex(num int) string {
	if num == 0 {
		return "0"
	}

	hexChars := "0123456789abcdef"
	result := ""

	for i := 0; num != 0 && i < 8; i++ {
		result = string(hexChars[num&0xf]) + result
		num >>= 4
	}

	return result
}

func main() {
	fmt.Println(toHex(26))
}
