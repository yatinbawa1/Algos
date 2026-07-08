package main

import (
	"algos/neet-150/medium"
	"fmt"
)

func main() {
	v := []int{1, 2, 2, 3, 3, 3, 4}
	k := 2

	fmt.Println(medium.TopKFrequent(v, k))
}
