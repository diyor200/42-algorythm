package main

type Solution struct{}

const (
	open  = '('
	close = ')'
)

func (Solution) Is_valid(s string) bool {
	balance := 0

	for i := range s {
		if s[i] == open {
			balance++
		} else if s[i] == close {
			balance--
		}

		if balance < 0 {
			return false
		}
	}

	return balance == 0
}

// leetcode: https://leetcode.com/problems/valid-parentheses/description/

func (Solution) isValid(s string) bool {
	n := len(s)
	if n%2 != 0 {
		return false
	}

	stack := make([]rune, n)
	top := -1
	pairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, ch := range s {
		if ch == '(' || ch == '{' || ch == '[' {
			top++
			stack[top] = ch
		} else {
			if top == -1 || stack[top] != pairs[ch] {
				return false
			}
			top-- // pop
		}
	}

	return top == -1
}
