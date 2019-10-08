package main

import (
	"encoding/json"
	"github.com/nats-io/nats.go"
	"log"
	"os"
)

const (
	todoSubject = "todoSubject"
)

func main() {
	initialUri := os.Getenv("INITIAL_URI")

	if initialUri == "" {
		log.Fatal("Missing INITIAL_URI system property")
	}

	log.Printf("Initializing feeder on url: %s", initialUri)

	// connect to NATS server
	nc, err := nats.Connect(os.Getenv("NATS_URI"))
	if err != nil {
		log.Fatalf("Error while connecting to nats server: %s", err)
	}
	defer nc.Close()

	log.Printf("Feeding url %s to web-crawler", initialUri)

	bytes, err := json.Marshal(initialUri)
	if err != nil {
		log.Fatalf("Error while serializing message into json: %s", err)
	}

	if err := nc.Publish(todoSubject, bytes); err != nil {
		log.Fatalf("Error while publishing message: %s", err)
	}
}
