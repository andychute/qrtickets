package qrtickets

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/gorilla/mux"
	"math/big"
	"net/http"
	"crypto/rand"
)

// Ticket - Outlines a digital ticket
type Ticket struct {
	TicketNumber, Sig1, Sig2 string
}

// TicketNumber - Plain text 
type TicketNumber struct {
	ID []byte
	Sig1,Sig2 *big.Int
}

// signTicket - Signs the Ticket Number
func (n *TicketNumber) sign() {
	conf := ConfLoad()
	sig1, sig2, _ := ecdsa.Sign(rand.Reader, &conf.PrivateKey, n.ID)
	n.Sig1, n.Sig2 = sig1,sig2
}

// GenTicket - Read ticket number from URL and Generate a QR Code
func GenTicket (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// Setup the Ticketnumber and sign it
	var ticketnum = TicketNumber{ID: []byte(vars["hash"])}
	ticketnum.sign()

	
	fmt.Fprintf(w, "sig1: %#v \n sig2: %#v \n message: %#v",ticketnum.Sig1,ticketnum.Sig2,vars["hash"])	
}

// VerifySignature - Read hash, sig1, and sig2 from HTTP handler and verify
func VerifySignature(w http.ResponseWriter, r *http.Request) {
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
