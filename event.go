package qrtickets

import (
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"net/http"
	"time"
)

// Event - define a performance / event
type Event struct {
	StartTime    time.Time     `json:"start_time"`
	EndTime      time.Time     `json:"end_time"`
	Headline     string        `json:"headline" datastore:",noindex"`
	Description  string        `json:"description" datastore:",noindex"`
	URL          string        `json:"url"`
	DateAdded    time.Time     `json:"date_added" datastore:",noindex"`
	DatastoreKey datastore.Key `datastore:"-"`
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
		StartTime:   st,
		EndTime:     et,
		DateAdded:   at,
		Headline:    r.FormValue("headline"),
		Description: r.FormValue("description"),
		URL:         r.FormValue("event_url"),
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
		k = datastore.NewIncompleteKey(ctx, "event", nil)
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
