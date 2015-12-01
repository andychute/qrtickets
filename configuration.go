package qrtickets

import (
	"crypto/ecdsa"
)

// Config - Load the private key into the config
type Config struct {
	ecdsa.PrivateKey
}



/*
func LoadConfig() Config {
	// Load Configuration
	// Loads the configuration in JSON format from conf.json into the Configuration struct

	path, _ := filepath.Abs("../src/bitbucket.org/capnfuzz/qrtickets/conf.json")
	file, filerr := os.Open(path)
	if filerr != nil {
		fmt.Println("error:", filerr)
	}
	decoder := json.NewDecoder(file)

	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}

	return configuration
}
*/
