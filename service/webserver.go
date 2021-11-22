package service

import (
	"go-remember/persistence"
	"log"
	"net/http"
)

func StartWebServer(port string, r *persistence.ReminderRepository) {
	log.Println("Starting webserver on port " + port)
	err := http.ListenAndServe(":"+port, NewRouter(r))
	if err != nil {
		log.Println("Error occured starting Webserver on port " + port + ", reason: " + err.Error())
	}
}
