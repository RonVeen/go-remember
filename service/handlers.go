package service

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go-remember/persistence"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func GetReminder(w http.ResponseWriter, r *http.Request, repo *persistence.ReminderRepository) {
	w.Header().Set("content-type", "application/json")

	params := mux.Vars(r)
	paramId := params["id"]

	if paramId == "" {
		w.WriteHeader(416)
	}

	objectId, err := primitive.ObjectIDFromHex(paramId)
	if err != nil {
		w.WriteHeader(416)
	}

	reminder := repo.FindReminder(objectId)
	if reminder == nil {
		w.WriteHeader(401)
		return
	}

	json.NewEncoder(w).Encode(reminder)
}

func GetAllReminders(w http.ResponseWriter, r *http.Request, repo *persistence.ReminderRepository) {
	w.Header().Set("content-type", "application/json")
	reminders := repo.FindAllReminder()
	if reminders == nil {
		w.WriteHeader(500)
	}
	err := json.NewEncoder(w).Encode(reminders)
	if err != nil {
		w.WriteHeader(500)
	}
}

func CreateReminder(w http.ResponseWriter, r *http.Request, repo *persistence.ReminderRepository) {
	w.Header().Set("content-type", "application/json")

	var re persistence.Reminder
	if err := json.NewDecoder(r.Body).Decode(&re); err != nil {
		w.WriteHeader(500)
		return
	}
	result := repo.AddReminder(re)
	json.NewEncoder(w).Encode(&result)
}

func UpdateReminder(w http.ResponseWriter, r *http.Request, repo *persistence.ReminderRepository) {
	w.Header().Set("content-type", "application/json")

	params := mux.Vars(r)
	id := params["id"]
	if id == "" {
		w.WriteHeader(500)
	}

	var re persistence.Reminder
	if err := json.NewDecoder(r.Body).Decode(&re); err != nil {
		w.WriteHeader(500)
		return
	}

	re.UID, _ = primitive.ObjectIDFromHex(id)
	result := repo.UpdateReminder(re)
	json.NewEncoder(w).Encode(&result)
}

func DeleteReminder(w http.ResponseWriter, r *http.Request, repo *persistence.ReminderRepository) {
	w.Header().Set("content-type", "application/json")

	params := mux.Vars(r)
	uid := params["id"]
	if uid == "" {
		w.WriteHeader(416)
		return
	}

	objectId, err := primitive.ObjectIDFromHex(uid)
	if err != nil {
		w.WriteHeader(416)
		return
	}
	result := repo.DeleteReminder(objectId)

	if result {
		w.WriteHeader(204)
	} else {
		w.WriteHeader(404)
	}

}
