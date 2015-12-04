package qrtickets

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/gorilla/mux"
	"math/big"
	"net/http"
)

// Ticket - Outlines a digital ticket
type Ticket struct {
	TicketNumber, Sig1, Sig2 string
}

// VerifySignature - Read hash, sig1, and sig2 from HTTP handler and verify
func VerifySignature(w http.ResponseWriter, r *http.Request) {

	// Reads {qrCode} from URL and outputs image/png bytestream
	vars := mux.Vars(r)
	var hash []byte

	sig1 := new(big.Int)
	sig2 := new(big.Int)

	sig1.SetString(vars["sig1"], 10)
	sig2.SetString(vars["sig2"], 10)
	hash = []byte(vars["hash"])

	// Load the Configuration
	conf := ConfLoad()
	pubkey := conf.PublicKey

	result := ecdsa.Verify(&pubkey, hash, sig1, sig2)
	if result != true {
		fmt.Fprintf(w, "Unable to verify signature")
	} else {
		fmt.Fprintf(w, "Yay!")
	}

}
