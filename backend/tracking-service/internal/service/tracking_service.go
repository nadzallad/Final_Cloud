package service

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"

	"tracking-service/internal/entity"
	"tracking-service/internal/repository"
)

type TrackingService struct {
	repo *repository.TrackingRepository
}

func NewTrackingService(
	repo *repository.TrackingRepository,
) *TrackingService {

	return &TrackingService{
		repo: repo,
	}
}

func (s *TrackingService) CreateTracking(
	tracking entity.Tracking,
) error {

	tracking.CreatedAt = time.Now()

	err := s.repo.Create(tracking)
	if err != nil {
		return err
	}

	notification := map[string]interface{}{
		"source":  "tracking",
		"event":   tracking.Status,
		"message": tracking.Note,
		"no_resi": tracking.NoResi,
	}

	body, err := json.Marshal(notification)
	if err != nil {
		return err
	}

	_, err = http.Post(
		"http://localhost:8088/notification",
		"application/json",
		bytes.NewBuffer(body),
	)

	if err != nil {
		return err
	}

	return nil
}

func (s *TrackingService) GetTrackingByResi(
	noResi string,
) ([]entity.Tracking, error) {

	return s.repo.GetTrackingByResi(noResi)
}