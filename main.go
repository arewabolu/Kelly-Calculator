package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func uiBase() fyne.CanvasObject {
	Box := container.NewVBox()

	BalEnt := widget.NewEntry()
	winEnt := widget.NewEntry()
	oddsEnt := widget.NewEntry()
	lossEntry := widget.NewLabel("")

	form := widget.NewForm(
		widget.NewFormItem("Currren Balance", BalEnt),
		widget.NewFormItem("Win percentage", winEnt),
		widget.NewFormItem("Loss Percentage", lossEntry),
		widget.NewFormItem("Odds offered", oddsEnt),
	)
	Box.Add(form)
	cont := container.NewMax(container.NewBorder(nil, nil, nil, nil, form))
	return cont
}

func main() {
	app := app.New()
	wind := app.NewWindow("Kelly Calculator")

	wind.SetContent(uiBase())
	wind.ShowAndRun()
}
