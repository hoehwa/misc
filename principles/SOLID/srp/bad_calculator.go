package srp

import "fmt"

// BadCalculator is a struct that represents a calculator with multiple responsibilities.
type BadCalculator struct{}

// Add is a method of BadCalculator that adds two integers and returns the result.
func (calc *BadCalculator) Add(a, b int) int {
	return a + b
}

// Subtract is a method of BadCalculator that subtracts one integer from another and returns the result.
func (calc *BadCalculator) Subtract(a, b int) int {
	return a - b
}

// Multiply is a method of BadCalculator that multiplies two integers and returns the result.
func (calc *BadCalculator) Multiply(a, b int) int {
	return a * b
}

// Divide is a method of BadCalculator that divides one integer by another and returns the result.
func (calc *BadCalculator) Divide(a, b int) int {
	return a / b
}

// ConvertToBinary is a method of BadCalculator that converts an integer to its binary representation.
func (calc *BadCalculator) ConvertToBinary(n int) string {
	return fmt.Sprintf("%b", n)
}

// ConvertToHex is a method of BadCalculator that converts an integer to its hexadecimal representation.
func (calc *BadCalculator) ConvertToHex(n int) string {
	return fmt.Sprintf("%x", n)
}
