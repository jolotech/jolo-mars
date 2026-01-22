package utils


import (
	"math/rand"
	"os"
	"strconv"
	"time"
)

func GenerateOTP() string {
	if os.Getenv("ENV") == "development" || os.Getenv("APP_MODE") == "test" {
		return "123456"
	}

	rand.Seed(time.Now().UnixNano())
	return strconv.Itoa(100000 + rand.Intn(900000))
}


func OTPWaitError(wait int) string {
	return "Please try again after " + strconv.Itoa(wait) + " seconds"
}


func IsOTPExpired(updatedAt time.Time) bool {
	return time.Since(updatedAt).Minutes() > OTPExpiryMinutes
}


func CanResendOTP(updatedAt time.Time) (bool, int) {
	elapsed := time.Since(updatedAt).Seconds()
	if elapsed < OTPIntervalSeconds {
		return false, OTPIntervalSeconds - int(elapsed)
	}
	return true, 0
}
