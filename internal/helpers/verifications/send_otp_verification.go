package otp_helpers


import "os"


func SendSMS(phone, otp string) bool {
	// you can switch providers here
	// like SMS_module or SmsGateway

	// TODO: integrate real provider
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
