package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config holds all environment variables
type Config struct {
	DBUser               string   // valid 
	DBPassword           string   // valid
	DBHost               string   // valid
	DBPort               string   // valid
	DBName               string   // valid
	ServerPort           string   // valid
	RedisAddr            string
	RedisPort            string
	AppSecrete           string
	PaystackSecrete      string
	PHPBaseURL           string
	CloudinaryName       string
	CloudinaryApiKey     string
	CloudinaryApiSecrete string    
	AppEnv               string    // valid values: "development", "staging", "production"
	AppVersion           string    // valid


	// Twilio configuration
	SID                  string
	Token                string
	MessagingServiceSID  string
	OTPTemplate          string
	FROM                 string
}

// LoadConfig loads .env variables or system env
func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️  No .env file found, using system environment variables")
	}

	cfg := &Config{

		// Database configuration
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBName:     os.Getenv("DB_NAME"),
		ServerPort: os.Getenv("PORT"),

		RedisAddr:  os.Getenv("REDIS_ADDR"),
		RedisPort:  os.Getenv("REDIS_PORT"),
		PaystackSecrete: os.Getenv("PAYSTACK_SECRET"),
		PHPBaseURL: os.Getenv("PHP_BASE_URL"),
		CloudinaryName: os.Getenv("CLOUDINAR_NAME"),
		CloudinaryApiKey: os.Getenv("CLOUDINAR_API_KEY"),
		CloudinaryApiSecrete: os.Getenv("CLOUDINAR_API_SECRETE"),
		AppSecrete: os.Getenv("APP_SIGNATURE_SECRET"),
		AppEnv:    os.Getenv("ENV"),
		AppVersion: os.Getenv("APP_VERSION"),

		// Twilio configuration
		SID:            os.Getenv("TWILIO_SID"),
		Token:          os.Getenv("TWILIO_TOKEN"),
		MessagingServiceSID:   os.Getenv("TWILIO_MESSAGING_SERVICE_SID"),
		OTPTemplate:           os.Getenv("TWILIO_OTP_TEMPLATE"),
		FROM:                  os.Getenv("TWILIO_FROM"),
	}

	return cfg
}
