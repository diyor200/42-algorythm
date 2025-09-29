package main

type Solution struct{}

func (Solution) Generate(n int) [][]int {
	if n == 0 {
		return [][]int{}
	}

	row := []int{1}
	res := make([][]int, n)
	res[0] = row

	for i := 1; i < n; i++ {
		row = generateRow(row)
		res[i] = row
	}

	return res
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
