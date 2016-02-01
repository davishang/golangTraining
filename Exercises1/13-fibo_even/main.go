package main

import "fmt"

func EvenFibonacciSum(x int64) (int64, error) {
	var (
		sum int64 = 0 // Sum of even Fibonacci numbers
		Fi int64 = 2 // Current Fibonacci number
		Fi_prev int64 = 1 // Previous Fibonacci number
		err error = nil // Initialize error variable
	)
	for Fi < x {
		if ( Fi % 2 == 0 ) { //If the number is even than
			sum += Fi // add number to sum
		}
		Fi = Fi + Fi_prev // Calculate next number in sequence
		Fi_prev = Fi - Fi_prev // Calculate previous number in sequence
	}
	return sum, err
}

func main() {
	var (
		x int64 = 4000000 // Limit for Fibonacci number sequence
		sum, err = EvenFibonacciSum(x)
	)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Sum of Even Fibonacci numbers less than",x,"is",sum)
	}
}
