package main

import "fmt"

func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	} else {
		return (fib(n-1) + fib(n-2))
	}
}

func main() {
	var i, answer int
	for {
		result := fib(i)
		if result > 4000000 {
			break
		}
		if result%2 == 0 {
			answer += result
		}
		i++
	}

	fmt.Printf("Answer: %d\n", answer)
}
