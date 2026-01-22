package jobs

import (
	"log"

	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
)

func StartJobScheduler(db *gorm.DB) {
	c := cron.New()

	// Run every 5 minutes
	c.AddFunc("@every 5m", func() {
		if err := CleanupExpiredOTPs(db); err != nil {
			log.Println("OTP cleanup failed:", err)
		}
		log.Println("OTP cleanup completed successfully")
	})
	log.Println("Job scheduler started...")
	c.Start()
}
