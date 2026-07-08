package easy

// Method 1 -> Brute Force
// Method 2 -> Memoization

func hasDuplicate(nums []int) bool {
	store := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if store[i] > 1 {
			return true
		}

		store[i] = store[i] + 1
	}

	return false
}

func hasDuplicate2(nums []int) bool {
	store := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if store[i] > 1 {
			return true
		}

		store[i] = store[i] + 1
	}

	return false
}
