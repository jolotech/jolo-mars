package router

import (
	// "log"
	"net/http"
	"time"

	// "github.com/gin-gonic/gin"
	// "github.com/gin-contrib/cors"
	// "github.com/jolotech/jolo-mars/internal/app/middlewares"
	"github.com/jolotech/jolo-mars/internal/app/middlewares"
	// "github.com/jolotech/jolo-mars/internal/helpers"
	"github.com/jolotech/jolo-mars/internal/app/dependencies"
    "github.com/jolotech/jolo-mars/internal/infrastructure/database"
	"github.com/jolotech/jolo-mars/config"




	"github.com/gin-gonic/gin"
)



func InitRoutes(container *dependencies.Container) *gin.Engine {
	// logger.InitLogger()
	cfg := config.LoadConfig()


	router := gin.New()

	// Middlewares can go here (logging, recovery, etc.)
	router.SetTrustedProxies(nil)
	router.Use(middlewares.CORS())

	router.Use(gin.Recovery())
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*")




	// ✅ Health check
	// router.GET("/health", func(c *gin.Context) {
	// 	helpers.SuccessResponse(c, nil, "✅ Jolo Delivery server is healthy", http.StatusOK)
	// })

	// router.GET("/health", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "health.html", gin.H{
	// 		"time": time.Now().Format("02 Jan 2006, 15:04:05"),
	// 	})
	// })

	router.GET("/health", func(c *gin.Context) {
	// DB health check
	dbOK := true
	sqlDB, err := database.DB.DB()
	if err != nil || sqlDB.Ping() != nil {
		dbOK = false
	}

	// uptime
	uptime := time.Since(appStartTime).Round(time.Second)

	c.HTML(http.StatusOK, "health.html", gin.H{
		"time":    time.Now().Format("02 Jan 2006, 15:04:05"),
		"uptime": uptime.String(),
		"version": cfg.AppVersion,
		"dbOK":    dbOK,
	})
})



// 	router.GET("/health", func(c *gin.Context) {
// 	c.Header("Content-Type", "text/html")
// 	c.String(http.StatusOK, `
// <!DOCTYPE html>
// <html lang="en">
// <head>
// 	<meta charset="UTF-8">
// 	<title>Jolo Delivery – Health Status</title>
// 	<style>
// 		body {
// 			font-family: Arial, sans-serif;
// 			background: #0f172a;
// 			color: #e5e7eb;
// 			display: flex;
// 			align-items: center;
// 			justify-content: center;
// 			height: 100vh;
// 		}
// 		.card {
// 			background: #020617;
// 			padding: 40px;
// 			border-radius: 12px;
// 			text-align: center;
// 			box-shadow: 0 10px 30px rgba(0,0,0,0.4);
// 		}
// 		.status {
// 			font-size: 22px;
// 			color: #22c55e;
// 			margin-top: 10px;
// 		}
// 		.time {
// 			margin-top: 15px;
// 			font-size: 14px;
// 			color: #94a3b8;
// 		}
// 	</style>
// </head>
// <body>
// 	<div class="card">
// 		<h1>🚀 Jolo Delivery API</h1>
// 		<p class="status">✅ Server is Healthy</p>
// 		<p class="time">Checked at: `+time.Now().Format("02 Jan 2006, 15:04:05")+`</p>
// 	</div>
// </body>
// </html>
// `)
// })


// 	router.POST("/webhook/test", func(c *gin.Context) {
//     var payload map[string]interface{}
//     if err := c.BindJSON(&payload); err != nil {
//         c.JSON(400, gin.H{"error": "invalid json"})
//         return
//     }

//     log.Println("Received webhook:", payload)

//     c.JSON(200, gin.H{
//         "message": "Webhook received successfully",
//     })
// })


    // router.GET("/topup/verify", paymentHandler.VerifyTopUp)


	// Register routes for each feature
	// routes.PartnerRoutes(router, container.PartnerHandler, container.OrderHandler, container.PartnerRepository, container.DistanceHandler, container.AuditService, container.AdminRepo, container.PaymentHandler, container.StoreHandler)
	// routes.AdminRoutes(router, container.AdminHandler, container.AdminAudit)
	// routes.CallBackRoutes(router, container.PaymentHandler)

	// 404 handler
	NotFoundHandler(router)
	return router
}
