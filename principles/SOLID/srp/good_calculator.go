package srp

// Calculator is a struct that represents a calculator.
type Calculator struct{}

// Add is a method of Calculator that adds two integers and returns the result.
func (calc *Calculator) Add(a, b int) int {
	return a + b
}

// Subtract is a method of Calculator that subtracts one integer from another and returns the result.
func (calc *Calculator) Subtract(a, b int) int {
	return a - b
}

// Multiply is a method of Calculator that multiplies two integers and returns the result.
func (calc *Calculator) Multiply(a, b int) int {
	return a * b
}

// Divide is a method of Calculator that divides one integer by another and returns the result.
func (calc *Calculator) Divide(a, b int) int {
	return a / b
}
