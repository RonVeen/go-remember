package service

import (
	"github.com/gorilla/mux"
	"go-remember/types"
	"net/http"
)

func NewRouter(ac types.AppConfig) *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/reminder", func(writer http.ResponseWriter, request *http.Request) {
		GetAllReminders(writer, request, ac)
	}).Methods("GET")

	router.HandleFunc("/reminder/{uid}", func(writer http.ResponseWriter, request *http.Request) {
		GetReminder(writer, request, ac)
	}).Methods("GET")

	router.HandleFunc("/reminder", func(writer http.ResponseWriter, request *http.Request) {
		CreateReminder(writer, request, ac)
	}).Methods("POST")

	router.HandleFunc("/reminder/{uid}", func(writer http.ResponseWriter, request *http.Request) {
		UpdateReminder(writer, request, ac)
	}).Methods("PUT")

	router.HandleFunc("/reminder/{uid}", func(writer http.ResponseWriter, request *http.Request) {
		DeleteReminder(writer, request, ac)
	}).Methods("DELETE")

	return router
}
