package removeduplicatesfromsortedarray

func RemoveDuplicates(nums []int) int {
	for i := 0; i < len(nums)-1; i++ {
		consumeIndex := 0
		for _, val := range nums[i+1:] {
			if nums[i] == val {
				consumeIndex++
			}
		}

		nums = append(nums[:i], nums[i+consumeIndex:]...)
	}

	return len(nums)
}

func RemoveDuplicateTwoPointer(nums []int) int {

	//  1 1 2 3 3 4 4 4
	//  . <

	// Step 1
	// check pos 0 and 1
	// return true
	// so skip this
	//  1 1 2 3 3 4 4 4
	//	.   <
	// Step 2
	// Not Same, so move slow pointer 1 step
	//  1 1 2 3 3 4 4 4
	//	  . <
	//. Exchange slows value with fast
	//  1 2 2 3 3 4 4 4
	//	  . <
	// Step 3
	// 1 2 2 3 4 4 4
	//	 .   <
	// Different from slow so move pointer slow up and replace
	//  1 2 2 3 4 4 4
	//	    . <
	//  1 2 3 3 4 4 4
	//      . <
	// Step 4
	//  1 2 3 3 4 4 4
	//      .   <
	// Different So  Move pointer and replace
	//  1 2 3 4 4 4 4
	//        . <
	// Step 5
	// Rest Everything is same so keep on moving
	// Fast till end
	//  1 2 3 4 4 4 4
	//        .     <
	// return slowPointer

	if len(nums) == 0 {
		return 0
	}

	sp := 0
	for i := 1; i <= len(nums)-1; i++ {
		if nums[sp] != nums[i] {
			sp += 1
			nums[sp] = nums[i]
		}
	}

	return sp + 1
}
