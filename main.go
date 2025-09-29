package main

type Solution struct{}

// T: O(n)
// S: O(n)
func (Solution) RemoveStars(word string) string {
	if len(word) == 2 {
		return ""
	}

	stack := customStack{}
	for i := range word {
		if word[i] == '*' {
			stack.pop()
		} else {
			stack.push(word[i])
		}
	}

	return string(stack.data)
}

type customStack struct {
	data []byte
	size int
}

func (s *customStack) pop() {
	s.data = s.data[:s.size-1]
	s.size--
}

func (s *customStack) push(char byte) {
	s.data = append(s.data, char)
	s.size++
}

// T: O(n)
// S: O(n)
func (Solution) RemoveStarsOptimized(word string) string {
	stack := make([]byte, 0, len(word))

	for i := range word {
		if word[i] == '*' {
			stack = stack[:len(stack)-1] // removing element from last T: O(1)
		} else {
			stack = append(stack, []byte(word)[i])
		}
	}

	return string(stack)
}
