package qrtickets

import (
	"github.com/gorilla/mux"
	"net/http"
)

// NewRouter - Wrap with custom Logger to enable request logging
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name, route.AdminOnly)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}

	return router
}
