package utils


import (
	"math/rand"
	"os"
	"strconv"
	"time"
)

func GenerateOTP() string {
	if os.Getenv("APP_MODE") == "test" {
		return "123456"
	}

	rand.Seed(time.Now().UnixNano())
	return strconv.Itoa(100000 + rand.Intn(900000))
}


func OTPWaitError(wait int) string {
	return "Please try again after " + strconv.Itoa(wait) + " seconds"
}