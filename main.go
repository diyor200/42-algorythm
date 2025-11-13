package main

//func reverseStr(s string, k int) string {
//	var res = make([]byte, len(s))
//
//	for i := range res {
//		res[i] = s[i]
//	}
//
//	start := 0
//	to := 2*k - 1
//	for start < len(s) {
//		switch {
//		case to > len(s):
//			reverse(res, start, len(s)-1)
//			break
//		case start+2*k < len(s):
//			res[start], res[start+1] = res[start+1], res[start]
//		case start+2*k <= len(s):
//		}
//
//		start = to + 1
//		to = start + 2*k - 1
//	}
//
//	return string(res)
//}

func reverse(chars []byte, from, to int) {
	for from < to {
		chars[from], chars[to] = chars[to], chars[from]
		from++
		to--
	}
}

func reverseStr(s string, k int) string {
	res := []byte(s)
	n := len(res)

	for start := 0; start < n; start += 2 * k {
		end := start + k - 1
		if end >= n {
			end = n - 1
		}
		reverse(res, start, end)
	}

	return string(res)
}
