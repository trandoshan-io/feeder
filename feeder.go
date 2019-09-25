package main

import (
	"encoding/json"
	"github.com/nats-io/nats.go"
	"log"
	"os"
)

func main() {
	if os.Getenv("INITIAL_URI") == "" {
		log.Fatal("Missing INITIAL_URI system property")
	}

	log.Println("Initializing feeder on url: ", os.Getenv("INITIAL_URI"))

	// connect to NATS server
	nc, err := nats.Connect(os.Getenv("NATS_URI"))
	if err != nil {
		log.Fatal("Error while connecting to nats server: ", err)
	}
	defer nc.Close()

	log.Println("Feeding url " + os.Getenv("INITIAL_URI") + " to web-crawler")

	bytes, err := json.Marshal(os.Getenv("INITIAL_URI"))
	if err != nil {
		log.Fatal("Error while serializing message into json: ", err)
	}

	if err := nc.Publish("todoSubject", bytes); err != nil {
		log.Fatal("Error while publishing message: ", err)
	}
}
