package calculator_test

import (
	calculator "github.com/itsmesuraj20/github-action-automation"

	"testing"
)

func TestAdd(t *testing.T) {
	result := calculator.Add(5, 3)
	expected := 8

	if result != expected {
		t.Errorf("Add(5,3) = %d; expected %d", result, expected)
	}
}

func TestSubtract(t *testing.T) {
	result := calculator.Subtract(5, 3)
	expected := 2
	if result != expected {
		t.Errorf("Subtract(5,3) = %d; expected %d ", result, expected)
	}
}

func TestMultiply(t *testing.T) {
	result := calculator.Multiply(5, 3)
	expected := 15
	if result != expected {
		t.Errorf("Multiply(5, 3) = %d; expected %d", result, expected)
	}
}

func TestDivide(t *testing.T) {
	result := calculator.Divide(6, 3)
	expected := 2
	if result != expected {
		t.Errorf("Divide(6, 3) = %d; expected %d", result, expected)
	}

	result = calculator.Divide(6, 0)
	expected = -1
	if result != expected {
		t.Errorf("Divide(6, 0) = %d; expected %d", result, expected)
	}
}
