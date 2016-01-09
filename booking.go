package qrtickets

import (
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"time"
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
		e1.PerformerKey = performer
	}

	// Load the Event Key
	if len(r.FormValue("event")) > 0 {
		event, err := datastore.DecodeKey(r.FormValue("event"))
		if err != nil {
			JSONError(&w, err.Error())
			return
		}
		e1.EventKey = event
	}

	// Add the booking to the Datastore
	k, err := e1.Store(ctx)
	if err != nil {
		JSONError(&w, err.Error())
		return
	}
	
	e1.DatastoreKey = *k
	return
}
