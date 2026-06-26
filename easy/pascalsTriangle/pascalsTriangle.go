package pascalstriangle

func Generate(n int) [][]int {

	if n == 1 {
		val := make([][]int, 1)
		val[0] = make([]int, 1)

		val[0][0] = 1

		return val
	}

	currentRow := make([]int, n)
	previousRows := Generate(n - 1)
	for i := 0; i < n; i++ {
		if i > 0 && i < n-1 {
			currentRow[i] = previousRows[n-2][i-1] + previousRows[n-2][i]
		} else {
			currentRow[i] = 1
		}
	}

	return append(previousRows, currentRow)

}

func Generate2(n int) []int {
	inSpaceSlice := make([]int, n+1) // A single slice, that will be used again and again

	// In Space Changes Make Next slice change reflect wrong value
	// there seems to be need to remember the last value
	//. 0.     1.        2.          3.           4.
	// [1] -> [1,1] -> [1,2,1] -> [1,3,3,1] -> [1,4,6,4,1]
	// when I am at Step [1,3,3,1]				  I
	//                      . x[1,2,1,0] ->  [1,3,1,0]
	//                            V                .x[1,3,(V + I),0]
	// I am this position
	// I know that whatever I am going to write is going to overwrite
	// the next value that is going to be useful for next run
	// My goal is to preserve this value for next cycle

	// While this solution works, there is a better way to do it
	// Using Backwards Setting
	// Basically Currently, I am setting everything from left to right,
	// Causing some data to be overwritten before use
	// Backwards setting actually prevents this from happening, setting
	// from left causes data to only be overwritten once that space is no longer needed
	// This is because the pascal triangle is sum of i - 1 + i elements

	for j := 0; j <= n; j++ {
		preservedValue := 1
		for i := 0; i <= j; i++ {
			if i > 0 && i <= j-1 {
				temp := inSpaceSlice[i]
				inSpaceSlice[i] = inSpaceSlice[i] + preservedValue
				preservedValue = temp
			} else {
				inSpaceSlice[i] = 1
			}
		}
	}

	return inSpaceSlice
}
