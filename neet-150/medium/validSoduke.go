package medium

const (
	ROW    = 9
	COL    = 9
	SUBBOX = 9
)

func IsValidSudoku(board [][]byte) bool {
	// First 9 Elements Are Rows, Next 9 Elements are Cols and Last 9 are SUBBOX
	hashmapArray := make([]map[byte]struct{}, (ROW + COL + SUBBOX))
	for i := 0; i < ROW+COL+SUBBOX; i++ {
		hashmapArray[i] = make(map[byte]struct{})
	}

	for rowIndex, rowVal := range board {
		for colIndex, elementVal := range rowVal {
			if elementVal != byte('.') {
				// Check Row
				if _, ok := hashmapArray[rowIndex][elementVal]; !ok {
					hashmapArray[rowIndex][elementVal] = struct{}{}
				} else {
					return false
				}

				//Check Col
				if _, ok := hashmapArray[9+colIndex][elementVal]; !ok {
					hashmapArray[9+colIndex][elementVal] = struct{}{}
				} else {
					return false
				}

				boxMapIdx := ((rowIndex/3)*3 + (colIndex / 3)) + 18
				if _, ok := hashmapArray[boxMapIdx][elementVal]; !ok {
					hashmapArray[boxMapIdx][elementVal] = struct{}{}
				} else {
					return false
				}
			}
		}
	}

	return true
}
