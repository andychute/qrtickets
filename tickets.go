package qrtickets

import (
	"bytes"
	"code.google.com/p/rsc/qr"
	"crypto/ecdsa"
	"crypto/rand"
	"fmt"
	"github.com/gorilla/mux"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"math/big"
	"net/http"
)

// Ticket - Outlines a digital ticket
type Ticket struct {
	OrderID  string        `json:"order_id"`
	EventKey datastore.Key `json:"event"`
	Valid    bool          `json:"valid"`
	Claimed  bool          `json:"claimed"`
}

// Store - Store the ticket entry in google datastore
func (t *Ticket) Store(ctx context.Context) (*datastore.Key, error) {

	// Stash the entry in the datastore
	key, err := datastore.Put(ctx, datastore.NewIncompleteKey(ctx, "ticket", &t.EventKey), t)
	if err != nil {
		return nil, err
	}

	return key, nil
}

// TicketNumber - Plain text
type TicketNumber struct {
	ID         []byte
	Sig1, Sig2 *big.Int
}

// sign - Generates the signatures for the ticket ID utilizing the PrivateKey loaded from app.yaml
func (n *TicketNumber) sign() {
	conf := ConfLoad()
	sig1, sig2, err := ecdsa.Sign(rand.Reader, &conf.PrivateKey, n.ID)
	if err != nil {
		panic(err)
	}

	// Add the signature to the ticket number entry
	n.Sig1, n.Sig2 = sig1, sig2
}

// verify - Verifies the ticket's signatures against it's ID using the PublicKey loaded from app.yaml
func (n *TicketNumber) verify() bool {
	conf := ConfLoad()
	return ecdsa.Verify(&conf.PublicKey, n.ID, n.Sig1, n.Sig2)
}

// AddTicket - Adds a valid ticket for the event and stores it in the datastore
func AddTicket(w http.ResponseWriter, r *http.Request) {
	var buffer bytes.Buffer
	vars := mux.Vars(r)

	// Load the event datastore key
	event, err := datastore.DecodeKey(vars["eventId"])

	if err != nil {
		panic(err)
	}

	// Create an appengine context
	ctx := appengine.NewContext(r)
	// fmt.Fprintf("%#v",ctx)

	// Build the ticket entry
	t := Ticket{
		OrderID:  "",
		EventKey: *event,
		Valid:    true,
	}

	// Store the ticket
	k, err := t.Store(ctx)
	if err != nil {
		panic(err)
	}

	// Create the Ticket Num
	var ticketnum = TicketNumber{ID: []byte(k.Encode())}
	ticketnum.sign()

	// Generate the text string to encode
	buffer.WriteString(ticketnum.Sig1.String())
	buffer.WriteString("/")
	buffer.WriteString(ticketnum.Sig2.String())
	buffer.WriteString("/")
	buffer.WriteString(string(k.Encode()))

	// Generate the QR code for the hash and two signatures
	code, err := qr.Encode(buffer.String(), qr.L)
	code.Scale = 4

	if err != nil {
		panic(err)
	}

	imgByte := code.PNG()
	w.Header().Set("Content-Type", "image/png")
	w.WriteHeader(http.StatusOK)
	w.Write(imgByte)
}

// GenTicket - Sign the ticket number provided through the URL and Generate a QR Code
func GenTicket(w http.ResponseWriter, r *http.Request) {
	var buffer bytes.Buffer

	// Load the variables from the path using mux
	vars := mux.Vars(r)

	// Setup the Ticket and sign it
	var ticketnum = TicketNumber{ID: []byte(vars["hash"])}
	ticketnum.sign()

	// Generate the text string to encode
	buffer.WriteString(ticketnum.Sig1.String())
	buffer.WriteString("/")
	buffer.WriteString(ticketnum.Sig2.String())
	buffer.WriteString("/")
	buffer.WriteString(vars["hash"])

	// Generate the QR code for the hash and two signatures
	code, err := qr.Encode(buffer.String(), qr.L)
	if err != nil {
		panic(err)
	}

	imgByte := code.PNG()
	w.Header().Set("Content-Type", "image/png")
	w.WriteHeader(http.StatusOK)
	w.Write(imgByte)

	//	fmt.Fprintf(w, "sig1: %#v \n sig2: %#v \n message: %#v",ticketnum.Sig1,ticketnum.Sig2,vars["hash"])
}

// VerifySignature - Read hash, sig1, and sig2 from HTTP handler and verify
func VerifySignature(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	fmt.Fprintf(w, "%#v", r.Header["Cache-Control"])

	// Define container variables
	var hash []byte
	sig1 := new(big.Int)
	sig2 := new(big.Int)

	// Assign the variables
	sig1.SetString(vars["sig1"], 10)
	sig2.SetString(vars["sig2"], 10)
	hash = []byte(vars["hash"])

	// Setup the signatures
	ticketnum := TicketNumber{ID: hash, Sig1: sig1, Sig2: sig2}
	if ticketnum.verify() != true {
		fmt.Fprintf(w, "Unable to verify signature (priv method)")
	} else {
		fmt.Fprintf(w, "Ya!")
	}
}
