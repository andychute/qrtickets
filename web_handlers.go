package qrtickets
import (

	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

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

	fmt.Fprintln(w, "Marshaled Private Key:", string(privkey))
	fmt.Fprintln(w, "Marshaled Public Key:", string(publikey))
	fmt.Fprintln(w, "Curve: ", pubkey.Curve)
	fmt.Fprintf(w, "Curve: Private: %#v\nPublic: %#v\n\nSignature:\n%v\n%v\n\nVerified: %v", privatekey, pubkey, sig1, sig2, result)

}

// TestSign - Return signatures for a signed message
func TestSign(w http.ResponseWriter, r *http.Request) {
	conf := ConfLoad()

	// Returns a Public / Private Key Pair
	// Read JSON config from app.yaml
	if v := os.Getenv("PRIV_KEY"); v != "" {
		err := json.Unmarshal([]byte(v), &conf)
		if err != nil {
			fmt.Printf("%#v", conf)
			panic(err)
		}
	}

	// Get the public key
	var pubkey ecdsa.PublicKey
	pubkey = conf.PublicKey

	// Try signing a message
	message := []byte("99999999")
	sig1, sig2, err := ecdsa.Sign(rand.Reader, &conf.PrivateKey, message)
	if err != nil {
		panic(err)
	}

	// Try verifying the signature
	result := ecdsa.Verify(&pubkey, message, sig1, sig2)
	if result != true {
		panic("Unable to verify signature")
	}

	fmt.Fprintf(w, "message: %#v\n\nsig1: %#v\nsig2: %#v", string(message[:]), sig1, sig2)

}

// WebConfLoad - Load Configuration and display to browser
func WebConfLoad(w http.ResponseWriter, r *http.Request) {
	var conf Config

	// Returns a Public / Private Key Pair
	// Uses eliptic curve cryptography

	// Read JSON config from app.yaml
	if v := os.Getenv("PRIV_KEY"); v != "" {
		err := json.Unmarshal([]byte(v), &conf)
		if err != nil {
			fmt.Printf("%#v", conf)
			panic(err)
		}
	}

	// Create the curve
	conf.PublicKey.Curve = elliptic.P224()
	fmt.Fprintf(w, "%#v", conf)

}
