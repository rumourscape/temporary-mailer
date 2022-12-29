package pages

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"github.com/rumourscape/temp-mail-app/components"
)

var vGap = layout.NewSpacer()

func Start() *fyne.Container {
	form := components.LoginForm()
	form.Hide()

	title := canvas.NewText("Temporary Mailer", color.White)
	title.Alignment = fyne.TextAlignCenter
	title.TextStyle.Bold = true
	title.TextSize = 30
	title.Resize(fyne.NewSize(400, 50))

	oldAccount := widget.NewButton("Login with an existing Account", func() { form.Show() })

	newAccount := widget.NewButton("Create a new Account", func() {})
	newAccount.Importance = widget.HighImportance

	vContainer := container.NewVBox(title, vGap, form, oldAccount, newAccount, vGap)
	vContainer.Resize(fyne.NewSize(400, 350))
	vContainer.Move(fyne.NewPos(200, 50))

	return container.NewWithoutLayout(vContainer)
}
