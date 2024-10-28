package main

import "fmt"

/*
The prime factors of 13195 are 5, 7, 13 and 29.
What is the largest prime factor of the number 600851475143?
*/
func Problem3() {
	result := 0
	targetNumber := 600851475143

	for i := 2; i <= targetNumber; i++ {
		if targetNumber/i == 1 {
			break
		}

		if targetNumber%i != 0 {
			continue
		}

		for targetNumber%i == 0 {
			targetNumber /= i
		}

		result = targetNumber
	}

	fmt.Println("Problem3:", result) // >> 6857
}
