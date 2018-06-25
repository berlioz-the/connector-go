package berlioz

import (
	"os"
	"fmt"
	"log"
	"strconv"
	"net/http"
	"encoding/json"

	// "github.com/gorilla/websocket"
	
)

// Peers is TBD.
type Peers interface {
	Do(*http.Request) (*http.Response, error)
	// Get(string) (*http.Response, error)
	GetRandomEndpoint() (Endpoint, error)
}

type peers map[string]Endpoint

// TODO: Shouldn't the pointer implement the interface rather than the value?

func (p peers) GetRandomEndpoint() (Endpoint, error) {
	for _, v := range map[string]Endpoint(p) {
		return v, nil
	}
	return Endpoint{}, fmt.Errorf("No endpoints found")
}

func (p peers) Do(req *http.Request) (*http.Response, error) {
	ep, err := p.GetRandomEndpoint()
	if err != nil {
		return nil, err
	}

	req.URL.Scheme = ep.Protocol
	req.URL.Host = ep.Address + ":" + strconv.Itoa(int(ep.Port))
	// req.URL.Host = "localhost:40002"

	// TODO: Manipulate headers.

	log.Printf("Request: %s", req.URL.String())
	resp, err := new(http.Client).Do(req)

	// TODO: Handle retry.

	return resp, err
}

// func (p *peers) Get(path string) (*http.Response, error) {
// 	ep := p.getRandomEndpoint()

// 	epURL := url.URL{
// 		Scheme: ep.Protocol,
// 		Host:   ep.Address + ":" + string(ep.Port),
// 		Path:   path,
// 	}

// 	req, err := http.NewRequest("GET", epURL.String(), nil)
// 	// Handle error

// 	resp, err := new(http.Client).Get(epURL.String())

// 	// Handle retry.

// 	return resp, err
// }

type agentResponsePeers struct {
	Database map[string]map[string]cloudResource `json:"database,omitempty"`
	Queue    map[string]map[string]cloudResource `json:"queue,omitempty"`
	Service  map[string]map[string]peers         `json:"service,omitempty"`
	Cluster  map[string]map[string]peers         `json:"cluster,omitempty"`
}

type agentResponse struct {
	Endpoints map[string]Endpoint `json:"endpoints,omitempty"`
	Policies  policy              `json:"policies,omitempty"`
	Peers     agentResponsePeers  `json:"peers,omitempty"`
}

// Request is TBD.
func Request(kind string, name string, ep string) Peers {
	if kind == "service" {
		return registry.Peers.Service[name][ep]
	}
	panic(fmt.Sprintf("Unhandled kind: %s", kind))
}

var registry agentResponse

func init() {
	wsURL := os.Getenv("BERLIOZ_AGENT_PATH")
	initClient(wsURL, handleMessage)
}

func handleMessage(message []byte) {

	var resp agentResponse
	err := json.Unmarshal(message, &resp)
	if err != nil {
		log.Printf("[handleMessage] Encountered error while parsing response: %s", err)
		return
	}

	registry = resp
	
	b, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		log.Fatalf("[handleMessage] Encountered error while encoding response back: %s", err)
	}
	fmt.Printf("[handleMessage] Updated registry: %s\n", string(b))
}
