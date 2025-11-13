package main

//func reverseWords(s string) string {
//	words := []byte(s)
//	var start *int
//	end := 0
//	for i := 0; i < len(words); i++ {
//		if i+1 == len(words) && start != nil {
//			end = i
//			reverse(words, *start, end)
//			break
//		}
//		if words[i] != ' ' {
//			if start == nil {
//				a := i
//				start = &a
//			} else {
//				end = i
//			}
//			continue
//		}
//
//		if start != nil {
//			reverse(words, *start, end)
//			start = nil
//		}
//	}
//
//	return string(words)
//}

func reverseWords(s string) string {
	n := len(s)
	var res = make([]byte, n)
	for i := range s {
		res[i] = s[i]
	}

	end := 0
	for i := 0; i < n; i++ {
		if i+1 == n {
			reverse(res, i-end, i)
			break
		}
		if res[i] != ' ' {
			end++
			continue
		}

		if end != 0 {
			reverse(res, i-end, i-1)
			end = 0
		}
	}

	return string(res)
}

func reverse(word []byte, from, to int) {
	for from < to {
		word[from], word[to] = word[to], word[from]
		from++
		to--
	}
}
