package stack

// IsValid reports whether every opening bracket in s is closed by the matching
// bracket in the correct order. It accepts only '(', ')', '{', '}', '[' and
// ']'. The input is not modified.
func IsValid(s string) bool {
	parentheses := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := []rune{}
	for _, val := range s {
		if val == ')' || val == ']' || val == '}' {
			if len(stack) == 0 {
				return false
			}

			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if top != parentheses[val] {
				return false
			}

		} else {
			stack = append(stack, val)
		}
	}

	if len(stack) >= 1 {
		return false
	}

	return true
}
