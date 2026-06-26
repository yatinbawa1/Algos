package issubsequence

import "testing"

func TestIsSubSequence(t *testing.T) {
	val := [][]string{{"b", "abc"}, {"abc", "ahbgdc"}, {"axc", "ahbgdc"}}
	shouldBe := []bool{true, true, false}
	for pos, v := range val {
		result := IsSubsequence(v[0], v[1])

		if result != shouldBe[pos] {
			t.Errorf("IsSubString(%s, %s) failed: expected %v, got %v", v[0], v[1], shouldBe[pos], result)
		}
	}
}
