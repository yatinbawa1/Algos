package jumpgame

// Basically in this problem
// we are at an index i
// value of that index tells us
// what max value can we jump
// our goal is to reach len(nums) - 1
// The problem is that from each index
// the value I can jump is different and a simple
// sum will not work

// At any index i, i can go any index between
// 0 and nums[i]
// {3,2,1,0,4}
// I can jump from i=0 to i=3 -> max = 3
// i=1, with value[i] => 2, i can go 0 -> i safely using max, then
// i can go i + value[i] => 1 + 2 => 3 max remains same
// i=2  with value[i] => 1, I can go to 0 -> safely using max, then
// i can go i + value[i] => 2 + 1 => 3
// And So forth at end of loop max < len(value) - 1 : Return False

// {3,2,2,0,4}
// Here I can go to i=2 using value[0]
// then from their i + value[2] = 4 -> Return true

func CanJump(nums []int) bool {

	currentReach := 0

	for i := 1; i < len(nums); i++ {

		if i > currentReach {
			return false
		}

		currentReach = max(currentReach, i+nums[i])

		if currentReach >= len(nums)-1 {
			return true
		}
	}

	return currentReach >= len(nums)-1
}
