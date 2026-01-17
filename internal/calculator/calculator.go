package calculator

import (
	"strconv"
	"strings"
)

// Calculator holds the state for a basic calculator
type Calculator struct {
	display  string
	operand  float64
	operator string
	newInput bool
}

// New creates a new Calculator
func New() *Calculator {
	return &Calculator{
		display:  "0",
		newInput: true,
	}
}

// Display returns the current display value
func (c *Calculator) Display() string {
	return c.display
}

// InputDigit appends a digit to the display
func (c *Calculator) InputDigit(d string) {
	if c.newInput {
		c.display = d
		c.newInput = false
	} else {
		if c.display == "0" && d != "." {
			c.display = d
		} else {
			c.display += d
		}
	}
}

// InputDecimal adds a decimal point if not present
func (c *Calculator) InputDecimal() {
	if c.newInput {
		c.display = "0."
		c.newInput = false
		return
	}
	if !strings.Contains(c.display, ".") {
		c.display += "."
	}
}

// InputOperator sets the operator and stores the current operand
func (c *Calculator) InputOperator(op string) {
	val, _ := strconv.ParseFloat(c.display, 64)
	if c.operator != "" && !c.newInput {
		c.operand = c.calculate(c.operand, val, c.operator)
		c.display = formatResult(c.operand)
	} else {
		c.operand = val
	}
	c.operator = op
	c.newInput = true
}

// Calculate performs the pending operation
func (c *Calculator) Calculate() {
	if c.operator == "" {
		return
	}
	val, _ := strconv.ParseFloat(c.display, 64)
	result := c.calculate(c.operand, val, c.operator)
	c.display = formatResult(result)
	c.operator = ""
	c.newInput = true
}

// Clear resets the calculator
func (c *Calculator) Clear() {
	c.display = "0"
	c.operand = 0
	c.operator = ""
	c.newInput = true
}

// Backspace removes the last character
func (c *Calculator) Backspace() {
	if c.newInput || len(c.display) <= 1 {
		c.display = "0"
		c.newInput = true
		return
	}
	c.display = c.display[:len(c.display)-1]
	if c.display == "" || c.display == "-" {
		c.display = "0"
	}
}

func (c *Calculator) calculate(a, b float64, op string) float64 {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		if b == 0 {
			return 0 // simple handling
		}
		return a / b
	}
	return b
}

func formatResult(val float64) string {
	if val == float64(int64(val)) {
		return strconv.FormatInt(int64(val), 10)
	}
	return strconv.FormatFloat(val, 'f', -1, 64)
}
