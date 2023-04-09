package twilio

import (
	"fmt"
	"os"
	"github.com/twilio/twilio-go"
	"github.com/joho/godotenv"
)

var TwilioClient *twilio.RestClient
var TwilioPhoneNumber string

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("Error loading .env file")
	}

	accountSid := os.Getenv("ACCOUNT_SID")
	authToken := os.Getenv("AUTH_TOKEN")
	TwilioPhoneNumber = os.Getenv("PHONE_NUMBER")

	TwilioClient = twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})
}