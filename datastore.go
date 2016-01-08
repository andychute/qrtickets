package qrtickets

import (
	"encoding/json"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"net/http"
)

func getEvents(r *http.Request, limit int) ([]Event, error) {
	ctx := appengine.NewContext(r)
	q := datastore.NewQuery("Event").Order("-StartDate").Limit(limit)

	var results []Event
	k, err := q.GetAll(ctx, &results)

	if err != nil {
		return nil, err
	}

	for i := range results {
		results[i].DatastoreKey = *k[i]
	}

	return results, nil
}

// EventList - Retrieve event objects from google datastore (most recent first)
func EventList(w http.ResponseWriter, r *http.Request) {
	results, err := getEvents(r, 50)
	if err != nil {
		JSONError(&w, err.Error())
	} else {
		json.NewEncoder(w).Encode(results)
	}
}
