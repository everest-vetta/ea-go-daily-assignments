package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddTwoNumbers(t *testing.T) {
	assert.Equal(t, 5.0, add(2.0, 3.0))
}

func TestAddTwoNegativeNumbers(t *testing.T) {
	assert.Equal(t, -5.0, add(-2.0, -3.0))
}

func TestSubtractTwoNumbers(t *testing.T) {
	assert.Equal(t, 5.0, subtract(10.0, 5.0))
}

func TestSubtractWithFirstNegativeNumbers(t *testing.T) {
	assert.Equal(t, -9.0, subtract(-4, 5))
}

func TestDivideTwoNumbers(t *testing.T) {
	result, _ := divide(10.0, 5.0)
	assert.Equal(t, 2.0, result)
}

func TestDivideByZero(t *testing.T) {
	_, err := divide(3, 0)

	if err.reason != "Division by zero is not possible" {
		t.Error("Divison by zero is permissible")
	}
}

func TestMultiplyTwoNumbers(t *testing.T) {
	assert.Equal(t, 10.0, multiply(2.0, 5.0))
}

func TestMultiplyNegativeAndPositiveNumber(t *testing.T) {
	assert.Equal(t, -10.0, multiply(-2.0, 5.0))
}

func TestMultiplyTwoNegativeNumbers(t *testing.T) {
	assert.Equal(t, 10.0, multiply(-2.0, -5.0))
}

func TestSquareRoot0(t *testing.T) {
	result, _ := squareRoot(0)
	assert.Equal(t, 0.0, result)
}

func TestSquareRoot4(t *testing.T) {
	result, _ := squareRoot(4)
	assert.Equal(t, 2.0, result)
}

func TestSquareRootOfNegativeNumbers(t *testing.T) {
	_, err := squareRoot(-4)
	if err.reason != "cannot find squareroot of negative numbers" {
		t.Error("square root of negative number is permissible")
	}
}
