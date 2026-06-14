package repository

import (
	"context"

	"tracking-service/internal/entity"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type TrackingRepository struct {
	collection *mongo.Collection
}

func NewTrackingRepository(
	collection *mongo.Collection,
) *TrackingRepository {

	return &TrackingRepository{
		collection: collection,
	}
}

func (r *TrackingRepository) Create(
	tracking entity.Tracking,
) error {

	_, err := r.collection.InsertOne(
		context.TODO(),
		tracking,
	)

	return err
}