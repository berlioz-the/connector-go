package main

import (
	"fmt"
	"time"

    "app"
)

func main() {
	time.Sleep(1 * time.Second)

	// berlioz.DoSomething()

	fmt.Printf("--- PEERS: %v\n", berlioz.Peers("service", "app", "client").All())
	fmt.Printf("--- INDEXED PEER: %v\n", berlioz.Peers("service", "app", "client").Get("1"))
	fmt.Printf("--- RANDOM PEER: %v\n", berlioz.Peers("service", "app", "client").Random())

	resp, body, err := berlioz.Request("service", "app", "client").Get("/")

	if err != nil {
		fmt.Printf("Response Error: %s\n", err)
	} else {
		fmt.Printf("Response Status Code: %s\n", resp.Status)
		fmt.Printf("Response Body: %s\n", body)
	}
}
