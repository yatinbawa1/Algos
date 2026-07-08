package easy

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	hashmap := make(map[byte]int, len(s))

	for pos := range s {
		hashmap[s[pos]] = hashmap[s[pos]] + 1
		hashmap[t[pos]] = hashmap[t[pos]] - 1
	}

	for _, val := range hashmap {
		if val != 0 {
			return false
		}
	}

	return true
}
