package qrtickets

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

// EventShow - Return details for specific event with id=vars['eventId']
func EventShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	eventID := vars["eventId"]
	if eventID == "" {
		http.Error(w, "No Event ID Provided", 500)
		return
	}

	event, err := LoadEvent(r, eventID)
	if err != nil {
		http.Error(w, "No Event ID Provided", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(event); err != nil {
		panic(err)
	}
}
