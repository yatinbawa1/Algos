package countbits

func CountBits(n int) []int {
	// reset knows the value of last
	// time the bits to left of current digits position
	// were empty

	reset := 1
	nums := make([]int, n+1)
	nums[0] = 0
	for i := 1; i <= n; i++ {
		if reset*2 == i {
			reset = i
		}
		nums[i] = nums[i-reset] + 1

	}

	return nums
}
