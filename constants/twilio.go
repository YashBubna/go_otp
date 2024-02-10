package constants

import (
	"os"

	"github.com/twilio/twilio-go"
)

var TWILIO_ACCOUNT_SID string
var TWILIO_AUTH_TOKEN string
var VERIFY_SERVICE_SID string
var Client *twilio.RestClient

func InitializeTwilioClient() {
	TWILIO_ACCOUNT_SID = os.Getenv("TWILIO_ACCOUNT_SID")
	TWILIO_AUTH_TOKEN = os.Getenv("TWILIO_AUTH_TOKEN")
	VERIFY_SERVICE_SID = os.Getenv("VERIFY_SERVICE_SID")
	Client = twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: TWILIO_ACCOUNT_SID,
		Password: TWILIO_AUTH_TOKEN,
	})

}
