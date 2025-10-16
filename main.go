package main

// T: O(n)
// S: O(n)
// func findCenter(edges [][]int) int {
// 	counter := make(map[int]int)
// 	n := len(edges)
// 	for i := 0; i < n; i++ {
// 		counter[edges[i][0]]++
// 		counter[edges[i][1]]++
// 	}

// 	res := 0
// 	for key := range counter {
// 		if counter[key] == n {
// 			res = key
// 			break
// 		}
// 	}
// 	return res
// }

// T: O(1)
// S: O(1)
func findCenter(edges [][]int) int {
	if edges[0][0] == edges[1][0] || edges[0][0] == edges[1][1] {
		return edges[0][0]
	}

	return edges[0][1]
}
