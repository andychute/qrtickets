package qrtickets

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/json"
	"os"
)

// Config - Load the private key into the config
type Config struct {
	ecdsa.PrivateKey
}

// ConfLoad - Load configuration from app.yaml
func ConfLoad() *Config {
	var conf *Config

	// Read JSON config from app.yaml
	if v := os.Getenv("PRIV_KEY"); v != "" {
		err := json.Unmarshal([]byte(v), conf)
		if err != nil {
			panic(err)
		}
	}

	// Create the curve
	conf.PublicKey.Curve = elliptic.P224()

	// Return the conf

	// Try signing a message
	// message := []byte("99999999")
	// sig1, sig2, err := ecdsa.Sign(rand.Reader, &conf.PrivateKey, message)
	// if err != nil {
	// 	panic(err)
	// }

	// // Try verifying the signature
	// result := ecdsa.Verify(&conf.PublicKey, message, sig1, sig2)
	// if result != true {
	// 	panic("Unable to verify signature")
	// } else {
	// 	fmt.Fprintf(w, "sig1: %#v\nsig2: %#v", sig1, sig2)
	// }

	return conf
}
