package service

import "net/http"

func GetReminder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.Write([]byte("{\n \"Result\" : \"GetReminder\" \n}"))
}

func GetAllReminders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.Write([]byte("{\n \"Result\" : \"GetAllReminders\" \n}"))
}

func CreateReminder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.Write([]byte("{\n \"Result\" : \"CreateReminder\" \n}"))
}

func UpdateReminder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.Write([]byte("{\n \"Result\" : \"UpdateReminder\" \n}"))
}

func DeleteReminder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.Write([]byte("{\n \"Result\" : \"DeleteReminder\" \n}"))
}
