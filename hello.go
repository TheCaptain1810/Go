package main

import (
	"fmt"
	"math"
)

func addition(numOne int, numTwo int) int {
	return numOne + numTwo
}
func subtraction(numOne int, numTwo int) int {
	return numOne - numTwo
}
func multiplication(numOne int, numTwo int) int {
	return numOne * numTwo
}
func division(numOne int, numTwo int) int {
	if numTwo == 0 {
		panic("division by zero") // Panic for division by zero in integer division
	}
	return numOne / numTwo
}
func floatDivision(numOne, numTwo float64) float64 {
	if numTwo == 0 {
		return math.Inf(1) // Return positive infinity for division by zero
	}
	return numOne / numTwo
}

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("Let's say our two number are 10 and 3")
	fmt.Println("Addition of those numbers is: ", addition(10, 3))
	fmt.Println("Subtraction of those numbers is: ", subtraction(10, 3))
	fmt.Println("Multiplication of those numbers is: ", multiplication(10, 3))
	fmt.Println("Division of those numbers is: ", division(10, 3))
	fmt.Printf("Floating point division of 10 and 3 is: %.4f\n", floatDivision(10, 3))
}
