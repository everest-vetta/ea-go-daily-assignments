package main

import (
	"fmt"
)

func main() {
	var firstNumber, secondNumber float64
	var operator string

	fmt.Println("Enter first number")
	fmt.Scan(&firstNumber)

	fmt.Println("Enter second number")
	fmt.Scan(&secondNumber)

	fmt.Println("Enter the operator (+, -, /, %, tan, sin, cos)")
	fmt.Scan(&operator)

	switch operator {
	case "+":
		result := add(firstNumber, secondNumber)
		fmt.Println(result)
	case "-":
		result := subtract(firstNumber, secondNumber)
		fmt.Println(result)
	case "/":
		result, _ := divide(firstNumber, secondNumber)
		fmt.Println(result)
	case "*":
		result := multiply(firstNumber, secondNumber)
		fmt.Println(result)
	case "tan":
		result := tan(firstNumber)
		fmt.Println(result)
	case "sin":
		sin, _ := sincos(firstNumber)
		fmt.Println(sin)
	case "cos":
		_, cos := sincos(firstNumber)
		fmt.Println(cos)
	}
}
