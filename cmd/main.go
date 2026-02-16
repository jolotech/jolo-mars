package main

import (
	

	"github.com/joho/godotenv"
	"github.com/jolotech/jolo-mars/config"
	"github.com/jolotech/jolo-mars/internal/app"
	"github.com/jolotech/jolo-mars/internal/infrastructure/database"
	"github.com/jolotech/jolo-mars/internal/infrastructure/jobs"
	"github.com/jolotech/jolo-mars/internal/app/dependencies"

)

func main() {

	var container *dependencies.Container

	// if len(os.Args) < 2 {
	// 	fmt.Println("Usage: go run cmd/main.go <Your Name>")
	// 	return
	// }
	// fmt.Printf("Hello, %s!\n", os.Args[1])


    godotenv.Load()
    config.LoadConfig()

    database.ConnectDB()
	
	container = dependencies.Init()
	container.BoostrapAdminHandler.Run()
    // redisInfra.ConnectRedis()
    // ctx := context.Background()

    // redis.ConnectRedis()

    // container := dependencies.Init()
    // container.Worker.Start(ctx)

    // ---- Cron Job ----
	// c := cron.New()

	// every 15 minutes auto debit check
	// c.AddJob("*/2 * * * *", container.AutoDebit)
	// c.Start()



    // --- Initialize Webhook System ---
    // webhookRepo := repositories.NewWebhookRepository(database.DB)
    // webhookQueue := queue.NewWebhookQueue(redisInfra.RDB, "webhook_queue")
    // webhookService := service.NewWebhookService(webhookRepo, webhookQueue)

    // w := worker.NewWebhookWorker(webhookRepo, webhookQueue, webhookService)
    // w.Start(ctx)

    // log.Println("Webhook worker started...")

    // Start background job scheduler
    jobs.StartJobScheduler(database.DB)

    // Start the main API server
    app.StartServer()
}
