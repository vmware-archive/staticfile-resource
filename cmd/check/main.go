package main

import (
	"encoding/json"
	"fmt"
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

	fmt.Println("[{}]")
}
