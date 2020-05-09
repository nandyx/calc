package main

import "testing"

func TestAddSuccess(t *testing.T) {
	var values []string = []string{"2", "3"}
	expected := 5
	actual, _ := add(values)

	if actual != expected {
		t.Errorf("TestAddSuccess. Actual:  %v, Expected:  %v", actual, expected)
	}

}
func TestCalcFailParse(t *testing.T) {
	input := "2+p"
	operator := "+"
	expected := "Error: Error de sintaxis"
	actual := calc(input, operator)

	if actual != expected {
		t.Errorf("Fail TestCalcFailParse. Actual %v, Expected %v", actual, expected)
	}
}
func TestCalcFailOperator(t *testing.T) {
	input := "2&3"
	operator := "+"
	expected := "La operación no es una adición"
	actual := calc(input, operator)

	if actual != expected {
		t.Errorf("Fail TestCalcFailOperator. Actual %v, Expected %v", actual, expected)
	}
}
