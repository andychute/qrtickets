package qrtickets

import "net/http"

// Route - Define information necessary to route a url request to handler function
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

var routes = []Route{{
	"",
	"GET",
	"/api/v1/events",
	EventList,
}, {
	"EventShow",
	"GET",
	"/api/v1/events/{eventId}",
	EventShow,
}, {
	"GenQR",
	"GET",
	"/gencode/{qrCode}",
	GenQR,
}, {
	"GenSignature",
	"GET",
	"/gensig",
	GenSignature,
}, {
	"VerifySig",
	"GET",
	"/api/v1/tickets/{sig1}/{sig2}/{hash}",
	VerifySignature,
}, {
	"GenerateTicket",
	"GET",
	"/api/v1/tickets/generate/{hash}",
	GenTicket,
}, {
	"LoadConf",
	"GET",
	"/loadconf",
	WebConfLoad,
}, {
	"TestSign",
	"GET",
	"/testsign",
	TestSign,
}}
