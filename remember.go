package main

import (
	"go-remember/persistence"
	"go-remember/service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	var db = persistence.NewDB()
	defer persistence.Disconnect(db)

	repo := persistence.ReminderRepository{DB: db}
	repo.FindAllReminder()
	repo.FindReminder(primitive.NewObjectID())

	service.StartWebServer("9999")

}
