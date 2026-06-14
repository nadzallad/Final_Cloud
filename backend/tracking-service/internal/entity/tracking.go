package entity

import "time"

type Tracking struct {
	TrackingID      string    `json:"tracking_id" bson:"tracking_id"`
	NoResi          string    `json:"no_resi" bson:"no_resi"`
	Status          string    `json:"status" bson:"status"`
	CurrentLocation string    `json:"location" bson:"location"`
	Note            string    `json:"note" bson:"note"`
	CreatedAt       time.Time `json:"created_at" bson:"created_at"`
}