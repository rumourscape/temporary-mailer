package mailer

import "github.com/felixstrobel/mailtm"

func InitClient() *mailtm.MailClient {
	client, err := mailtm.NewMailClient()
	if err != nil {
		panic(err)
	}
	return client
}
