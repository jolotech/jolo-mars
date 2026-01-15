package dependencies

import (
	// "github.com/jolotech/Logistic-gateway/internal/app/handlers/admin"
	// "github.com/jolotech/Logistic-gateway/internal/app/handlers/partners"
	// "github.com/jolotech/Logistic-gateway/internal/infrastructure/database"
	// "github.com/jolotech/Logistic-gateway/internal/repository"
	// "github.com/jolotech/Logistic-gateway/internal/service"
	// "github.com/jolotech/Logistic-gateway/internal/queue"
	// "github.com/jolotech/Logistic-gateway/internal/worker"
	// "github.com/jolotech/Logistic-gateway/internal/jobs"
	// "github.com/jolotech/Logistic-gateway/internal/infrastructure/redis"	
	// "github.com/jolotech/jolo-mars/config"
)

type Container struct {
	// PartnerHandler *partners.PartnerHandler
	// OrderHandler *partners.OrderHandler
	// AdminHandler   *admin.AdminHandler
	// PartnerRepository *repositories.PartnerRepository
	// DistanceHandler *partners.DistanceHandler
	// Worker *worker.WebhookWorker
	// AdminAudit *admin.AuditHandler
	// AuditService *service.AuditService
	// AdminRepo *repositories.AdminRepository
	// AutoDebit *jobs.AutoDebitJob
	// PaymentHandler *partners.PaymentHandler
	// StoreHandler *partners.StoreHandler
}


// Init initializes all repositories, services, and handlers
func Init() *Container {
	// cfg := config.LoadConfig()



	// Repositories
	// partnerRepo := repositories.NewPartnerRepository(database.DB)
	// orderRepo := repositories.NewOrderRepository(database.DB)
	// adminRepo := repositories.NewAdminRepository(database.DB)
	// webhookRepo := repositories.NewWebhookRepository(database.DB)
	// auditRepo := repositories.NewAuditRepository(database.DB)
	// storeRepo := repositories.NewStoreRepository(database.DB, cfg.PHPBaseURL)

    // queue := queue.NewWebhookQueue(redis.RDB, "Webhook")

	// Services
	// partnerService := service.NewPartnerService(partnerRepo)
	// webhookService := service.NewWebhookService(webhookRepo, queue, partnerRepo)
	// webhookManager := service.NewWebhookManager(webhookService)
	// orderService := service.NewOrderService(orderRepo, webhookManager)
	// adminService := service.NewAdminService(adminRepo)
	// distanceService := service.NewDistanceService(cfg.PHPBaseURL + "api/v1/config/distance-api")
	// auditService := service.NewAuditService(auditRepo)
	// paymentService := service.NewPaymentService(cfg.PaystackSecrete)
	// storeService := service.NewStoreService(storeRepo)


	// autoDebitJobs := jobs.NewAutoDebitJob(partnerRepo, paymentService)



	// Handlers
	// partnerHandler := partners.NewPartnerHandler(partnerService, webhookService, partnerRepo, auditService, paymentService)
	// orderHandler := partners.NewOrderHandler(orderService, distanceService, storeRepo)
	// adminHandler := admin.NewAdminHandler(adminService, adminRepo, auditService)
	// distanceHandler := partners.NewDistanceHandler(distanceService)
	// worker := worker.NewWebhookWorker(webhookRepo, queue, webhookService, partnerRepo)
	// auditHandler := admin.NewAuditHandler(auditService)
	// paymentHandler := partners.NewPaymentHandler(partnerRepo, paymentService)
	// storeHamdler := partners.NewStoreHandler(storeService, *partnerRepo, storeRepo)






	return &Container{
		// PartnerHandler: partnerHandler,
		// OrderHandler:   orderHandler,
		// AdminHandler:   adminHandler,
		// PartnerRepository: partnerRepo,
		// DistanceHandler: distanceHandler,
		// Worker : worker,
		// AdminAudit: auditHandler,
		// AuditService: auditService,
		// AdminRepo: adminRepo,
		// AutoDebit: autoDebitJobs,
		// PaymentHandler: paymentHandler,
		// StoreHandler: storeHamdler,
	}
}
