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

	systemInformationBox := tview.NewFrame(playerInfo).
		SetBorder(true).
		SetBackgroundColor(tcell.ColorBlack.TrueColor()).
		SetBorderColor(tcell.ColorGreen.TrueColor()).
		SetTitle("System Information")

	logBox := tview.NewFrame(tview.NewBox()).
		SetBorder(true).
		SetBackgroundColor(tcell.ColorBlack.TrueColor()).
		SetBorderColor(tcell.ColorGreen.TrueColor()).
		SetTitle("EVENT LOG")

	root := tview.NewGrid().
		SetColumns(-3, -1).
		SetRows(0).
		AddItem(logBox, 0, 0, 1, 1, 0, 0, false).
		AddItem(systemInformationBox, 0, 1, 1, 1, 0, 0, false)

	root.SetTitle("FRIEND COMPUTER - TOP SECRET").
		SetBackgroundColor(tcell.ColorBlack.TrueColor())

	err := tview.NewApplication().SetRoot(root, true).
		EnableMouse(true).
		SetFocus(root).
		Run()

	if err != nil {
		log.Fatalf("Cannot start ui: %s", err)
	}
}
