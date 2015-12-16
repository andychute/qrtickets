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
	"time"
)

// Ticket - Outlines a digital ticket
type Ticket struct {
	OrderID      string         `json:"order_id"`
	EventKey     *datastore.Key `json:"event" datastore:"-"`
	Valid        bool           `json:"valid"`
	Claimed      bool           `json:"claimed"`
	DatastoreKey datastore.Key  `datastore:"-"`
	DateAdded    time.Time      `json:"date_added"`
}

// Load - Takes a datastore.Key provided and loads it into the current Ticket object
func (t *Ticket) Load(ctx context.Context, k datastore.Key) error {
	err := datastore.Get(ctx, &k, t)
	t.DatastoreKey = k
	t.EventKey = k.Parent()

	if err = datastore.Get(ctx, &k, t); err != nil {
		return err
	}

	return nil
}

// Store - Store the ticket entry in google datastore
func (t *Ticket) Store(ctx context.Context) (*datastore.Key, error) {
	var k *datastore.Key

	// See if a key exists, or if a new one is required
	if t.DatastoreKey.Incomplete() {
		k = datastore.NewIncompleteKey(ctx, "ticket", t.EventKey)
		t.DateAdded = time.Now()
	} else {
		k = &t.DatastoreKey
	}

	// Stash the entry in the datastore
	key, err := datastore.Put(ctx, k, t)
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
		OrderID:  r.FormValue("order_id"),
		EventKey: event,
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
	code.Scale = 2

	if err != nil {
		panic(err)
	}

	imgByte := code.PNG()
	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Content-Disposition", `inline; filename="`+k.Encode()+`"`)
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
	code, err := qr.Encode(buffer.String(), qr.H)
	if err != nil {
		panic(err)
	}

	imgByte := code.PNG()

	w.Header().Set("Content-Type", "image/png")
	w.WriteHeader(http.StatusOK)
	w.Write(imgByte)

	//	fmt.Fprintf(w, "sig1: %#v \n sig2: %#v \n message: %#v",ticketnum.Sig1,ticketnum.Sig2,vars["hash"])
}

// ClaimTicket - Verify the ticket and claim it
func ClaimTicket(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// Define container variables
	var hash []byte
	sig1 := new(big.Int)
	sig2 := new(big.Int)

	// Assign the variables
	sig1.SetString(vars["sig1"], 10)
	sig2.SetString(vars["sig2"], 10)
	hash = []byte(vars["hash"])

	if VerifySignature(sig1, sig2, hash) {
		// Signature has been successfully verified, Claim it in the datastore
		ctx := appengine.NewContext(r)
		var t Ticket

		key, err := datastore.DecodeKey(vars["hash"])
		// Check for decoding errors
		if err != nil {
			JSONError(&w, "Unable to Decode Key from provided hash")
			return
		}

		// Map the results to the receiving object
		if err = t.Load(ctx, *key); err != nil {
			JSONError(&w, "Unable to Retrieve Key")
			return
		}

		// Check for ticket validity
		if t.Valid == false {
			JSONError(&w, "Ticket has been invalidated")
			return
		}

		if t.Claimed == true {
			JSONError(&w, "Ticket has already been claimed")
			return
		}

		// All good, update the ticket to claimed and resave it
		t.Claimed = true
		t.Store(ctx)

		fmt.Fprintf(w, "%#v", t)
		return
	}

	JSONError(&w, "Unable to Verify Signature")
	return
}

// VerifySignature - Read hash, sig1, and sig2 from HTTP handler and verify
func VerifySignature(s1, s2 *big.Int, hash []byte) bool {

	// Setup the signatures
	ticketnum := TicketNumber{ID: hash, Sig1: s1, Sig2: s2}
	return ticketnum.verify()
}
