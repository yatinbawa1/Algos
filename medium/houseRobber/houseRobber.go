package houserobber

func Rob(nums []int) int {
	m, p := 0, 0

	for i := 0; i < len(nums); i++ {
		hypothetical := max(m, p+nums[i])
		p = m
		m = hypothetical
	}

	return m
}
