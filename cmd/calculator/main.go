package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"image/color"

	"mopore.org/ralph_wiggum_opencode_template/internal/calculator"
)

func main() {
	a := app.New()
	w := a.NewWindow("Calculator")
	calc := calculator.New()

	// Display with black bg and bright green text
	display := canvas.NewText(calc.Display(), color.RGBA{0, 255, 0, 255})
	display.TextSize = 32
	display.Alignment = fyne.TextAlignTrailing
	displayBg := canvas.NewRectangle(color.Black)
	displayContainer := container.NewStack(displayBg, container.NewPadded(display))

	updateDisplay := func() {
		display.Text = calc.Display()
		display.Refresh()
	}

	makeDigitBtn := func(d string) *widget.Button {
		return widget.NewButton(d, func() {
			calc.InputDigit(d)
			updateDisplay()
		})
	}

	makeOpBtn := func(op string) *widget.Button {
		return widget.NewButton(op, func() {
			calc.InputOperator(op)
			updateDisplay()
		})
	}

	btnDot := widget.NewButton(".", func() {
		calc.InputDecimal()
		updateDisplay()
	})

	btnEq := widget.NewButton("=", func() {
		calc.Calculate()
		updateDisplay()
	})

	btnC := widget.NewButton("C", func() {
		calc.Clear()
		updateDisplay()
	})

	btnBack := widget.NewButton("<-", func() {
		calc.Backspace()
		updateDisplay()
	})

	buttons := container.NewGridWithColumns(4,
		btnC, btnBack, makeOpBtn("/"), makeOpBtn("*"),
		makeDigitBtn("7"), makeDigitBtn("8"), makeDigitBtn("9"), makeOpBtn("-"),
		makeDigitBtn("4"), makeDigitBtn("5"), makeDigitBtn("6"), makeOpBtn("+"),
		makeDigitBtn("1"), makeDigitBtn("2"), makeDigitBtn("3"), btnEq,
		makeDigitBtn("0"), btnDot,
	)

	content := container.NewBorder(displayContainer, nil, nil, nil, buttons)
	w.SetContent(content)

	// Keyboard support
	w.Canvas().SetOnTypedKey(func(ev *fyne.KeyEvent) {
		switch ev.Name {
		case fyne.KeyReturn, fyne.KeyEnter:
			calc.Calculate()
		case fyne.KeyEscape:
			calc.Clear()
		case fyne.KeyBackspace:
			calc.Backspace()
		}
		updateDisplay()
	})

	w.Canvas().SetOnTypedRune(func(r rune) {
		switch r {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			calc.InputDigit(string(r))
		case '.':
			calc.InputDecimal()
		case '+', '-', '*', '/':
			calc.InputOperator(string(r))
		case '=':
			calc.Calculate()
		}
		updateDisplay()
	})

	w.Resize(fyne.NewSize(300, 400))
	w.ShowAndRun()
}
