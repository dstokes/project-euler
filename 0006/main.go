package main

import (
	"fmt"
	"math"
)

// Find the difference between the sum of the squares
// of the first one hundred natural numbers and the square
// of the sum.
func main() {
	var i, max, sum, squares float64
	for i, max = 1, 100; i <= max; i++ {
		sum += i
		squares += math.Pow(i, 2)
	}
	fmt.Printf("Answer: %v\n", math.Pow(sum, 2)-squares)
}
