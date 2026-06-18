package service

import (
	"time"

	"notification-service/internal/entity"
	"notification-service/internal/repository"

	"github.com/google/uuid"
)

var Clients = make(map[chan entity.Notification]bool)

type NotificationService struct {
	repo *repository.NotificationRepository
}

func NewNotificationService(
	repo *repository.NotificationRepository,
) *NotificationService {

	return &NotificationService{
		repo: repo,
	}
}

func (s *NotificationService) CreateNotification(
	notification entity.Notification,
) error {

	notification.NotificationID = uuid.New().String()
	notification.CreatedAt = time.Now()

	err := s.repo.Create(notification)

	if err != nil {
		return err
	}

	for ch := range Clients {
		ch <- notification
	}

	return nil
}

func (s *NotificationService) GetNotificationsByOrderID(
	orderID string,
) ([]entity.Notification, error) {

	return s.repo.FindByOrderID(orderID)
}

func (s *NotificationService) GetAllNotifications() (
	[]entity.Notification,
	error,
) {
	return s.repo.GetAll()
}

func (s *NotificationService) GetNotificationsByUserID(
	userID int,
) ([]entity.Notification, error) {

	return s.repo.FindByUserID(userID)
}
