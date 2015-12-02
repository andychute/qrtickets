package qrtickets

import (
	"code.google.com/p/rsc/qr"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"strings"
)

// GenQR - Generate a PNG QR code based on URL argument
func GenQR(w http.ResponseWriter, r *http.Request) {

	// QR Code Generation Function
	// Reads {qrCode} from URL and outputs image/png bytestream
	vars := mux.Vars(r)
	code, err := qr.Encode(vars["qrCode"], qr.H)
	if err != nil {
		panic(err)
	}

	imgByte := code.PNG()
	w.Header().Set("Content-Type", "image/png")
	w.WriteHeader(http.StatusOK)
	w.Write(imgByte)
}

// ConfLoad - Load configuration from app.yaml
func ConfLoad(w http.ResponseWriter, r *http.Request) {
	var conf Config

	if v := os.Getenv("PRIV_KEY"); v != "" {
		json.NewDecoder(strings.NewReader(v)).Decode(&conf)
	}
	conf.PublicKey.Curve = elliptic.P224()

	// Try signing a message
	/* message := []byte("This is a test")
	sig1, sig2, err := ecdsa.Sign(rand.Reader, conf.PrivateKey, message)

	if err != nil {
		panic(err)
	}


	// Try verifying the signature
	result := ecdsa.Verify(&conf.PublicKey, message, sig1, sig2)
	if result != true {
		panic("Unable to verify signature")
	}
	*/
	// Log the output
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%#v \n%#v", conf)

}

// GenSignature - Debugging function to sign a message
func GenSignature(w http.ResponseWriter, r *http.Request) {
	// Returns a Public / Private Key Pair
	// Uses eliptic curve cryptography

	// Generate a public / private key pair
	privatekey := new(ecdsa.PrivateKey)

	// Generate an elliptic curve using NIST P-224
	ecurve := elliptic.P224()
	privatekey, err := ecdsa.GenerateKey(ecurve, rand.Reader)

	if err != nil {
		panic(err)
	}

	// Marshal the JSON
	privkey, _ := json.Marshal(privatekey)
	publikey, _ := json.Marshal(privatekey.Public())

	// Get the public key
	var pubkey ecdsa.PublicKey
	pubkey = privatekey.PublicKey

	// Try signing a message
	message := []byte("This is a test")
	sig1, sig2, err := ecdsa.Sign(rand.Reader, privatekey, message)

	// Try verifying the signature
	result := ecdsa.Verify(&pubkey, message, sig1, sig2)
	if result != true {
		panic("Unable to verify signature")
	}

	fmt.Fprintln(w, "Marshaled Private Key:", privkey)
	fmt.Fprintln(w, "Marshaled Public Key:", publikey)
	fmt.Fprintln(w, "Curve: ", pubkey.Curve)
	fmt.Fprintf(w, "Curve: Private: %#v\nPublic: %#v\n\nSignature:\n%v\n%v\n\nVerified: %v", privatekey, pubkey, sig1, sig2, result)

	// Now
}
