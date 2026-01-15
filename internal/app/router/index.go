package router

import (
	"log"
	"net/http"

	// "github.com/gin-gonic/gin"
	// "github.com/gin-contrib/cors"
	// "github.com/jolotech/jolo-mars/internal/app/middlewares"
	"github.com/jolotech/jolo-mars/internal/app/middlewares"
	"github.com/jolotech/jolo-mars/internal/helpers"



	"github.com/gin-gonic/gin"
)



func InitRoutes() *gin.Engine {
	// logger.InitLogger()

	router := gin.New()

	// Middlewares can go here (logging, recovery, etc.)
	router.SetTrustedProxies(nil)
	router.Use(middlewares.CORS())

	router.Use(gin.Recovery())
	router.Use(gin.Logger())




	// ✅ Health check
	router.GET("/health", func(c *gin.Context) {
		helpers.SuccessResponse(c, nil, "✅ Jolo Logistics API Gateway is healthy", http.StatusOK)
	})

	router.POST("/webhook/test", func(c *gin.Context) {
    var payload map[string]interface{}
    if err := c.BindJSON(&payload); err != nil {
        c.JSON(400, gin.H{"error": "invalid json"})
        return
    }

    log.Println("Received webhook:", payload)

    c.JSON(200, gin.H{
        "message": "Webhook received successfully",
    })
})


    // router.GET("/topup/verify", paymentHandler.VerifyTopUp)


	// Register routes for each feature
	// routes.PartnerRoutes(router, container.PartnerHandler, container.OrderHandler, container.PartnerRepository, container.DistanceHandler, container.AuditService, container.AdminRepo, container.PaymentHandler, container.StoreHandler)
	// routes.AdminRoutes(router, container.AdminHandler, container.AdminAudit)
	// routes.CallBackRoutes(router, container.PaymentHandler)

	// 404 handler
	NotFoundHandler(router)
	return router
}
