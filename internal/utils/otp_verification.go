package utils


import (
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

func GenerateOTP() string {
	if os.Getenv("ENV") == "development" || os.Getenv("APP_MODE") == "test" {
		return "123456"
	}

	rand.Seed(time.Now().UnixNano())
	return strconv.Itoa(100000 + rand.Intn(900000))
}


// func OTPWaitError(wait int) string {
// 	return "Please try again after " + strconv.Itoa(wait) + " seconds"
// }

func OTPWaitError(wait int) string {
	if wait <= 0 {
		return "Please try again shortly"
	}

	hours := wait / 3600
	minutes := (wait % 3600) / 60
	seconds := wait % 60

	var parts []string

	if hours > 0 {
		if hours == 1 {
			parts = append(parts, "1 hour")
		} else {
			parts = append(parts, strconv.Itoa(hours)+" hours")
		}
	}

	if minutes > 0 {
		if minutes == 1 {
			parts = append(parts, "1 minute")
		} else {
			parts = append(parts, strconv.Itoa(minutes)+" minutes")
		}
	}

	if seconds > 0 && hours == 0 { 
		// Only show seconds if less than 1 hour
		if seconds == 1 {
			parts = append(parts, "1 second")
		} else {
			parts = append(parts, strconv.Itoa(seconds)+" seconds")
		}
	}

	return "Please try again after " + strings.Join(parts, " ")
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


// Generate a new TOTP key (issuer + account)
func Generate2faTOTPKey(issuer, accountName string) (*otp.Key, error) {
	return totp.Generate(totp.GenerateOpts{
		Issuer:      issuer,
		AccountName: accountName, // usually email
		Period:      30,
		SecretSize:  20,
		Digits:      otp.DigitsSix,
		Algorithm:   otp.AlgorithmSHA1,
	})
}


// func Verify2faTOTP(code, secret string) bool {
// 	// Allow +/- 1 step for clock drift
// 	return totp.ValidateCustom(code, secret, time.Now(), totp.ValidateOpts{
// 		Period:    30,
// 		Skew:      1,
// 		Digits:    otp.DigitsSix,
// 		Algorithm: otp.AlgorithmSHA1,
// 	})
// }