package qrtickets

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// JSONError - Throw a JSON Error
func JSONError(w *http.ResponseWriter, m string) {
	// Display Error to HTTP Handler
	fmt.Fprintf(*w, `{error: true,message: "%s"}`, m)
}

// Logger - Add additional logging to http.Handler
func Logger(inner http.Handler, name string, admin bool) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Log the request
		log.Printf(
			"%s\t%q\t%s\t%s\t",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)

		// Check to see if the method is admin only
		if admin != false {
			// Load the Ticket Auth Header from the request
			if len(r.Header["X-Ticket-Auth"]) > 0 {
				pass := r.Header["X-Ticket-Auth"][0]
				if VerifyAuth(pass) {
					// Valid Request
					log.Println(w, "LOGIN SUCCESS")
					inner.ServeHTTP(w, r)
				} else {
					JSONError(&w, "Could not verify HTTP password")
				}
			} else {
				JSONError(&w, "No Auth Header Provided")
			}
		} else {
			inner.ServeHTTP(w, r)
		}
	})
}
