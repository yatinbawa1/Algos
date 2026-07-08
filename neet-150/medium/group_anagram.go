package medium

func groupAnagrams(strings []string) [][]string {
	hashmap := make(map[[26]uint8][]string)

	for _, s := range strings {
		arr := [26]uint8{}

		for _, val := range s {
			arr[val-97] += 1
		}

		hashmap[arr] = append(hashmap[arr], s)
	}

	rv := [][]string{}

	for _, v := range hashmap {
		rv = append(rv, v)
	}

	return rv
}
