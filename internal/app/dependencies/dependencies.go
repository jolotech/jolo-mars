package dependencies

import (
	// "github.com/jolotech/Logistic-gateway/internal/app/handlers/admin"
	"github.com/jolotech/jolo-mars/internal/app/handlers/auth"
	bootstrap_handler "github.com/jolotech/jolo-mars/internal/app/handlers/boostrap"
	"github.com/jolotech/jolo-mars/internal/infrastructure/database"
	"github.com/jolotech/jolo-mars/internal/repository/boostrap"
	guest_repo "github.com/jolotech/jolo-mars/internal/repository/guest"
	admin_repo "github.com/jolotech/jolo-mars/internal/repository/admin"
	"github.com/jolotech/jolo-mars/internal/repository/user"
	guest_service "github.com/jolotech/jolo-mars/internal/services/guest"
	"github.com/jolotech/jolo-mars/internal/services/user"

	// boostrap "github.com/jolotech/jolo-mars/internal/services/admin/boostrap"
	"github.com/jolotech/jolo-mars/internal/services/boostrap"
	// "github.com/jolotech/Logistic-gateway/internal/queue"
	// "github.com/jolotech/Logistic-gateway/internal/worker"
	// "github.com/jolotech/Logistic-gateway/internal/jobs"
	// "github.com/jolotech/Logistic-gateway/internal/infrastructure/redis"
	// "github.com/jolotech/jolo-mars/config"
)

type Container struct {
	UserAuthHandler *auth.UserAuthHandler
	BoostrapAdminHandler *bootstrap_handler.BootstrapHandler
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
	guestRepo := guest_repo.NewGuestRepo(database.DB)
	userMainRepo := user_repository.NewUserMainRepository(database.DB, guestRepo)
	userAuthRepo := user_repository.NewUserAuthRepository(database.DB, userMainRepo)
	adminAuthRepo := admin_repo.NewAdminAuthRepo(database.DB)
	adminMainRepo := admin_repo.NewAdminMainRepository(database.DB)
	bootstrapAdminRepo := admin_repository.NewAdminBoostrapRepository(database.DB)
	
	// orderRepo := repositories.NewOrderRepository(database.DB)
	// adminRepo := repositories.NewAdminRepository(database.DB)
	// webhookRepo := repositories.NewWebhookRepository(database.DB)
	// auditRepo := repositories.NewAuditRepository(database.DB)d
	// storeRepo := repositories.NewStoreRepository(database.DB, cfg.PHPBaseURL)

    // queue := queue.NewWebhookQueue(redis.RDB, "Webhook")

	// Services
	userAuthService := services.NewAuthService(userAuthRepo, userMainRepo, adminMainRepo, database.DB)
	guestService:= guest_service.NewGuestService(guestRepo)
	bootstrapService := bootstrap_service.NewBootstrapService(adminAuthRepo, bootstrapAdminRepo)

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
	userAuthHandler := auth.NewUserAuhHandler(userAuthService, guestService)
	bootstrapHandler := bootstrap_handler.NewBootstrapHandler(bootstrapService)

	// orderHandler := partners.NewOrderHandler(orderService, distanceService, storeRepo)
	// adminHandler := admin.NewAdminHandler(adminService, adminRepo, auditService)
	// distanceHandler := partners.NewDistanceHandler(distanceService)
	// worker := worker.NewWebhookWorker(webhookRepo, queue, webhookService, partnerRepo)
	// auditHandler := admin.NewAuditHandler(auditService)
	// paymentHandler := partners.NewPaymentHandler(partnerRepo, paymentService)
	// storeHamdler := partners.NewStoreHandler(storeService, *partnerRepo, storeRepo)






	return &Container{
		UserAuthHandler: userAuthHandler,
		BoostrapAdminHandler: bootstrapHandler,

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
