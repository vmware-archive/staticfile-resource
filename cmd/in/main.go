package main

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/christianang/staticfile-resource/api"
)

func main() {
	destinationDirectory := os.Args[1]

	var inRequest api.InRequest
	err := json.NewDecoder(os.Stdin).Decode(&inRequest)
	if err != nil {
		log.Fatal("failed to decode: ", err)
	}

	in := api.NewIn()

	for _, file := range inRequest.Source.Files {
		err = in.WriteDataToDestination(
			filepath.Join(destinationDirectory, file.Filename),
			[]byte(file.Data),
		)
		if err != nil {
			log.Fatal("failed to write data to file: ", err)
		}
	}

	json.NewEncoder(os.Stdout).Encode(api.Response{
		api.ResponseVersion{
			Version: "latest",
		},
	})
}
