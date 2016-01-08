package qrtickets

import (
	"google.golang.org/appengine/datastore"
	"time"
)

// Booking - A performer scheduled to perform at an event
type Booking struct {
	EventKey     datastore.Key `json:"event"`
	PerformerKey datastore.Key `json:"performer"`
	SetTime      time.Time     `json:"setTime"`
}
