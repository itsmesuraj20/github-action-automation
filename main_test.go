package calculator_test

import (
	calculator "github-actions-automation"

	"testing"

	"https://github.com/itsmesuraj20/github-action-automation/calculator"
)

func TestAdd(t *testing.T){
	result := calculator.Add(5,3)
	expected := 8

	if result != expected { 
		t.Errorf("Add(5,3) = %d; expected %d" )
	}
}