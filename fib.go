package main

import "fmt"

func fib(which int) int {
	if which <= 0 {
		return which
	}
	firstNumber, secondNumber := 0, 1
	for i := 1; i < which; i++ {
		firstNumber, secondNumber = secondNumber, firstNumber+secondNumber
	}
	return secondNumber
}

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Printf("fib %d: %d\n", i, fib(i))
	}
}
