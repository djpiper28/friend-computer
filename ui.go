package main

import (
	"log"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type CurrentInput struct {
	Input string
}

func (c *CurrentInput) Backspace() {
	if len(c.Input) == 0 {
		return
	}
	c.Input = c.Input[0 : len(c.Input)-1]
}

func (c *CurrentInput) AddChar(ch byte) {
	c.Input += string(ch)
}

type Ui struct {
}

func WithUi(func(Ui)) {
	playerInfo := tview.NewTextView().
		SetDynamicColors(true).
		SetText("[green]Players[white]: Testing")

	systemInformationBox := tview.NewGrid().
		SetColumns(-1).
		SetRows(-1).
		AddItem(playerInfo, 0, 0, 1, 1, 0, 0, false)

	systemInformationBox.SetBorder(true).
		SetTitle("System Information")

	commandInput := tview.NewInputField().
		SetLabel("Command").
		SetPlaceholder("'help' for assistance").
		SetDoneFunc(func(key tcell.Key) {

		})

	logBox := tview.NewGrid().
		SetColumns(-1).
		SetRows(-100, -1, 1).
		AddItem(tview.NewBox().
			SetTitle("Test").
			SetBorder(true),
			0, 0, 1, 1, 0, 0, false).
		AddItem(tview.NewBox(), 1, 0, 1, 1, 0, 0, false).
		AddItem(commandInput, 2, 0, 1, 1, 0, 0, true)

	logBox.SetBorder(true).
		SetTitle("EVENT LOG")

	root := tview.NewGrid().
		SetColumns(-3, -1).
		SetRows(0).
		AddItem(logBox, 0, 0, 1, 1, 0, 0, false).
		AddItem(systemInformationBox, 0, 1, 1, 1, 0, 0, false)

	root.SetTitle("FRIEND COMPUTER - TOP SECRET")

	app := tview.NewApplication().SetRoot(root, true).
		EnableMouse(true).
		SetFocus(commandInput)

	err := app.Run()

	if err != nil {
		log.Fatalf("Cannot start ui: %s", err)
	}
}
