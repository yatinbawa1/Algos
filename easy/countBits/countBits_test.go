package countbits

import "testing"

func TestCountBits(t *testing.T) {
	val := []int{2}
	shouldBe := [][]int{{0, 1, 1}}
	for pos, v := range val {
		result := CountBits(v)
		for p, r := range result {
			if r != shouldBe[pos][p] {
				t.Errorf("CountBits() failed: expected %v, got %v\n", shouldBe[pos], result)
			}
		}
	}
}
