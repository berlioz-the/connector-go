package main

import (
	"log"
	"time"

    "app"
)

func main() {
	time.Sleep(1 * time.Second)

	log.Printf("--- PEERS: %v\n", berlioz.Peers("service", "app", "client").All())
	log.Printf("--- INDEXED PEER: %v\n", berlioz.Peers("service", "app", "client").Get("1"))
	log.Printf("--- RANDOM PEER: %v\n", berlioz.Peers("service", "app", "client").Random())

	resp, body, err := berlioz.Request("service", "app", "client").Get("/")

	if err != nil {
		log.Printf("Response Error: %s\n", err)
	} else {
		log.Printf("Response Status Code: %s\n", resp.Status)
		log.Printf("Response Body: %s\n", body)
	}
}
