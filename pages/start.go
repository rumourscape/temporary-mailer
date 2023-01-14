package pages

import (
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"github.com/rumourscape/temporary-mailer/components"
	"github.com/rumourscape/temporary-mailer/mailer"
)

var gap = layout.NewSpacer()

func Start(win *fyne.Window) *fyne.Container {
	mainCanvas := (*win).Canvas()

	errorPopup := components.LoginError(&mainCanvas, "")

	form := components.LoginForm()
	form.Hide()
	form.OnSubmit = func() {
		email := form.Items[0].Widget.(*widget.Entry)
		password := form.Items[1].Widget.(*widget.Entry)
		login := mailer.Login(email.Text, password.Text)

		if login {
			// Login successful
			log.Println("Login successful")
			SetPage(win, "dashboard")
		} else {
			// Login failed
			log.Println("Login failed")
			errorPopup = components.LoginError(&mainCanvas, "Invalid Email or Password")
			errorPopup.Show()
		}
	}

	title := canvas.NewText("Temporary Mailer", color.White)
	title.Alignment = fyne.TextAlignCenter
	title.TextStyle.Bold = true
	title.TextSize = 30
	title.Resize(fyne.NewSize(400, 50))

	newAccount := widget.NewButton("Create a new Account", func() {
		err := mailer.NewAccount()
		if err != nil {
			errorPopup = components.LoginError(&mainCanvas, err.Error())
			errorPopup.Show()
		}
		SetPage(win, "dashboard")
	})
	newAccount.Importance = widget.HighImportance

	oldAccount := widget.NewButton("Login with an existing Account", func() {})
	oldAccount.OnTapped = func() { form.Show(); oldAccount.Hide(); newAccount.SetText("Create a new Account instead") }

	vContainer := container.NewVBox(title, gap, form, oldAccount, newAccount, gap)
	vContainer.Resize(fyne.NewSize(400, 400))
	vContainer.Move(fyne.NewPos(280, 70))

	return container.NewWithoutLayout(vContainer, errorPopup)
}
