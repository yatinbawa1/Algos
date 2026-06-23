package climbstairs

// Problem Set:
// Can Climb N number of stairs by doing 1 or 2 steps
// How many distinct Patters Are their

// N Remaining Function
func nRemainingSteps(n int) int {
	// If less than 2 Stairs Are Left
	// There can only be either 1 or 0 to Climb it
	if n < 2 {
		return n
	}

	return 2 * nRemainingSteps(n-1)
}

func ClimbStairs(n int) int {
	return nRemainingSteps(n)
}
