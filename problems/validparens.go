package answers

type stack struct {
	slice []string
}

func (s *stack) push(char string) {
	s.slice = append(s.slice, char)
}

func (s *stack) peek() string {
	return s.slice[len(s.slice)-1]
}

func (s *stack) pop() string {
	var str = s.peek()
	s.slice = s.slice[0 : len(s.slice)-1]
	return str
}

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	stk := new(stack)
	pairs := map[string]string{
		"(": ")",
		"[": "]",
		"{": "}",
	}

	for i := 0; i <= len(s)-1; i++ {
		if len(stk.slice) == 0 {
			stk.push(pairs[string(s[i])])
			continue
		}
		exp := stk.peek()
		if exp != string(s[i]) {
			stk.push(pairs[string(s[i])])
		} else {
			stk.pop()
		}
	}

	if len(stk.slice) == 0 {
		return true
	}
	return false
}
