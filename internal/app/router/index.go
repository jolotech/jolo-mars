package router

import (
	// "log"
	// "net/http"
	"time"

	// "github.com/gin-gonic/gin"
	// "github.com/gin-contrib/cors"
	// "github.com/jolotech/jolo-mars/internal/app/middlewares"
	"github.com/jolotech/jolo-mars/internal/app/handlers"
	"github.com/jolotech/jolo-mars/internal/app/middlewares"
	"github.com/jolotech/jolo-mars/internal/app/router/routes"

	// "github.com/jolotech/jolo-mars/internal/helpers"
	"github.com/jolotech/jolo-mars/internal/app/dependencies"
	// "github.com/jolotech/jolo-mars/internal/infrastructure/database"
	"github.com/jolotech/jolo-mars/config"

	"github.com/gin-gonic/gin"
)




func InitRoutes(container *dependencies.Container) *gin.Engine {
	// logger.InitLogger()
	cfg := config.LoadConfig()

	healthHandler := &handlers.HealthHandler{
		StartTime: time.Now(),
		Version:   cfg.AppVersion,
	}

	router := gin.New()

	// Middlewares can go here (logging, recovery, etc.)
	router.SetTrustedProxies(nil)
	router.Use(middlewares.CORS())

	router.Use(gin.Recovery())
	router.Use(gin.Logger())
	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*.html")


     router.GET("/health", healthHandler.UI)
	 router.GET("/health/json", healthHandler.JSON)

	// Initialize user routes
	routes.UserRoutes(router, container.UserAuthHandler)
	routes.AdminRoutes(router, container.AdminAuthHandler)

	// 404 handler
	NotFoundHandler(router)
	return router
}
