package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/jolotech/jolo-mars/internal/app/handlers/auth"
	"github.com/jolotech/jolo-mars/internal/app/middlewares"
)


func AdminRoutes(
	r *gin.Engine,
	authHandler *auth.AdminAuthHandler,
	// orderHandler *partners.OrderHandler,
	// partnerRepo *repositories.PartnerRepository,
	// distanceHandler *partners.DistanceHandler,
	// auditService *service.AuditService,
	// adminRepo *repositories.AdminRepository,
	// paymentHandler *partners.PaymentHandler,
	// StoreHandler *partners.StoreHandler,
) {

		public := r.Group("/admin")
		{
			auth := public.Group("/auth")
			{
				auth.POST("/login", authHandler.Login)
				auth.POST("/setup-2fa", authHandler.Setup2FA)
				auth.Use(middlewares.AdminAuthMiddleware())
				auth.Use(middlewares.RequireAdminTokenPurpose("2fa"))
				auth.POST("/confirm-2fa", authHandler.Confirm2FA)
			}

			dash := public.Group("/dash")
			dash.Use(middlewares.AdminAuthMiddleware())
			{
			    dash.Use(middlewares.RequireAdminTokenPurpose("pwd_change"))
			    dash.PUT("/change-password", authHandler.ChangePassword)
		    }
		}
}