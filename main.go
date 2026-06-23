package main

import (
	"fmt"

	"algos/climbstairs"
)

func main() {
	fmt.Println("--- Testing Climb Stairs Implementation ---")

	testCases := []int{0, 1, 2, 3, 4, 5}

	for _, n := range testCases {
		result := climbstairs.ClimbStairs(n)
		fmt.Printf("Distinct patterns to climb %d stairs: %d\n", n, result)
	}
}
