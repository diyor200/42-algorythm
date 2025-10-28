package main

func getRow(rowIndex int) []int {
	return pascalsTriangle(rowIndex + 1)
}

func pascalsTriangle(n int) []int {
	row := []int{1}

	for i := 1; i < n; i++ {
		row = generateRow(row)
	}

	return row
}

func generateRow(prev []int) []int {
	row := make([]int, len(prev)+1)
	row[0] = 1
	row[len(row)-1] = 1

	for i := 0; i < len(prev)-1; i++ {
		row[i+1] = prev[i] + prev[i+1]
	}

	return row
}
