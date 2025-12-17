package main

import "fmt"

func arrangeCoins(n int) int {
	rows := 0
	for i := 1; n-i >= 0; i++ {
		n -= i
		rows++
	}

	return rows
}

func main() {
	fmt.Println(arrangeCoins(11))
}
