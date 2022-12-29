package mailer

import "net/mail"

// email validation
func Emailvalidator(email string) error {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return err
	}

	return nil
}
