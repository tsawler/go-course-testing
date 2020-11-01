package main

import (
	"errors"
	"log"
)

func main() {
	result, err := divide(100, 0)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("result of division is", result)
}

func divide(x, y int) (int, error) {
	var result int
	if y == 0 {
		return result, errors.New("cannot divide by 0")
	}
	result = x / y
	return result, nil
}
