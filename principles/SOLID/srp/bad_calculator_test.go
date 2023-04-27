package srp_test

import (
	"fmt"
	"testing"

	"github.com/mindulle/misc/principles/SOLID/srp"
)

// ExampleBadCalculator_Add demonstrates how to use the BadCalculator.Add method.
func ExampleBadCalculator_Add() {
	calc := &srp.Calculator{}
	result := calc.Add(1, 2)
	fmt.Println(result)
	// Output: 3
}

// ExampleBadCalculator_Subtract demonstrates how to use the BadCalculator.Subtract method.
func ExampleBadCalculator_Subtract() {
	calc := &srp.Calculator{}
	result := calc.Subtract(3, 2)
	fmt.Println(result)
	// Output: 1
}

// ExampleBadCalculator_Multiply demonstrates how to use the BadCalculator.Multiply method.
func ExampleBadCalculator_Multiply() {
	calc := &srp.Calculator{}
	result := calc.Multiply(2, 3)
	fmt.Println(result)
	// Output: 6
}

// ExampleBadCalculator_Divide demonstrates how to use the BadCalculator.Divide method.
func ExampleBadCalculator_Divide() {
	calc := &srp.Calculator{}
	result := calc.Divide(6, 2)
	fmt.Println(result)
	// Output: 3
}

func TestBadCalculator(t *testing.T) {
	calc := &srp.Calculator{}

	t.Run("TestAdd", func(t *testing.T) {
		result := calc.Add(1, 2)

		if result != 3 {
			t.Errorf("BadCalculator.Add(1, 2) = %d; want 3", result)
		}
	})

	t.Run("TestSubtract", func(t *testing.T) {
		result := calc.Subtract(3, 2)

		if result != 1 {
			t.Errorf("BadCalculator.Subtract(3, 2) = %d; want 1", result)
		}
	})

	t.Run("TestMultiply", func(t *testing.T) {
		result := calc.Multiply(2, 3)

		if result != 6 {
			t.Errorf("BadCalculator.Multiply(2, 3) = %d; want 6", result)
		}
	})

	t.Run("TestDivide", func(t *testing.T) {
		result := calc.Divide(6, 2)

		if result != 3 {
			t.Errorf("BadCalculator.Divide(6, 2) = %d; want 3", result)
		}
	})
}
