package qrtickets

import "net/http"

// Route - Define information necessary to route a url request to handler function
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
	AdminOnly   bool
}

var routes = []Route{{
	"",
	"GET",
	"/api/v1/events",
	EventList,
	false,
}, {
	"AddEvent",
	"POST",
	"/api/v1/events/create",
	AddEvent,
	true,
}, {
	"EventShow",
	"GET",
	"/api/v1/events/{eventId}",
	EventShow,
	false,
}, {
	"GenSignature",
	"GET",
	"/gensig",
	GenSignature,
	true,
}, {
	"ClaimTicket",
	"GET",
	"/api/v1/tickets/{sig1:[0-9]+}/{sig2:[0-9]+}/{hash:[-0-9a-zA-Z]+}/claim",
	ClaimTicket,
	false,
}, {
	"AddTicket",
	"GET",
	"/api/v1/events/{eventId:[-0-9a-zA-Z]+}/ticket/add",
	AddTicket,
	true,
}, {
	"LoadConf",
	"GET",
	"/loadconf",
	WebConfLoad,
	true,
}, {
	"TestSign",
	"GET",
	"/testsign",
	TestSign,
	true,
}}
