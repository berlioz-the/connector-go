package main

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/berlioz-the/connector-go"
)

func main() {
	time.Sleep(2 * time.Second)

	fmt.Println(berlioz.Request("service", "app", "client").GetRandomEndpoint())

	req := &http.Request{URL: &url.URL{Path: "/"}}
	resp, err := berlioz.Request("service", "app", "client").Do(req)

	fmt.Printf("Response: %s, error: %s\n", resp, err)
}
