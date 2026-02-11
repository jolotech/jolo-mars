package app

import (
	"log"
	"net/http"
	"os"

	"github.com/jolotech/jolo-mars/internal/app/dependencies"
	"github.com/jolotech/jolo-mars/internal/app/router"
)

func StartServer() {
	container := dependencies.Init()

	r := router.InitRoutes(container)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3456"
	}

	log.Printf("🚀 Server running on :%s\n", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
}