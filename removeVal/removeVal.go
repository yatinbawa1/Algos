package removeval

func RemoveElement(nums []int, val int) int {
	pt2 := len(nums) - 1

	i := 0
	for i <= pt2 {
		if nums[i] == val {
			nums[i] = nums[pt2]
			pt2--
		} else {
			i++
		}
	}

	return i
}
