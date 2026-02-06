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

// Run on app start
// func (h *BootstrapHandler) Run() {
// 	log.Println("[bootstrap] starting super admin bootstrap...")

// 	res, err := h.service.EnsureSuperAdminFromEnvSilently()
// 	if err != nil {
// 		// Your request says "fails silently" when admin exists.
// 		// But if we get a real error (DB down), log it so you know.
// 		log.Printf("[bootstrap] super admin bootstrap error: %v\n", err)
// 		return
// 	}

// 	if res != nil && res.Created {
// 		log.Println("[bootstrap] super admin created successfully")

// 		// If we generated a password, print it ONCE to server logs.
// 		// If you don't want this in logs, remove this.
// 		if res.TempPassword != "" {
// 			log.Printf("[bootstrap] TEMP PASSWORD (copy now): %s\n", res.TempPassword)
// 		} else {
// 			log.Println("[bootstrap] password came from SUPER_ADMIN_PASSWORD env")
// 		}
// 	}
// }



func (h *BootstrapHandler) Run() {
	log.Println("[bootstrap] starting super admin bootstrap...")

	res, err := h.service.EnsureSuperAdminFromEnvSilently()
	if err != nil {
		log.Printf("[bootstrap] error: %v\n", err)
		return
	}

	if res == nil {
		log.Println("[bootstrap] skipped (nil result)")
		return
	}

	if !res.Created {
		log.Printf("[bootstrap] skipped: %s\n", res.Reason)
		return
	}

	log.Println("[bootstrap] ✅ super admin created successfully")
	if res.TempPassword != "" {
		log.Printf("[bootstrap] TEMP PASSWORD (copy now): %s\n", res.TempPassword)
	} else {
		log.Println("[bootstrap] password came from SUPER_ADMIN_PASSWORD env")
	}
}
