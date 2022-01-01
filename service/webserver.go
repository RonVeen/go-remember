package service

import (
	"go-remember/types"
	"log"
	"net/http"
	"strconv"
)

func StartWebServer(ac types.AppConfig) {
	log.Println("Starting webserver on port " + strconv.Itoa(ac.Port))
	err := http.ListenAndServe(":"+strconv.Itoa(ac.Port), NewRouter(ac))
	if err != nil {
		log.Println("Error occured starting Webserver on port " + strconv.Itoa(ac.Port) + ", reason: " + err.Error())
	}
}
