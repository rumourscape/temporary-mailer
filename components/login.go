package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"github.com/rumourscape/temporary-mailer/mailer"
)

func LoginForm() *widget.Form {
	email := widget.NewEntry()
	email.Validator = mailer.Emailvalidator

	password := widget.NewPasswordEntry()

	// Create the form
	form := &widget.Form{
		Items:      []*widget.FormItem{{Text: "Email", Widget: email}, {Text: "Password", Widget: password}},
		SubmitText: "Login",
	}

	return form
}

func LoginError(canvas *fyne.Canvas, text string) *widget.PopUp {
	popup := widget.NewModalPopUp(container.NewVBox(), *canvas)
	popup.Resize(fyne.NewSize(400, 50))
	popup.Move(fyne.NewPos(200, 400))
	popup.Hide()

	label := widget.NewLabel(text)
	label.Alignment = fyne.TextAlignCenter

	button := widget.Button{Text: "OK", OnTapped: func() { popup.Hide() }}

	popup.Content.(*fyne.Container).Add(label)
	popup.Content.(*fyne.Container).Add(&button)

	return popup
}
