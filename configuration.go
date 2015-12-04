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
	var conf Config

	// Read JSON config from app.yaml
	if v := os.Getenv("PRIV_KEY"); v != "" {
		err := json.Unmarshal([]byte(v), &conf)
		if err != nil {
			panic(err)
		}
	}

	// Create the curve
	conf.PublicKey.Curve = elliptic.P224()
	return &conf
}
