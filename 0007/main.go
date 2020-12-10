package main

import (
	"fmt"
)

func isPrime(n int64) bool {
	if n <= 1 {
		return false
	}
	if n == 2 {
		return true
	}
	var i int64
	for i = 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// What is the 10001st prime number?
func main() {
	var n, c int64
	for c, n = 1, 1; c < 10001; {
		n += 2
		if isPrime(n) {
			c += 1
		}
	}
	fmt.Printf("Answer: %d\n", n)
}
