package mailer

import (
	"log"

	"github.com/felixstrobel/mailtm"
)

var client = InitClient()

func Login(email, password string) bool {
	token, err := client.GetAuthTokenCredentials(email, password)
	if err != nil || token == "" {
		log.Println(err)
		return false
	}

	//log.Println(token)
	return true
}

func NewAccount() error {
	_, err := client.CreateAccount()
	if err != nil {
		log.Println("Error creating Account: " + err.Error())
		return err
	}

	_, err = client.GetAuthToken()
	if err != nil {
		log.Println("Error getting AuthToken: " + err.Error())
		return err
	}

	return nil
}

func GetMessages(page int) ([]mailtm.Message, error) {
	messages, err := client.GetMessages(page)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return messages, nil
}

func GetText(id string) string {
	message, err := client.GetMessageByID(id)
	if err != nil {
		log.Println(err)
		return ""
	}

	return message.Text
}

func Logout() {
	client = InitClient()
}

func DeleteAccount() {
	err := client.DeleteAccountByID(client.Account.ID)
	if err != nil {
		log.Println("Error deleting Account: " + err.Error())
	}
}
