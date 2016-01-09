package qrtickets

import (
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"net/http"
)

// Venue - A physical location that hosts events
type Venue struct {
	Name    string `json:"name" datastore:",noindex"`
	Address string `json:"address" datatore:",noindex"`
	URL     string `json:"url" datastore:",noindex"`

	DatastoreKey datastore.Key `json:"venue_id" datastore:"-"`
}

// AddVenue - Add a venue from form input and save to Datastore
func AddVenue(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	// Create the Venue object
	p1 := Venue{
		Name: r.FormValue("name"),
		Address: r.FormValue("address"),
		URL:  r.FormValue("url"),
	}

	// Add the venue to the Datastore
	k, err := p1.Store(ctx)
	if err != nil {
		JSONError(&w, err.Error())
		return
	}

	p1.DatastoreKey = *k
	return
}

// Store - Stores the current venue into Google datastore
func (p *Venue) Store(ctx context.Context) (*datastore.Key, error) {
	var k *datastore.Key

	// See if a key exists, or if a new one is required
	if p.DatastoreKey.Incomplete() {
		k = datastore.NewIncompleteKey(ctx, "Venue", nil)
	} else {
		k = &p.DatastoreKey
	}

	// Stash the entry in the datastore
	key, err := datastore.Put(ctx, k, p)
	if err != nil {
		return nil, err
	}

	return key, nil
}

