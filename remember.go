package main

import (
	"go-remember/persistence"
	"go-remember/service"
	"go-remember/types"
	"os"
	"strconv"
	//"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	var db = persistence.NewDB()
	defer persistence.Disconnect(db)

	repo := persistence.ReminderRepository{DB: db}
	port := 9999
	if p := os.Getenv("REMINDER_PORT"); p != "" {
		port, _ = strconv.Atoi(p)
	}

	c := &types.AppConfig{
		Repo: repo,
		Port: port,
	}

	service.StartWebServer(*c)

}
