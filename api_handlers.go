package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"html"
	"net/http"
	"time"
)

func EventShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	eventId := vars["eventId"]
	fmt.Fprintln(w, "Event ID:", eventId)
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func ApiIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func EventList(w http.ResponseWriter, r *http.Request) {
	layout := "2006-01-02 15:04:05"

	ts, _ := time.Parse(layout, "2015-11-27 00:33:00")
	te, _ := time.Parse(layout, "2015-11-28 01:13:00")
	events := Events{
		Event{Headline: "Write presentation", StartTime: ts, EndTime: te, Description: "It's an event"},
		Event{Headline: "Present Presentation", StartTime: ts, EndTime: te, Description: "It's an event"},
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(events); err != nil {
		panic(err)
	}
}

