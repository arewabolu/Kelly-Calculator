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

	box := container.NewVBox()

	form := widget.NewForm(
		widget.NewFormItem("Currren Balance", BalEnt),
		widget.NewFormItem("Win percentage", winEnt),

		widget.NewFormItem("Odds offered", oddsEnt),
	)

	cont := container.NewBorder(form, nil, nil, nil)
	submitButn := &widget.Button{Text: "Calculate"}
	box.Add(submitButn)
	submitButn.OnTapped = func() {

		PStr, _ := winBind.Get()
		BStr, _ := oddsBind.Get()
		bal, _ := BalBind.Get()

		P, _ := kellycalc.PCalc(PStr)
		B, _ := kellycalc.BCalc(BStr)
		f := kellycalc.Fcalc(P, B)
		balance, _ := kellycalc.BalCalc(bal)

		//Ran out of names
		use := kellycalc.WithBal(balance, f)
		fBind := binding.BindFloat(&use)
		fStr := binding.FloatToStringWithFormat(fBind, "Value is: %0.2f")
		disp := widget.NewLabelWithData(fStr)
		for _, obj := range box.Objects {

			if obj == submitButn {
				continue
			}
			box.Remove(obj)

		}
		box.Add(disp)

	}

	cont.Add(box)
	return cont
}

func main() {
	app := app.New()
	wind := app.NewWindow("Kelly Calculator")
	wind.Resize(fyne.NewSize(400, 400))

	wind.SetContent(uiBase())
	wind.ShowAndRun()
}
