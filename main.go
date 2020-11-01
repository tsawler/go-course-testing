package main

import (
	"errors"
	"log"
)

// main is the entry point for the application
func main() {
	result, err := divide(100, 0)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("result of division is", result)
}

// divide takes two integers (dividend and divisor) and performs integer division. If the divisor is zero,
// it returns 0 and an error; otherwise, it returns the result of the division and nil.
func divide(x, y int) (int, error) {
	var result int
	if y == 0 {
		return result, errors.New("cannot divide by 0")
	}
	result = x / y
	return result, nil
}
