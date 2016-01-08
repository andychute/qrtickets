package qrtickets

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"net/http"
	"time"
)

// Event - define a performance / event
type Event struct {
	StartDate   time.Time `json:"startDate"`
	DoorTime    time.Time `json:"doorTime"`
	EndDate     time.Time `json:"endDate"`
	Name        string    `json:"name" datastore:",noindex"`
	Description string    `json:"description" datastore:",noindex"`
	URL         string    `json:"url"`
	DateAdded   time.Time `json:"date_added" datastore:",noindex"`

	Promoter string        `json:"promoter" datastore:",noindex"`
	Image    string        `json:"image" datastore:",noindex"`
	Venue    datastore.Key `json:"Venue"`

	// Additional Datastore Variables
	DatastoreKey datastore.Key `json:"event_id" datastore:"-"`
}

// LoadEvent - Accepts a key to look up in datastore
// returns event object
func LoadEvent(r *http.Request, id string) (*Event, error) {
	ctx := appengine.NewContext(r)
	k, err := datastore.DecodeKey(id)

	if err != nil {
		return nil, err
	}

	var e Event
	if err = datastore.Get(ctx, k, &e); err != nil {
		return nil, err
	}

	return &e, nil
}

// AddEvent - Add an Event and save to Datastore
func AddEvent(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	// Set the timestamps
	at := time.Now()
	const timeformat = "2006-01-02 15:04:05"
	st, _ := time.Parse(timeformat, r.FormValue("start_time"))
	et, _ := time.Parse(timeformat, r.FormValue("end_time"))

	// Create the event object
	e1 := Event{
		StartDate:   st,
		EndDate:     et,
		DateAdded:   at,
		Name:        r.FormValue("headline"),
		Description: r.FormValue("description"),
		URL:         r.FormValue("event_url"),
		Promoter:    r.FormValue("promoter"),
		Image:       r.FormValue("poster_file"),
	}

	// Add the event to the Datastore
	k, err := e1.Store(ctx)
	if err != nil {
		JSONError(&w, err.Error())
		return
	}

	e1.DatastoreKey = *k
	return
}

// Store - Stores the current event into Google datastore
func (e *Event) Store(ctx context.Context) (*datastore.Key, error) {
	var k *datastore.Key

	// See if a key exists, or if a new one is required
	if e.DatastoreKey.Incomplete() {
		k = datastore.NewIncompleteKey(ctx, "Event", nil)
	} else {
		k = &e.DatastoreKey
	}

	// Stash the entry in the datastore
	key, err := datastore.Put(ctx, k, e)
	if err != nil {
		return nil, err
	}

	return key, nil
}

// Load - Loads the event from Google datastore into the event object
func (e *Event) Load(ctx context.Context, k datastore.Key) error {
	err := datastore.Get(ctx, &k, e)
	e.DatastoreKey = k

	if err != nil {
		return err
	}

	return nil
}

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
