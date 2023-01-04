package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"

	"github.com/rumourscape/temporary-mailer/components"
	"github.com/rumourscape/temporary-mailer/pages"
)

var windowSize = fyne.NewSize(960, 540)

func main() {
	a := app.New()

	win := a.NewWindow("Temporary Mailer")
	win.Resize(windowSize)
	win.SetPadded(false)
	win.CenterOnScreen()
	win.SetFixedSize(true)

	bg := components.BackgroundImage()
	home := pages.Start(&win)

	MainContainer := container.NewMax(bg, home)

	win.SetContent(MainContainer)

	win.ShowAndRun()
}
