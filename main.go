package main

import "fmt"

// Subtract takes two integers, minuend and subtrahend, and returns their difference.
// It returns an integer representing the result of minuend - subtrahend.
func Subtract(minuend, subtrahend int) int {
	return minuend - subtrahend
}

func main() {
	// Example usage
	result := Subtract(10, 5)
	fmt.Println("10 - 5 =", result)
}
