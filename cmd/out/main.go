package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/pivotal-cf/staticfile-resource/api"
)

func main() {
	var outRequest api.OutRequest
	err := json.NewDecoder(os.Stdin).Decode(&outRequest)
	if err != nil {
		log.Fatal("failed to decode: ", err)
	}

	json.NewEncoder(os.Stdout).Encode(api.Response{
		api.ResponseVersion{
			Version: "latest",
		},
	})
}
