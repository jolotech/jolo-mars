package guest_service

import (
	"net/http"

	// "github.com/google/uuid"
	"github.com/jolotech/jolo-mars/internal/models"
	guest_repo "github.com/jolotech/jolo-mars/internal/repository/guest"
	"github.com/jolotech/jolo-mars/types"
	// "github.com/google/uuid"
)

type GuestService struct {
	guestRepo *guest_repo.GuestRepo
}

func NewGuestService(guestRepo *guest_repo.GuestRepo) *GuestService {
	return &GuestService{guestRepo: guestRepo}
}

// ================= CREATE GUEST ===============
func (s *GuestService) GuestRequest(req types.GuestRequest) (string, any, int, error) {
	guest := &models.Guest{
		IPAddress: req.IPAddress,
		FCMToken:  req.FCMToken,
	}

	if err := s.guestRepo.CreateGuest(guest); err != nil {
		return "Something went wrong", nil, http.StatusInternalServerError, err
	}
	return "guest verified", types.GuestResponse{GuestID: guest.PublicID}, http.StatusOK, nil
}
