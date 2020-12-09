package main

import (
	"fmt"
	"sort"
)

func isPalindrome(num int) bool {
	orig := num
	reversed := 0

	for num > 0 {
		last := num % 10
		reversed = reversed*10 + last
		num /= 10
	}
	return reversed == orig
}

// Find the largest palindrome made from the product of two 3-digit numbers.
func main() {
	var palindromes []int
	for i := 999; i > 99; i-- {
		for ii := 999; ii > 99; ii-- {
			num := i * ii
			if isPalindrome(num) {
				palindromes = append(palindromes, num)
			}
		}
	}
	sort.Ints(palindromes)
	fmt.Printf("Answer: %d\n", palindromes[len(palindromes)-1])
}
