package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/jolotech/jolo-mars/internal/app/handlers/auth"
	// "github.com/jolotech/jolo-mars/internal/app/middlewares"
	// "github.com/jolotech/jolo-mars/internal/repository"
	// "github.com/jolotech/jolo-mars/internal/services"
)

func UserRoutes(
	r *gin.Engine,
	authHandler *auth.UserAuthHandler,
	// orderHandler *partners.OrderHandler,
	// partnerRepo *repositories.PartnerRepository,
	// distanceHandler *partners.DistanceHandler,
	// auditService *service.AuditService,
	// adminRepo *repositories.AdminRepository,
	// paymentHandler *partners.PaymentHandler,
	// StoreHandler *partners.StoreHandler,
) {
	// --------------------------
	// PUBLIC PARTNER ROUTES
	// --------------------------
	public := r.Group("/v1")
	{
		auth := public.Group("/auth")
		{
			// auth.POST("/login", authHandler.Login)
			auth.POST("/register", authHandler.Register)
			auth.POST("/guest-request", authHandler.GuestRequest)
			auth.POST("/verify-otp", authHandler.VerifyOTP)
			auth.POST("/resend-otp", authHandler.ResendOTP)
			auth.POST("/forget-password", authHandler.ForgetPassword)
			auth.PUT("/reset-password", authHandler.ResetPassword)
		}

		// --------------------------
		// PARTNER DASHBOARD ROUTES (/api/partners/u/*)
		// Protected with JWT
		// --------------------------
		// dash := public.Group("/dash")
		// dash.Use(middlewares.DashAuthMiddleware())
		// {
		//     dash.POST("/topup/init", paymentHandler.InitTopUp)
		// 	public.GET("/topup/verify", paymentHandler.VerifyTopUp)
		// 	dash.PUT("/debit/limit", paymentHandler.UpdateAutoDebitLimit)

		// 	dash.POST("/change-password", authHandler.ChangePassword)
		// 	dash.POST("/update", authHandler.UpdatePartner)
		// 	dash.PUT("/auto_debit/enable", authHandler.EnableAutodebit)
		// 	dash.POST("/store/create", StoreHandler.RegisterStore)

		// 	// Webhook logs
		// 	webhook := dash.Group("/webhook")
		// 	{
		// 		webhook.GET("/logs", authHandler.WebhookLogs)
		// 	}

		// 	// Audit logs for partners
		// 	audit := dash.Group("/audit")
		// 	{
		// 		audit.GET("/logs", authHandler.AuditTrail)
		// 	}
		// }

		// // --------------------------
		// // API KEY PROTECTED ROUTES
		// // /api/v1/*
		// // --------------------------
		// apiProtected := r.Group("/api/v1")
		// apiProtected.Use(middlewares.AuditTrailMiddleware(auditService, partnerRepo, adminRepo))
		// apiProtected.Use(middlewares.APIAuthMiddleware(partnerRepo))
		// {
		// 	order := apiProtected.Group("/orders")
		// 	{
		// 		order.POST("/create", orderHandler.CreateOrder)
		// 		// order.GET("/", orderHandler.ListOrders)
		// 		// order.GET("/:id", orderHandler.GetOrder)
		// 		// order.PUT("/:id/status", orderHandler.UpdateOrderStatus)
		// 	}

		// 	distance := apiProtected.Group("/distance")
		// 	{
		// 		distance.POST("/", distanceHandler.Distance)
		// 	}
		// }
	}


}
