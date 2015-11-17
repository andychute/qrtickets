package main
import (
	"fmt"
	"html"
	"log"
	"net/http"
	
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	
	// Setup routes
	router.HandleFunc("/",Index)
	router.HandleFunc("/api/v1",ApiIndex)
	router.HandleFunc("/api/v1/events",EventList)
	router.HandleFunc("/api/v1/events/{eventId}", EventShow)
	
	log.Fatal(http.ListenAndServe(":8080",router))
}

func EventShow (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	eventId := vars["eventId"]
	fmt.Fprintln(w, "Event ID:", eventId)
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func ApiIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func EventList(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

