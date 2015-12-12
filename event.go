package qrtickets

import (
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"net/http"
	"time"
)

// Event - define a performance / event
type Event struct {
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	Headline    string    `json:"headline"`
	Description string    `json:"description"`
	Tickets     []*Ticket `json:"-"`
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
