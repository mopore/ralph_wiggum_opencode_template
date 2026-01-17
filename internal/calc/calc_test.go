package calc

import (
	"sync"
	"testing"
)

func TestAddition(t *testing.T) {
	e := New()
	e.Input("7")
	e.Operator("+")
	e.Input("5")
	e.Equals()
	if e.Display() != "12" {
		t.Errorf("7+5 = %s, want 12", e.Display())
	}
}

func TestSubtraction(t *testing.T) {
	e := New()
	e.Input("9")
	e.Operator("-")
	e.Input("4")
	e.Equals()
	if e.Display() != "5" {
		t.Errorf("9-4 = %s, want 5", e.Display())
	}
}

func TestMultiplication(t *testing.T) {
	e := New()
	e.Input("6")
	e.Operator("*")
	e.Input("7")
	e.Equals()
	if e.Display() != "42" {
		t.Errorf("6*7 = %s, want 42", e.Display())
	}
}

func TestDivision(t *testing.T) {
	e := New()
	e.Input("8")
	e.Operator("/")
	e.Input("2")
	e.Equals()
	if e.Display() != "4" {
		t.Errorf("8/2 = %s, want 4", e.Display())
	}
}

func TestClear(t *testing.T) {
	e := New()
	e.Input("1")
	e.Input("2")
	e.Input("3")
	e.Clear()
	if e.Display() != "0" {
		t.Errorf("Clear: %s, want 0", e.Display())
	}
}

func TestBackspace(t *testing.T) {
	e := New()
	e.Input("1")
	e.Input("2")
	e.Input("3")
	e.Backspace()
	if e.Display() != "12" {
		t.Errorf("Backspace: %s, want 12", e.Display())
	}
}

// TestConcurrentOps uses sync.WaitGroup.Go (Go 1.25 feature)
func TestConcurrentOps(t *testing.T) {
	var wg sync.WaitGroup
	wg.Go(func() {
		e := New()
		e.Input("1")
		e.Operator("+")
		e.Input("1")
		e.Equals()
		if e.Display() != "2" {
			t.Errorf("concurrent: 1+1 = %s", e.Display())
		}
	})
	wg.Go(func() {
		e := New()
		e.Input("2")
		e.Operator("*")
		e.Input("3")
		e.Equals()
		if e.Display() != "6" {
			t.Errorf("concurrent: 2*3 = %s", e.Display())
		}
	})
	wg.Wait()
}
