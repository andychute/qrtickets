package savoytickets

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/api/v1",
		Index,
	},
	Route{
		"",
		"GET",
		"/api/v1/events",
		EventList,
	},
	Route{
		"EventShow",
		"GET",
		"/api/v1/events/{eventId}",
		EventShow,
	},
	Route{
		"GenQR",
		"GET",
		"/gencode/{qrCode}",
		GenQR,
	},
	Route{
		"GenSignature",
		"GET",
		"/gensig",
		GenSignature,
	},
}
