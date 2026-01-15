package handlers

import (
	"net/http"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jolotech/jolo-mars/internal/infrastructure/database"
)

type HealthHandler struct {
	StartTime time.Time
	Version   string
}

func (h *HealthHandler) UI(c *gin.Context) {
	dbOK := true

	sqlDB, err := database.DB.DB()
	if err != nil || sqlDB.Ping() != nil {
		dbOK = false
	}

	uptime := time.Since(h.StartTime).Round(time.Second)

	// Get memory usage
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	memUsage := m.Alloc / 1024 / 1024 // in MB

	c.HTML(http.StatusOK, "health.html", gin.H{
		"time":       time.Now().Format("02 Jan 2006, 15:04:05"),
		"uptime":     uptime.String(),
		"version":    h.Version,
		"dbOK":       dbOK,
		"memUsageMB": memUsage,
	})
}




func (h *HealthHandler) JSON(c *gin.Context) {
	dbOK := true

	sqlDB, err := database.DB.DB()
	if err != nil || sqlDB.Ping() != nil {
		dbOK = false
	}

	uptime := time.Since(h.StartTime).Round(time.Second)

	// Memory usage
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	memUsage := m.Alloc / 1024 / 1024 // in MB

	c.JSON(http.StatusOK, gin.H{
		"dbOK":       dbOK,
		"uptime":     uptime.String(),
		"memUsageMB": memUsage,
		"version":    h.Version,
		"time":       time.Now().Format("02 Jan 2006, 15:04:05"),
	})
}