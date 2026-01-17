package calc

import (
	"strconv"
)

// Engine manages calculator state and operations
type Engine struct {
	display  string
	operand  float64
	operator string
	newInput bool
}

// New creates a new calculator engine
func New() *Engine {
	return &Engine{display: "0", newInput: true}
}

// Display returns current display value
func (e *Engine) Display() string {
	return e.display
}

// Input handles digit/decimal input
func (e *Engine) Input(s string) {
	if e.newInput {
		if s == "." {
			e.display = "0."
		} else {
			e.display = s
		}
		e.newInput = false
		return
	}
	if s == "." && containsDot(e.display) {
		return
	}
	e.display += s
}

// Operator sets pending operator and stores operand
func (e *Engine) Operator(op string) {
	val, _ := strconv.ParseFloat(e.display, 64)
	if e.operator != "" && !e.newInput {
		e.compute()
		val, _ = strconv.ParseFloat(e.display, 64)
	}
	e.operand = val
	e.operator = op
	e.newInput = true
}

// Equals computes result
func (e *Engine) Equals() {
	if e.operator == "" {
		return
	}
	e.compute()
	e.operator = ""
	e.newInput = true
}

// Clear resets calculator
func (e *Engine) Clear() {
	e.display = "0"
	e.operand = 0
	e.operator = ""
	e.newInput = true
}

// Backspace removes last char
func (e *Engine) Backspace() {
	if e.newInput || len(e.display) <= 1 {
		e.display = "0"
		e.newInput = true
		return
	}
	e.display = e.display[:len(e.display)-1]
}

func (e *Engine) compute() {
	b, _ := strconv.ParseFloat(e.display, 64)
	var result float64
	switch e.operator {
	case "+":
		result = e.operand + b
	case "-":
		result = e.operand - b
	case "*":
		result = e.operand * b
	case "/":
		if b == 0 {
			e.display = "Error"
			return
		}
		result = e.operand / b
	default:
		return
	}
	e.display = formatNumber(result)
}

func containsDot(s string) bool {
	for _, c := range s {
		if c == '.' {
			return true
		}
	}
	return false
}

func formatNumber(f float64) string {
	if f == float64(int64(f)) {
		return strconv.FormatInt(int64(f), 10)
	}
	return strconv.FormatFloat(f, 'f', -1, 64)
}
