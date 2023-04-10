package fyne

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/McCdama/govue/router"
)

func (fw *FyneApp) RunServer() {
	myWindow := fw.app.NewWindow("Server")

	content := widget.NewButtonWithIcon("Start the server", theme.SettingsIcon(), func() {
		log.Println("Server starts..")
		e := router.New()
		e.Logger.Fatal(e.Start(":8080"))
	})

	myWindow.Resize(fyne.NewSize(600, 600))
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
