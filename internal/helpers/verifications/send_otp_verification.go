package otp_helpers


import (
	"log"
	"strings"

	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
	"github.com/jolotech/jolo-mars/config"
)


func SendSMS(phone, otp string) bool {
	cfg := config.LoadConfig() // load from DB or env

	// match: if (isset($config) && $config['status'] == 1)
	// if cfg == nil || !cfg.Status {
	// 	return false
	// }

	// replace #OTP#
	message := strings.ReplaceAll(cfg.OTPTemplate, "#OTP#", otp)

	if cfg.FROM != "" {
		message = cfg.FROM + ":\n" + message
	}

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: cfg.SID,
		Password: cfg.Token,
	})

	params := &openapi.CreateMessageParams{}
	params.SetTo(phone)
	params.SetMessagingServiceSid(cfg.MessagingServiceSID)
	params.SetBody(message)

	_, err := client.Api.CreateMessage(params)
	if err != nil {
		log.Println("twilio sms error:", err)
		return false
	}

	return true
}


func SendOTPViaFirebase(phone, otp string) bool {
	// TODO: integrate Firebase OTP sending
	return true
}


func SendEmailOTP(email, otp, name string) bool {
	if os.Getenv("APP_MODE") == "test" {
		return true
	}

	// TODO: integrate real mail service
	return true
}
