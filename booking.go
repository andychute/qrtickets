package qrtickets

import (
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"time"
	"net/http"
	"golang.org/x/net/context"
)

// Booking - A performer scheduled to perform at an event
type Booking struct {
	EventKey     datastore.Key `json:"event"`
	PerformerKey datastore.Key `json:"performer"`
	SetTime      time.Time     `json:"setTime"`
	
	DatastoreKey datastore.Key  `json:"booking_id" datastore:"-"`
}

// AddBooking - Associate a performer to an event
func AddBooking(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	
	const timeformat = "2006-01-02 15:04:05 -0700"
	st, _ := time.Parse(timeformat, r.FormValue("set_time"))

	// Create the Venue object
	p1 := Booking{
		SetTime: st,
	}

	// Load the Performer Key
	if len(r.FormValue("performer")) > 0 {
		performer, err := datastore.DecodeKey(r.FormValue("performer"))
		if err != nil {
			JSONError(&w, err.Error())
			return
		}
		p1.PerformerKey = *performer
	}

	// Load the Event Key
	if len(r.FormValue("event")) > 0 {
		event, err := datastore.DecodeKey(r.FormValue("event"))
		if err != nil {
			JSONError(&w, err.Error())
			return
		}
		p1.EventKey = *event
	}

	// Add the booking to the Datastore
	k, err := p1.Store(ctx)
	if err != nil {
		JSONError(&w, err.Error())
		return
	}
	
	p1.DatastoreKey = k
	return
}

// Store - Stores the current booking into Google datastore
func (e *Booking) Store(ctx context.Context) (datastore.Key, error) {
	var k *datastore.Key

	// See if a key exists, or if a new one is required
	if e.DatastoreKey.Incomplete() {
		k = datastore.NewIncompleteKey(ctx, "Booking", nil)
	} else {
		k = &e.DatastoreKey
	}

	// Stash the entry in the datastore
	key, err := datastore.Put(ctx, k, e)
	if err != nil {
		panic(err)
	}

	return *key, nil
}

// Load - Loads the booking from Google datastore into the booking object
func (e *Booking) Load(ctx context.Context, k datastore.Key) error {
	err := datastore.Get(ctx, &k, e)
	e.DatastoreKey = k

	if err != nil {
		return err
	}

	return nil
}
