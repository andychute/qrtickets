package qrtickets

import (
	"encoding/json"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"net/http"
)

// EventList - Retrieve event objects from google datastore (most recent first)
func EventList(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	q := datastore.NewQuery("Event").Order("-StartTime").Limit(50)

	var results []Event
	k, err := q.GetAll(ctx, &results)
	if err != nil {
		JSONError(&w, err.Error())
	}

	for i := range results {
		results[i].DatastoreKey = *k[i]
	}

	json.NewEncoder(w).Encode(results)

}
