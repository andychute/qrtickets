package main

import (
	"fmt"
	"net/http"
	"crypto/elliptic"
	"crypto/ecdsa"
	"crypto/rand"
	"code.google.com/p/rsc/qr"
	"github.com/gorilla/mux"
)

func GenQR(w http.ResponseWriter, r *http.Request) {

	// QR Code Generation Function
	// Reads {qrCode} from URL and outputs image/png bytestream
	
	vars := mux.Vars(r)
	
	code, err := qr.Encode(vars["qrCode"], qr.H)
	if err != nil {
		panic (err)
	}

	imgByte := code.PNG()
	w.Header().Set("Content-Type", "image/png")
	w.WriteHeader(http.StatusOK)
	w.Write(imgByte)
}

func GenSignature(w http.ResponseWriter, r *http.Request) {
	// Returns a Public / Private Key Pair 
	// Uses eliptic curve cryptography

	// Generate a public / private key pair
	privatekey := new(ecdsa.PrivateKey);

	// Generate an elliptic curve using NIST P-224
	ecurve := elliptic.P224()

	privatekey, err := ecdsa.GenerateKey(ecurve,rand.Reader);
	if err != nil {
		panic (err)
	}
	
	// Get the public key
	pubkey := privatekey.Public()
	
	// Try signing a message
	message := []byte("This is a test")
	sig1, sig2, err := ecdsa.Sign(rand.Reader, privatekey, message)

	// Try verifying the signature
	result := ecdsa.Verify(pubkey,message, sig1, sig2);
	if result != true {
		panic("Unable to verify signature");
	}
	
	// Log the output
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Private: %#v\nPublic: %#v\n\nSignature:\n%v\n%v\n\nVerified: %v",privatekey,pubkey,sig1,sig2,result)

	// Now 
}
