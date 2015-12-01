package qrtickets

import "time"

// Event - define a performance / event
type Event struct {
	EventID               int
	StartTime, EndTime    time.Time
	Headline, Description string
	// Venue                 Venue
}

// Events - Collection of Event Objects
type Events []Event
