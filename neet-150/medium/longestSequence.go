package medium

func LongestConsecutive(nums []int) int {

	hashmap := make(map[int]struct{}, len(nums))

	for _, v := range nums {
		hashmap[v] = struct{}{}
	}

	globalLongest := 0

	for _, v := range nums {
		if _, ok := hashmap[v-1]; !ok {
			localLongest := 1

			for j := 1; ; j++ {
				if _, ok := hashmap[v+j]; !ok {
					break
				}
				localLongest++
			}

			globalLongest = max(localLongest, globalLongest)
		}
	}

	return globalLongest
}
