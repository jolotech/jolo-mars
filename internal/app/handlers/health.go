package handlers

import (
	"net/http"
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

	c.HTML(http.StatusOK, "health.html", gin.H{
		"time":    time.Now().Format("02 Jan 2006, 15:04:05"),
		"uptime":  uptime.String(),
		"version": h.Version,
		"dbOK":    dbOK,
	})
}
