package repository

import (
	"context"

	"notification-service/internal/entity"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type NotificationRepository struct {
	collection *mongo.Collection
}

func NewNotificationRepository(
	collection *mongo.Collection,
) *NotificationRepository {

	return &NotificationRepository{
		collection: collection,
	}
}

func (r *NotificationRepository) Create(
	notification entity.Notification,
) error {

	_, err := r.collection.InsertOne(
		context.TODO(),
		notification,
	)

	return err
}

func (r *NotificationRepository) FindByOrderID(
	orderID string,
) ([]entity.Notification, error) {

	filter := bson.M{
		"order_id": orderID,
	}

	cursor, err := r.collection.Find(
		context.TODO(),
		filter,
	)

	if err != nil {
		return nil, err
	}

	var notifications []entity.Notification

	if err := cursor.All(
		context.TODO(),
		&notifications,
	); err != nil {
		return nil, err
	}

	return notifications, nil
}

func (r *NotificationRepository) GetAll() ([]entity.Notification, error) {

	cursor, err := r.collection.Find(
		context.TODO(),
		bson.M{},
	)

	if err != nil {
		return nil, err
	}

	var notifications []entity.Notification

	if err := cursor.All(
		context.TODO(),
		&notifications,
	); err != nil {
		return nil, err
	}

	return notifications, nil
}


func (r *NotificationRepository) FindByUserID(
	userID int,
) ([]entity.Notification, error) {

	filter := bson.M{
		"user_id": userID,
	}

	cursor, err := r.collection.Find(
		context.TODO(),
		filter,
	)

	if err != nil {
		return nil, err
	}

	var notifications []entity.Notification

	if err := cursor.All(
		context.TODO(),
		&notifications,
	); err != nil {
		return nil, err
	}

	return notifications, nil
}
