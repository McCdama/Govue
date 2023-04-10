package fyne

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

type FyneWindow struct {
	app fyne.App
	win fyne.Window
}

func NewFyneWindow() *FyneWindow {
	return &FyneWindow{
		app: app.New(),
	}
}

func (fw *FyneWindow) RunUI() {
	myApp := app.New()
	w := myApp.NewWindow("Clock")

	clock := widget.NewLabel("")
	updateTime(clock)

	w.SetContent(clock)
	go func() {
		for range time.Tick(time.Second) {
			updateTime(clock)
		}
	}()
	w.ShowAndRun()
}

func updateTime(clock *widget.Label) {
	formatted := time.Now().Format("Time: 03:04:05")
	clock.SetText(formatted)
}

// func (fw *FyneWindow) ___RunUI() {
// 	myApp := app.New()
// 	myWindow := myApp.NewWindow("Server")

// 	content := widget.NewButtonWithIcon("Start the server", theme.SettingsIcon(), func() {
// 		log.Println("Server starts..")
// 		e := router.New()
// 		e.Logger.Fatal(e.Start(":8080"))
// 	})

// 	myWindow.Resize(fyne.NewSize(600, 600))
// 	myWindow.SetContent(content)
// 	myWindow.ShowAndRun()
// }
