package main

import (
	"kellyCalc/kellycalc"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func uiBase() fyne.CanvasObject {
	//Box := container.NewVBox()

	BalEnt := widget.NewEntry()
	BalEnt.PlaceHolder = "100.00"

	BalBind := binding.NewString()
	BalEnt.OnChanged = func(s string) {
		BalBind.Set(s)
		BalEnt.Bind(BalBind)
	}

	winEnt := widget.NewEntry()
	winBind := binding.NewString()
	winEnt.OnChanged = func(s string) {
		winBind.Set(s)
		winEnt.Bind(winBind)
	}
	winEnt.PlaceHolder = "100.00%"

	oddsEnt := widget.NewEntry()
	oddsBind := binding.NewString()
	oddsEnt.OnChanged = func(s string) {
		oddsBind.Set(s)
		oddsEnt.Bind(oddsBind)
	}

	lossEntry := widget.NewLabel("0%")
	box := container.NewVBox()

	form := widget.NewForm(
		widget.NewFormItem("Currren Balance", BalEnt),
		widget.NewFormItem("Win percentage", winEnt),
		widget.NewFormItem("Loss Percentage", lossEntry),
		widget.NewFormItem("Odds offered", oddsEnt),
	)

	cont := container.NewBorder(form, nil, nil, nil)
	submitButn := &widget.Button{Text: "Calculate"}
	box.Add(submitButn)
	submitButn.OnTapped = func() {

		PStr, _ := winBind.Get()
		BStr, _ := oddsBind.Get()
		P, _ := kellycalc.PCalc(PStr)
		B, _ := kellycalc.BCalc(BStr)
		f := kellycalc.Fcalc(P, B)

		fBind := binding.BindFloat(&f)
		fStr := binding.FloatToString(fBind)
		disp := widget.NewLabelWithData(fStr)
		box.Add(disp)

	}

	cont.Add(box)
	return cont
}

func main() {
	app := app.New()
	wind := app.NewWindow("Kelly Calculator")

	wind.SetContent(uiBase())
	wind.ShowAndRun()
}
