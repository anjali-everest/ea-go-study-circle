package main

import "testing"

func TestAdd(t *testing.T){
	result := add(5, 10)
	if result != 15 {
		t.Errorf("Addition failed")
	}
}

func TestSubtract(t *testing.T){
	result := subtract(10,5)
	if result != 5 {
		t.Errorf("Subtraction failed")
	}
}

func TestMultiply(t *testing.T){
	result := multiply(5, 10)
	if result != 50 {
		t.Errorf("Multiplication failed")
	}
}

func TestDivision(t *testing.T){
	result := divide(10, 5)
	if result != 2 {
		t.Errorf("Division failed")
	}
}


func TestSquareRoot(t *testing.T){
	result := squareRoot(9)
	if result != 3 {
		t.Errorf("SquareRoot failed")
	}
}

func TestSin(t *testing.T){
	result := sin(90)
	if result != 0.89 {
		t.Errorf("Sin failed")
	}
}

func TestCos(t *testing.T){
	result := cos(90)
	if result != -0.45 {
		t.Errorf("Cos failed")
	}
}