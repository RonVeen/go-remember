package service

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go-remember/persistence"
	"go-remember/reminders"
	"go-remember/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"strconv"
)

const UID = "uid"

func GetReminder(w http.ResponseWriter, r *http.Request, c types.AppConfig) {
	id := mux.Vars(r)[UID]
	orgR, success := FindReminder(id, c, w)
	if success == false {
		return
	}
	respondWithJSONReminder(w, 200, *orgR)
}

func GetAllReminders(w http.ResponseWriter, r *http.Request, c types.AppConfig) {
	var result []persistence.Reminder
	completed := r.URL.Query().Get("completed")
	switch completed {
	case "true":
		result = c.Repo.FindCompleted()
	case "false":
		result = c.Repo.FindUncompleted()
	case "":
		result = c.Repo.FindAllReminder()
	}
	respondWithJSONReminders(w, 200, result)
}

func CreateReminder(w http.ResponseWriter, r *http.Request, c types.AppConfig) {
	var re persistence.Reminder
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&re); err != nil {
		respondWithError(w, 400, err.Error())
		return
	}
	pr := c.Repo.AddReminder(re)
	respondWithJSONReminder(w, 201, *pr)
}

func UpdateReminder(w http.ResponseWriter, r *http.Request, c types.AppConfig) {
	id := mux.Vars(r)[UID]
	orgR, success := FindReminder(id, c, w)
	if success == false {
		return
	}

	var ur persistence.Reminder
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&ur); err != nil {
		respondWithError(w, 416, "Invalid body content")
		return
	}

	if ur.Title != "" {
		orgR.Title = ur.Title
	}
	if ur.Comment != "" {
		orgR.Comment = ur.Comment
	}
	orgR.Completed = ur.Completed

	orgR = c.Repo.UpdateReminder(*orgR)
	respondWithJSONReminder(w, 200, *orgR)

}

func DeleteReminder(w http.ResponseWriter, r *http.Request, c types.AppConfig) {
	id := mux.Vars(r)[UID]
	if id == "" {
		respondWithError(w, 416, "Missing 'UID' parameter")
		return
	}

	uid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		respondWithError(w, 416, "Invalid UID parameter value")
		return
	}
	success, count := c.Repo.DeleteReminder(uid)
	if success == false {
		respondWithError(w, 416, "Bad request")
		return
	}
	if count == 0 {
		respondWithJSON(w, 404, "Reminder not found for UID="+id)
	} else {
		respondWithJSON(w, 200, strconv.FormatInt(count, 10))
	}

}

func FindReminder(id string, c types.AppConfig, w http.ResponseWriter) (*persistence.Reminder, bool) {
	if id == "" {
		respondWithError(w, 416, "Missing 'UID' parameter")
		return &persistence.Reminder{}, false
	}

	uid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		respondWithError(w, 416, "Invalid UID parameter value")
		return &persistence.Reminder{}, false
	}

	orgR := c.Repo.FindReminder(uid)
	if orgR == nil {
		respondWithError(w, 404, "Reminder not found for UID="+id)
		return &persistence.Reminder{}, false
	}
	return orgR, true
}

func toModel(reminder persistence.Reminder) reminders.Reminder {
	uid := reminder.UID.Hex()
	r := reminders.Reminder{
		UID:       uid,
		Title:     reminder.Title,
		Comment:   reminder.Comment,
		Completed: reminder.Completed,
	}
	return r
}

func respondWithError(w http.ResponseWriter, rc int, msg string) {
	response, _ := json.Marshal(msg)
	writeResponse(w, rc, response)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	writeResponse(w, code, response)
}

func respondWithJSONReminder(w http.ResponseWriter, code int, reminder persistence.Reminder) {
	response, _ := json.Marshal(toModel(reminder))
	writeResponse(w, code, response)
}

func respondWithJSONReminders(w http.ResponseWriter, code int, ra []persistence.Reminder) {
	rs := make(reminders.Reminders, len(ra))

	for i, rx := range ra {
		rs[i] = toModel(rx)
	}

	response, _ := json.Marshal(&rs)
	writeResponse(w, code, response)
}

func writeResponse(w http.ResponseWriter, rc int, content []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(rc)
	w.Write(content)

}
