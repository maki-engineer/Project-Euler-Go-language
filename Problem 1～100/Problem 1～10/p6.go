package main

import (
	"fmt"
	"math"
)

/*
The sum of the squares of the first ten natural numbers is, 385.
The square of the sum of the first ten natural numbers is, 3025.
Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 - 285 = 2640.
Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
*/
func Problem6() {
	a := 0
	b := 0

	for i := 1; i <= 100; i++ {
		a += int(math.Pow(float64(i), 2))
		b += i
	}

	b = int(math.Pow(float64(b), 2))

	fmt.Println(int(math.Abs(float64(a - b)))) // >> 25164150
}
