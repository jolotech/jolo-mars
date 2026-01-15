package migrations

import (
	"fmt"
	"log"

	"github.com/jolotech/jolo-mars/internal/domain"
	"gorm.io/gorm"
)

// RunAll migrations for all given tables (structs)

func RunAll(db *gorm.DB, tables []string) error {
    log.Println("🚀 Starting all migrations...")

    // Ensure migration_histories table exists
    if err := db.AutoMigrate(&domain.MigrationHistory{}); err != nil {
        return fmt.Errorf("failed migration_histories migration: %v", err)
    }
    log.Println("✅ migration_histories table synced.")

    for _, table := range tables {
        log.Printf("🧩 Migrating table: %s", table)

        switch table {
        case "users":
            if err := db.AutoMigrate(&domain.User{}); err != nil {
                return fmt.Errorf("failed users migration: %v", err)
            }
            log.Println("✅ users table synced.")
        case "admins", "admin":
            if err := db.AutoMigrate(&domain.Admin{}); err != nil {
                return fmt.Errorf("failed admins migration: %v", err)
            }
            log.Println("✅ admins table synced.")

        // case "orders":
        //     if err := db.AutoMigrate(&domain.Order{}); err != nil {
        //         return fmt.Errorf("failed orders migration: %v", err)
        //     }
        //     log.Println("✅ orders table synced.")

        // case "webhook_events":
        //     if err := db.AutoMigrate(&domain.WebhookEvent{}); err != nil {
        //         return fmt.Errorf("failed webhook_events migration: %v", err)
        //     }
        //     log.Println("✅ webhook_events table synced.")

        // case "webhook_retry_logs":
        //     if err := db.AutoMigrate(&domain.WebhookRetryLog{}); err != nil {
        //         return fmt.Errorf("failed webhook_retry_logs migration: %v", err)
        //     }
        //     log.Println("✅ webhook_retry_logs table synced.")

        // case "audit_trails":
        //     if err := db.AutoMigrate(&domain.AuditTrail{}); err != nil {
        //         return fmt.Errorf("failed audit_trails migration: %v", err)
        //     }
        //     log.Println("✅ audit_trails table synced.")

        // case "stores":
        //     if err := db.AutoMigrate(&domain.Store{}); err != nil {
        //         return fmt.Errorf("failed stores migration: %v", err)
        //     }
        //     log.Println("✅ stores table synced.")

        default:
            log.Printf("⚠️ No migration defined for: %s", table)
        }
    }

    log.Println("✅ All migrations completed.")
    return nil
}
