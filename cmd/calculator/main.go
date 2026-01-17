package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"mopore.org/ralph_wiggum_opencode_template/internal/calc"
)

func main() {
	a := app.New()
	w := a.NewWindow("Calculator")
	engine := calc.New()

	// Display with black bg, bright green text
	display := canvas.NewText(engine.Display(), color.RGBA{0, 255, 0, 255})
	display.TextSize = 32
	display.Alignment = fyne.TextAlignTrailing
	displayBg := canvas.NewRectangle(color.Black)
	displayContainer := container.NewStack(displayBg, container.NewPadded(display))

	update := func() {
		display.Text = engine.Display()
		display.Refresh()
	}

	btn := func(label string, action func()) *widget.Button {
		return widget.NewButton(label, action)
	}

	digitBtn := func(d string) *widget.Button {
		return btn(d, func() { engine.Input(d); update() })
	}

	opBtn := func(op string) *widget.Button {
		return btn(op, func() { engine.Operator(op); update() })
	}

	buttons := container.NewGridWithColumns(4,
		digitBtn("7"), digitBtn("8"), digitBtn("9"), opBtn("/"),
		digitBtn("4"), digitBtn("5"), digitBtn("6"), opBtn("*"),
		digitBtn("1"), digitBtn("2"), digitBtn("3"), opBtn("-"),
		digitBtn("0"), digitBtn("."), btn("=", func() { engine.Equals(); update() }), opBtn("+"),
		btn("C", func() { engine.Clear(); update() }), btn("←", func() { engine.Backspace(); update() }),
	)

	content := container.NewBorder(displayContainer, nil, nil, nil, buttons)
	w.SetContent(content)

	// Keyboard support
	w.Canvas().SetOnTypedKey(func(ev *fyne.KeyEvent) {
		switch ev.Name {
		case fyne.KeyReturn, fyne.KeyEnter:
			engine.Equals()
		case fyne.KeyEscape:
			engine.Clear()
		case fyne.KeyBackspace:
			engine.Backspace()
		}
		update()
	})
	w.Canvas().SetOnTypedRune(func(r rune) {
		switch r {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '.':
			engine.Input(string(r))
		case '+', '-', '*', '/':
			engine.Operator(string(r))
		case '=':
			engine.Equals()
		}
		update()
	})

	w.Resize(fyne.NewSize(300, 400))
	w.ShowAndRun()
}
