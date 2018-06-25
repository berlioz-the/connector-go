package main

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
	"io/ioutil"

	// "github.com/berlioz-the/connector-go"
	"app"
)

func main() {
	time.Sleep(1 * time.Second)

	fmt.Println(berlioz.Request("service", "app", "client").GetRandomEndpoint())

	req := &http.Request{URL: &url.URL{Path: "/"}}
	resp, err := berlioz.Request("service", "app", "client").Do(req)

	fmt.Printf("Response: %s, error: %s\n", resp, err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("Response Body: %s\n", body)

}
