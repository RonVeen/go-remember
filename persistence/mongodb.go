package persistence

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func NewDB() *mongo.Database {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	database := client.Database("remember")
	return database
}

func Disconnect(db *mongo.Database) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	defer db.Client().Disconnect(ctx)
}
