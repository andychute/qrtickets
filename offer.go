package qrtickets

import (
	"google.golang.org/appengine/datastore"
	"time"
)

// Offer - An item for sale in relation to an event
type Offer struct {
	TicketPrice  float32       `json:"price" datastore:",noindex"`
	Currency     string        `json:"priceCurrency" datastore:",noindex"`
	EventKey     datastore.Key `json:"event_id"`
	Category     string        `json:"category" datatstore:"-"`
	ValidFrom    time.Time     `json:"validFrom"`
	ValidThrough time.Time     `json:"validThrough"`
}
