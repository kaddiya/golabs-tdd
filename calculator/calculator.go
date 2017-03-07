package calculator

import (
	"errors"
)

//Add performs the addition of two numbers
func Add(a, b int) int {
	return a + b
}

//Subtract performs the subtraction of two numbers
func Subtract(a, b int) int {
	return a - b
}

//Multiply performs multiplication of two numbers
func Multiply(a, b int) int {
	return a * b
}

//Divide performs the division of two numbers
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Divisor cant be equal to 0")
	}
	return a / b, nil
}
