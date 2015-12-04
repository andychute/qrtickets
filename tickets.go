package qrtickets

import (
	"net/http"
)

// Ticket - Outlines a digital ticket
type Ticket struct {
	TicketNumber, Sig1, Sig2 string
}

// VerifySignature - Read hash, sig1, and sig2 from HTTP handler and verify
func VerifySignature(w http.ResponseWriter, r *http.Request) {
	// Reads {qrCode} from URL and outputs image/png bytestream
	//	vars := mux.Vars(r)

}
