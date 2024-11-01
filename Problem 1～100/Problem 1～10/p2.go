package main

import "fmt"

/*
Each new term in the Fibonacci sequence is generated by adding the previous two terms.
By starting with 1 and 2, the first 10 terms will be: 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...
By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms.
*/
func Problem2() {
	fibonacciNumberList := []int{1, 2, 3, 5, 8, 13, 21, 34, 55, 89}
	result := 44

	for fibonacciNumberList[len(fibonacciNumberList)-1] <= 4_000_000 {
		nextFibonacciNumber := fibonacciNumberList[len(fibonacciNumberList)-1] + fibonacciNumberList[len(fibonacciNumberList)-2]
		fibonacciNumberList = append(fibonacciNumberList, nextFibonacciNumber)

		if nextFibonacciNumber%2 == 0 {
			result += nextFibonacciNumber
		}
	}

	fmt.Println("Problem2:", result) // >> 4613732
}
