package qrtickets

import (
	"log"
	"net/http"
	"fmt"
	"time"
	"os"
)

// VerifyAuth - Compare Header provided in HTTP request against app.yaml environment variable
func VerifyAuth (i string) bool {
	auth := os.Getenv("HTTP_AUTH")
	if i == "" || auth == "" || i != auth {
		log.Printf("%#v - %#v",auth,i)
		return false
	}
	return true
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
		if admin != false  {			
			// Load the Ticket Auth Header from the request
			pass := r.Header["X-Ticket-Auth"][0]
			if VerifyAuth(pass) {
				// Valid Request
				fmt.Fprintf(w,"Yay!!!!")
				log.Println(w,"LOGIN SUCCESS")
				inner.ServeHTTP(w, r)
			} else {			
				// Unable to validate Request
				fmt.Fprintf(w,`{error: true,message: "No / Invalid login information provided"}`)
				log.Println(w,"ACCESS DENIED")
			}
		} else {
			inner.ServeHTTP(w, r)
		}		
	})
}
