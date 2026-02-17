package bootstrap_handler

import (
	"log"

	"github.com/jolotech/jolo-mars/internal/services/boostrap"


)

type BootstrapHandler struct {
	service *bootstrap_service.BootstrapService
}

func NewBootstrapHandler(service *bootstrap_service.BootstrapService) *BootstrapHandler {
	return &BootstrapHandler{service: service}
}

// ============== BOOSTRAP AND CREATE SUPER ADMIN =================

func (h *BootstrapHandler) Run() {
	log.Println("[bootstrap] starting super admin bootstrap...")

	res, err := h.service.EnsureSuperAdminFromEnvSilently()
	if err != nil {
		log.Printf("[bootstrap] error: %v\n", err)
		return
	}

	if res == nil {
		log.Println("[bootstrap] skipped: nil result")
		return
	}

	if !res.Created {
		log.Printf("[bootstrap] skipped: %s\n", res.Reason)
		return
	}

	log.Println("[bootstrap] ✅ super admin created successfully")
}
