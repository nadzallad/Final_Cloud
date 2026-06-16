package config

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var NotificationCollection *mongo.Collection

func ConnectDB() {

	mongoURI := os.Getenv("MONGO_URI")

	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27018"
	}

	client, err := mongo.Connect(
		options.Client().ApplyURI(mongoURI),
	)

	if err != nil {
		panic(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		panic(err)
	}

	db := client.Database("notification_db")

	NotificationCollection =
		db.Collection("notifications")
}