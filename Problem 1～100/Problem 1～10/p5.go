package main

import (
	"fmt"
	"math/big"
)

/*
2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
*/
func gcd(a, b int64) int64 {
	return big.NewInt(a).GCD(nil, nil, big.NewInt(a), big.NewInt(b)).Int64()
}

func lcm(a, b int64) int64 {
	return (a * b) / gcd(a, b)
}

func Problem5() {
	result := int64(1)

	for i := int64(2); i <= 20; i++ {
		result = lcm(result, i)
	}

	fmt.Println(result) // >> 232792560
}
