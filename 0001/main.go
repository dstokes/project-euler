package main

import "fmt"

// Find the sum of all the multiples of 3 or 5 below 1000.
func main() {
	var sum int
	for i := 1; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	fmt.Printf("Result: %d\n", sum)
}
