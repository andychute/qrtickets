package savoytickets

import (
	"path/filepath"
    "encoding/json"
    "os"
    "fmt"
)

type Configuration struct {
    Key ConfKey
}

type ConfKey struct {
	P,N,B,Gx,Gy,BitSize,X,Y,D string
}

func LoadConfig() Configuration {

	// Load Configuration
	// Loads the configuration in JSON format from conf.json into the Configuration struct

	path, _ := filepath.Abs("../src/bitbucket.org/capnfuzz/savoytickets/conf.json");
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


