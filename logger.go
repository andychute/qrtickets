package qrtickets

import (
	"log"
	"net/http"
	"time"
)

// Logger - Add additional logging to http.Handler
func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)
		log.Printf(
			"%s\t%q\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}
