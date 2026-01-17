package calculator

import "testing"

func TestAdd(t *testing.T) {
	c := New()
	c.InputDigit("7")
	c.InputOperator("+")
	c.InputDigit("5")
	c.Calculate()
	if c.Display() != "12" {
		t.Errorf("7+5 expected 12, got %s", c.Display())
	}
}

func TestSubtract(t *testing.T) {
	c := New()
	c.InputDigit("9")
	c.InputOperator("-")
	c.InputDigit("4")
	c.Calculate()
	if c.Display() != "5" {
		t.Errorf("9-4 expected 5, got %s", c.Display())
	}
}

func TestMultiply(t *testing.T) {
	c := New()
	c.InputDigit("6")
	c.InputOperator("*")
	c.InputDigit("7")
	c.Calculate()
	if c.Display() != "42" {
		t.Errorf("6*7 expected 42, got %s", c.Display())
	}
}

func TestDivide(t *testing.T) {
	c := New()
	c.InputDigit("8")
	c.InputOperator("/")
	c.InputDigit("2")
	c.Calculate()
	if c.Display() != "4" {
		t.Errorf("8/2 expected 4, got %s", c.Display())
	}
}
