package easy

func twoSum(nums []int, target int) []int {
	hashmap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		difference := nums[i] - target
		if _, ok := hashmap[difference]; ok {
			return []int{hashmap[difference], i}
		} else {
			hashmap[nums[i]] = i
		}
	}

	return []int{-1, -1}
}
