package q20

// 20. Valid Parentheses
func isValid(s string) bool {
	stack := make([]uint8, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack = append(stack, s[i])
		} else if len(stack) == 0 || stack[len(stack)-1] != s[i] {
			if s[i] == ')' && (len(stack) == 0 || stack[len(stack)-1] != '(') {
				return false
			} else if s[i] == '}' && (len(stack) == 0 || stack[len(stack)-1] != '{') {
				return false
			} else if s[i] == ']' && (len(stack) == 0 || stack[len(stack)-1] != '[') {
				return false
			}
			stack = stack[0 : len(stack)-1]
		}
	}
	return len(stack) == 0
}
