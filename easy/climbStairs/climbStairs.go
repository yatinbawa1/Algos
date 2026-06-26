package climbStairs

// N Remaining Function
func nRemainingSteps(n int, index map[int]int) int {
	// There is only 1 Way Forward
	if n == 0 || n == 1 {
		return 1
	}

	if val, exist := index[n]; exist {
		return val
	}

	// Check all possible ways
	singleSkip := nRemainingSteps(n-1, index)
	doubleSkip := nRemainingSteps(n-2, index)

	index[n] = singleSkip + doubleSkip
	return singleSkip + doubleSkip
}

func ClimbStairs(n int) int {
	index := make(map[int]int)
	return nRemainingSteps(n, index)
}
