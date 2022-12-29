package components

import (
	"log"

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
		CancelText: "Cancel",
		OnSubmit: func() {
			if mailer.Login(email.Text, password.Text) {
				// Login successful
				log.Println("Login successful")
			} else {
				// Login failed
				log.Println("Login failed")
			}
		},
	}

	return form
}
