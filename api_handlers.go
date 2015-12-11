package qrtickets

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"net/http"
	"time"
)

// EventShow - Return details for specific event with id=vars['eventId']
func EventShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	eventID := vars["eventId"]
	if eventID == "" {
		http.Error(w, "No Event ID Provided", 500)
		return
	}

	ctx := appengine.NewContext(r)
	k, err := datastore.DecodeKey(eventID)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	var e Event
	if err = datastore.Get(ctx, k, &e); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Fprintf(w, "%#v", e)
}

// EventList - Returns JSON list of events
func EventList(w http.ResponseWriter, r *http.Request) {
	layout := "2006-01-02 15:04:05"

	ts, _ := time.Parse(layout, "2015-11-27 00:33:00")
	te, _ := time.Parse(layout, "2015-11-28 01:13:00")
	events := []Event{{Headline: "Write presentation", StartTime: ts, EndTime: te, Description: "It's an event"},
		{Headline: "Present Presentation", StartTime: ts, EndTime: te, Description: "It's an event"},
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(events); err != nil {
		panic(err)
	}
}
