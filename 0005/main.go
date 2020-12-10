package main

import "fmt"

// What is the smallest positive number that is evenly
// divisible by all of the numbers from 1 to 20?
func main() {
	for num := 0; ; {
		num += 20
		divisible := true
		for i := 2; i <= 20; i++ {
			if num%i != 0 {
				divisible = false
			}
		}
		if divisible {
			fmt.Printf("Answer: %d\n", num)
			break
		}
	}
}
