package qrtickets

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
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

// VerifyAuth - Compare Header provided in HTTP request against app.yaml environment variable
func VerifyAuth(i string) bool {
	auth := os.Getenv("HTTP_AUTH")
	if i == "" || auth == "" || i != auth {
		log.Printf("%#v - %#v", auth, i)
		return false
	}
	return true
}
