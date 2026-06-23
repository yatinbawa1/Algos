package lcp

// to find the longest common prefix
func Lcp(v []string) string {
	longestPrefixTillNow := v[0]

	for _, val := range v {
		shortestItem := min(len(longestPrefixTillNow), len(val))

		i := 0
		for i < shortestItem && longestPrefixTillNow[i] == val[i] {
			i++
		}

		longestPrefixTillNow = longestPrefixTillNow[:i]
	}

	return longestPrefixTillNow
}
