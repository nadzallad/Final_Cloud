package service

import (
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

	return nil
}

func (s *TrackingService) GetTrackingByResi(
	noResi string,
) ([]entity.Tracking, error) {

	return s.repo.GetTrackingByResi(noResi)
}