package main

type Solution struct{}

// FindJudge map based solution:
// T: O(n + len(trues)
// S: O(n)
// FindJudge
//func (_ Solution) FindJudge(n int, trust [][]int) int {
//	var (
//		from = make(map[int]int)
//		to   = make(map[int]int)
//	)
//
//	for i := range trust {
//		from[trust[i][0]]++
//		to[trust[i][1]]++
//	}
//
//	for i := 1; i < n+1; i++ {
//		if from[i] == 0 && to[i] == n-1 {
//			return i
//		}
//	}
//
//	return -1
//}

func (_ Solution) FindJudge(n int, trust [][]int) int {
	var count = make([]int, n+1)

	for i := range trust {
		count[trust[i][0]]--
		count[trust[i][1]]++
	}

	for i := 1; i < n+1; i++ {
		if count[i] == n-1 {
			return i
		}
	}

	return -1
}
