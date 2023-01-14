package mailer

import "github.com/felixstrobel/mailtm"

func InitClient() *mailtm.MailClient {
	client, err := mailtm.NewMailClient()
	if err != nil {
		panic(err)
	}

	client.GetDomains()

	return client
}

func GetAddress() string {
	return client.Account.Address
}

func GetPassword() string {
	return client.Account.Password
}

func GetId() string {
	return client.Account.ID
}

func GetToken() string {
	return client.Token
}
