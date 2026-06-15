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

	return s.repo.Create(tracking)
}