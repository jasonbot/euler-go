/*

Problem 3

The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?

*/

package main

import fmt "fmt"

func isprime(number uint64) bool {
	if number % 2 == 0 {
		return false
	}
	for i := number / 2; i > 1; i -= 2 {
		if number % i == 0 {
			return false;
		}
	}
	return true;
}

func main() {

	var mybignum uint64 = 600851475143

	for i := (mybignum / 2) - 1; i > 1; i -= 2 {
		if isprime(i) && mybignum % i == 0 {
			fmt.Printf("Factor: %v\n", i)
			return
		}
	}
}
