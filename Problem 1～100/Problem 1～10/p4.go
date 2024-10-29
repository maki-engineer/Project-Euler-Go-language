package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
A palindromic number reads the same both ways.
The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 * 99.
Find the largest palindrome made from the product of two 3-digit numbers.
*/
func Problem4() {
	result := 0

	for left := 100; left < 1000; left++ {
		for right := 100; right < 1000; right++ {
			intNum := left * right
			stringNum := strconv.Itoa(intNum)
			numberList := strings.Split(stringNum, "")
			isPalindromes := true

			for i := 0; i < len(numberList)/2; i++ {
				if numberList[i] != numberList[len(numberList)-1-i] {
					isPalindromes = false
					break
				}
			}

			if isPalindromes && (result < intNum) {
				result = intNum
			}
		}
	}

	fmt.Println(result) // >> 906609
}
