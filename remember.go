package main

import (
	"go-remember/persistence"
	"go-remember/service"
)

func main() {
	var db = persistence.NewDB()
	defer persistence.Disconnect(db)

	repo := persistence.ReminderRepository{DB: db}
	service.StartWebServer("9999", &repo)

}
