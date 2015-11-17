package main

import "time"

type Event struct {
	EventId               int
	StartTime, EndTime    time.Time
	Headline, Description string
	// Venue                 Venue
}

type Events []Event
