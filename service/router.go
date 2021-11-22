package service

import (
	"github.com/gorilla/mux"
	"go-remember/persistence"
	"net/http"
)

func NewRouter(repo *persistence.ReminderRepository) *mux.Router {

	var path = "/reminder"

	router := mux.NewRouter().StrictSlash(true)
	router.Methods("GET").Path(path).HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		GetAllReminders(writer, request, repo)
	})
	router.Methods("GET").Path(path + "/{id}").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		GetReminder(w, r, repo)
	})
	router.Methods("POST").Path(path).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		CreateReminder(w, r, repo)
	})
	router.Methods("PUT").Path(path + "/{id}").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		UpdateReminder(w, r, repo)
	})
	router.Methods("DELETE").Path(path + "/{id}").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		DeleteReminder(w, r, repo)
	})

	return router
}
