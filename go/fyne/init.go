package fyne

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

type FyneApp struct {
	app fyne.App
}

func NewFyne() *FyneApp {
	return &FyneApp{
		app: app.New(),
	}
}
