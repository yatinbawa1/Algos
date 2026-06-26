package issubsequence

func IsSubsequence(s string, t string) bool {
	salt := 0

	lenS := len(s)
	lenT := len(t)

	if s == "" {
		return true
	} else if t == "" {
		return false
	}

	for i := 0; i < lenT && salt < lenS; i++ {
		if s[salt] == t[i] {
			salt++
		}
	}

	return salt == lenS
}
