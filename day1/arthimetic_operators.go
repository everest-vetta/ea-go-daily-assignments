package main

import (
	"fmt"
	"math"
)

func add(first, second float64) (result float64) {
	result = first + second
	return
}

func subtract(first, second float64) (result float64) {
	result = first - second
	return
}

type DivisonError struct {
	first  float64
	second float64
	reason string
}

func (d DivisonError) Error() string {
	return fmt.Sprintf("Cannot divide %v by %v because of %v", d.first, d.second, d.reason)
}

func divide(first, second float64) (float64, DivisonError) {
	if second == 0 {
		return 0, DivisonError{first, second, "Division by zero is not possible"}
	}
	result := first / second
	return result, DivisonError{}
}

func multiply(first, second float64) float64 {
	return first * second
}

type squareRootError struct {
	reason string
}

func squareRoot(first float64) (float64, squareRootError) {
	if first < 0 {
		return 0, squareRootError{"cannot find squareroot of negative numbers"}
	}

	return math.Sqrt(first), squareRootError{}
}
