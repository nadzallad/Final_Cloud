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

func (r *TrackingRepository) GetTrackingByResi(
	noResi string,
) ([]entity.Tracking, error) {

	filter := map[string]interface{}{
		"no_resi": noResi,
	}

	cursor, err := r.collection.Find(
		context.TODO(),
		filter,
	)

	if err != nil {
		return nil, err
	}

	var tracking []entity.Tracking

	if err = cursor.All(
		context.TODO(),
		&tracking,
	); err != nil {
		return nil, err
	}

	return tracking, nil
}