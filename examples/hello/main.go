package main

import (
	"log"

    "app"
)

func main() {

    berlioz.Peers("service", "app", "client").Monitor(func(peers berlioz.PeerAccessor) {
        log.Printf("---------- PEER MONITOR -----------------")
        log.Printf("--- PEERS: %v\n", peers.All())
        if val, ok := peers.Get("1"); ok {
            log.Printf("--- INDEXED PEER: %v\n", val)
        }
        if val, ok := peers.Random(); ok {
            log.Printf("--- RANDOM PEER: %v\n", val)
        }
    })

	resp, body, err := berlioz.Request("service", "app", "client").Get("/")
	if err != nil {
		log.Printf("Response Error: %s\n", err)
	} else {
		log.Printf("Response Status Code: %s\n", resp.Status)
		log.Printf("Response Body: %s\n", body)
	}
}
