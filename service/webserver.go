package service

import (
	"log"
	"net/http"
)

func StartWebServer(port string) {
	log.Println("Starting webserver on port " + port)
	err := http.ListenAndServe(":"+port, NewRouter())
	if err != nil {
		log.Println("Error occured starting Webserver on port " + port + ", reason: " + err.Error())
	}
}
