package validParentheses

func isValid(s string) bool {
	// this is a stack
	// of runes
	// with LIFO Operations
	p := []rune{}
	for _, v := range s {

		// Check for each item
		// if its opening bracket
		// add it to stack
		switch v {
		case '(':
			p = append(p, ')')
		case '{':
			p = append(p, '}')
		case '[':
			p = append(p, ']')
		default:
			// if stacks len is 0
			// we know the starting is malformed
			// and we return
			if len(p) == 0 {
				return false
			} else {
				// if at any point
				// top of stack does not
				// match next coming item
				// we return false
				if v != p[len(p)-1] {
					return false
				}
			}
			p = p[:len(p)-1]
		}
	}

	return len(p) == 0
}
