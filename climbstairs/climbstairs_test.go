package climbstairs_test

import (
	"algos/climbstairs"
	"testing"
)

// TestClimbStairs runs a comprehensive suite of table-driven tests
func TestClimbStairs(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  int
	}{
		{name: "Zero Stairs", input: 0, want: 0},
		{name: "One Stair", input: 1, want: 1},
		{name: "Two Stairs", input: 2, want: 2},
		{name: "Three Stairs", input: 3, want: 4},
		{name: "Four Stairs", input: 4, want: 8},
		{name: "Five Stairs", input: 5, want: 16},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := climbstairs.ClimbStairs(tt.input)
			if got != tt.want {
				t.Errorf("ClimbStairs(%d) = %d; want %d", tt.input, got, tt.want)
			}
		})
	}
}

// BenchmarkClimbStairs allows you to measure the execution performance
func BenchmarkClimbStairs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = climbstairs.ClimbStairs(10)
	}
}
