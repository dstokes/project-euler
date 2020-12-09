package main

import (
	"fmt"
	"math"
)

// What is the largest prime factor of the number 600851475143 ?
func main() {
	input := 600851475143

	for {
		if input%2 == 0 {
			input /= 2
			continue
		}
		break
	}

	for i := 3; float64(i) <= math.Sqrt(float64(input)); i = i + 2 {
		for {
			if input%i == 0 {
				input /= i
				continue
			}
			break
		}
	}

	fmt.Printf("Answer: %d\n", input)
}
