package main

import (
	"errors"
	"fmt"
)

var ErrDivideByZero = errors.New("can't divide by zero")

func main() {
	fmt.Println("Divide 1 by 0")
	_, err := precheckDivide(1, 0)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error: %s", err.Error()))
	}
	fmt.Println("Divide 2 by 0")
	divide(2, 0)
}


func precheckDivide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}

	return divide(a, b), nil
}

func divide(a, b int) int {
	return a / b
}