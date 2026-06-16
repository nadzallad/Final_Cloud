package entity

import "time"

type Notification struct {
	NotificationID string    `json:"notification_id" bson:"notification_id"`

	NoResi  string `json:"no_resi,omitempty" bson:"no_resi,omitempty"`

	Source string `json:"source" bson:"source"`
	Event  string `json:"event" bson:"event"`

	Message string    `json:"message" bson:"message"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}