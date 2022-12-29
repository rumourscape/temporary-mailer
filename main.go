package main

import (
	_ "embed"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"

	"github.com/rumourscape/temp-mail-app/components"
	"github.com/rumourscape/temp-mail-app/pages"
)

var windowSize = fyne.NewSize(800, 450)

func main() {
	a := app.New()

	w := a.NewWindow("Temporary Mailer")
	w.Resize(windowSize)
	w.SetPadded(false)
	w.CenterOnScreen()
	w.SetFixedSize(true)

	bg := components.BackgroundImage()
	home := pages.Start()

	MainContainer := container.NewMax(bg, home)

	w.SetContent(MainContainer)
	//get the size of the window
	//w.Canvas().Size()

	w.ShowAndRun()
}
