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
	AppName              string


	// Twilio configuration
	SID                  string  // valid
	Token                string  // valid
	MessagingServiceSID  string  // valid
	OTPTemplate          string  // valid
	FROM                 string  // valid

	// Zoho SMTP configuration
	SMTPUser             string  // valid
	SMTPPass             string  // valid
	SMTPHost             string  // valid
	SMTPPort             string  // valid

	// Auth Token (JWT)
	AdminAuthSecret      string // valid
	AuthSecret           string // valid
	AuthExpIn            string // valid
    AuthPassExpIn        string // valid

	// Boostrap Admin
	BoostrapSuperAdmin   string // valid
	SuperAdminName       string // valid
	SuperAdminPassword   string // valid
	SuperAdminEmail      string // valid
	AdminLoginUrl        string // valid
	SupportEmail        string // valid

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
		AppName:    os.Getenv("APP_NAME"),

		// Twilio configuration
		SID:            os.Getenv("TWILIO_SID"),
		Token:          os.Getenv("TWILIO_TOKEN"),
		MessagingServiceSID:   os.Getenv("TWILIO_MESSAGING_SERVICE_SID"),
		OTPTemplate:           os.Getenv("TWILIO_OTP_TEMPLATE"),
		FROM:                  os.Getenv("TWILIO_FROM"),

		// Zoho SMTP configuration 
		SMTPUser: 	  os.Getenv("SMTP_USER"),
		SMTPPass: 	  os.Getenv("SMTP_PASSWORD"),
		SMTPHost:    os.Getenv("SMTP_HOST"),
		SMTPPort:    os.Getenv("SMTP_PORT"),

		// Auth Token (JWT)
		AdminAuthSecret: os.Getenv("ADMIN_JWT_SECRET"),
		AuthSecret: os.Getenv("JWT_SECRET"),
		AuthExpIn: os.Getenv("JWT_EXPIRES_IN"),
		AuthPassExpIn: os.Getenv("ADMIN_PASS_EXPIRES_IN"),


		// Boostrap super admin 
		BoostrapSuperAdmin: os.Getenv("BOOTSTRAP_SUPER_ADMIN"),
		SuperAdminName: os.Getenv("SUPER_ADMIN_NAME"),
		SuperAdminPassword: os.Getenv("SUPER_ADMIN_PASSWORD"),
		SuperAdminEmail: os.Getenv("SUPER_ADMIN_EMAIL"),
		AdminLoginUrl: os.Getenv("ADMIN_LOGIN_URL"),
		SupportEmail: os.Getenv("SUPPORT_EMAIL"),
	}

	return cfg
}
