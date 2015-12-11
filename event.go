package qrtickets

import "time"

// Event - define a performance / event
type Event struct {
	StartTime, EndTime    time.Time
	Headline, Description string
	Tickets               []*Ticket
}
