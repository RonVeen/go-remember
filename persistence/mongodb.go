package persistence

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

const mongoUrl = "MONGO_URL"

func NewDB() *mongo.Database {
	mongoUrl := os.Getenv(mongoUrl)
	if mongoUrl == "" {
		mongoUrl = "mongodb://mongoadmin:secret@localhost:27017"
	}
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoUrl))
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
