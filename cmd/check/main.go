package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/christianang/staticfile-resource/api"
)

func main() {
	var checkRequest api.InRequest
	err := json.NewDecoder(os.Stdin).Decode(&checkRequest)
	if err != nil {
		log.Fatal("failed to decode: ", err)
	}

	json.NewEncoder(os.Stdout).Encode([]api.ResponseVersion{
		{
			Version: "latest",
		},
	})
}
