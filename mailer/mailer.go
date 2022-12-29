package mailer

var client = InitClient()

func Login(email, password string) bool {
	err := getToken(email, password)
	if err != nil {
		return false
	}

	client.GetAuthToken()
	return true

}

func getToken(email, password string) error {
	_, err := client.GetAuthTokenCredentials(email, password)
	if err != nil {
		return err
	}

	return nil
}
