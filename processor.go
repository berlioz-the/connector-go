package berlioz

import (
	"encoding/json"
	"log"
)

func processMessage(data []byte) error {

	var message agentMessageModel
	err := json.Unmarshal(data, &message)
	if err != nil {
		log.Printf("[handleMessage] Encountered error while parsing response: %s", err)
		return err
	}

	if message.Endpoints != nil {

	}

	if message.Peers != nil {
		if message.Peers.Service != nil {
			for name, endpointMap := range message.Peers.Service {
				for endpoint, value := range endpointMap {
					path := make([]string, 2)
					path[0] = name
					path[1] = endpoint
					registry.setAsIndexedMap("service", path, value)
				}
			}
		}
		if message.Peers.Database != nil {
			for name, value := range message.Peers.Database {
				path := make([]string, 1)
				path[0] = name
				registry.setAsIndexedMap("database", path, value)
			}
		}
		if message.Peers.Queue != nil {
			for name, value := range message.Peers.Queue {
				path := make([]string, 1)
				path[0] = name
				registry.setAsIndexedMap("queue", path, value)
			}
		}
	}

	if message.Policies != nil {
		registry.set("policies", nil, message.Policies)
	}

	registry.debugOutput()
	return nil
}
