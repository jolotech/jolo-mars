package migrations

import (
	"fmt"
	"log"

	"github.com/jolotech/jolo-mars/internal/models"
	"gorm.io/gorm"
)

// RunAll migrations for all given tables (structs)

func RunAll(db *gorm.DB, tables []string) error {
    log.Println("🚀 Starting all migrations...")

    // Ensure migration_histories table exists
    if err := db.AutoMigrate(&models.MigrationHistory{}); err != nil {
        return fmt.Errorf("failed migration_histories migration: %v", err)
    }
    log.Println("✅ migration_histories table synced.")

    for _, table := range tables {
        log.Printf("🧩 Migrating table: %s", table)

        switch table {
        case "users":
            if err := db.AutoMigrate(&models.User{}); err != nil {
                return fmt.Errorf("failed users migration: %v", err)
            }
            log.Println("✅ users table synced.")
        case "admins":
            if err := db.AutoMigrate(&models.Admin{}); err != nil {
                return fmt.Errorf("failed admins migration: %v", err)
            }
            log.Println("✅ admins table synced.")

        case "business_settings":
            if err := db.AutoMigrate(&models.BusinessSetting{}); err != nil {
                return fmt.Errorf("failed business settings migration: %v", err)
            }
            log.Println("✅ business settings table synced.")

        case "notification_settings":
            if err := db.AutoMigrate(&models.NotificationSetting{}); err != nil {
                return fmt.Errorf("failed notification settings migration: %v", err)
            }
            log.Println("✅ notification settings table synced.")

        case "otp_verifications":
            if err := db.AutoMigrate(&models.OtpVerification{}); err != nil {
                return fmt.Errorf("failed otp verifications migration: %v", err)
            }
            log.Println("✅ otp verifications table synced.")

        case "user_notifications":
            if err := db.AutoMigrate(&models.UserNotification{}); err != nil {
                return fmt.Errorf("failed user notifications migration: %v", err)
            }
            log.Println("✅ user notifications table synced.")
        case "guests":
            if err := db.AutoMigrate(&models.Guest{}); err != nil {
                return fmt.Errorf("failed guest migration: %v", err)
            }
            log.Println("✅ guest table synced.")
        case "wallet_transactions":
            if err := db.AutoMigrate(&models.WalletTransaction{}); err != nil {
                return fmt.Errorf("failed wallet transactions migration: %v", err)
            }
            log.Println("✅ wallet transactions table synced.")
        case "carts":
            if err := db.AutoMigrate(&models.Cart{}); err != nil {
                return fmt.Errorf("failed carts migration: %v", err)
            }
            log.Println("✅ carts table synced.")
        default:
            log.Printf("⚠️ No migration defined for: %s", table)
        }
    }

    log.Println("✅ All migrations completed.")
    return nil
}
