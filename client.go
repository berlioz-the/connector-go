package berlioz

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
)

type messageHandler func([]byte) error

func initClient(handler messageHandler) {
	wsURL := GetEnvironmentVariable("BERLIOZ_AGENT_PATH")
	log.Printf("[initClient] AGENT WS URL: %s\n", wsURL)

	go func() {
		for {
			err := _clientTryProcess(wsURL, handler)
			if err != nil {
				log.Printf("[initClient] Error: %v\n", err)
			}

			time.Sleep(5 * time.Second)
		}
	}()
}

func _clientTryProcess(wsURL string, handler messageHandler) error {

	log.Printf("[_clientTryProcess] Start: %s\n", wsURL)

	dialer := websocket.DefaultDialer
	ws, _, err := dialer.Dial(wsURL, nil)
	if err != nil {
		return err
	}

	for {
		// log.Printf("[_clientTryProcess] loop... \n")
		_, message, err := ws.ReadMessage()
		if err != nil {
			return err
		}

		log.Printf("[_clientTryProcess] Message: %s \n", string(message[:]))

		err = handler(message)
		if err != nil {
			log.Printf("[_clientTryProcess] Error processing message. Error: %s. Message: %s \n", err, string(message[:]))
			return err
		}
	}

	log.Printf("[_clientTryProcess] End \n")
	return nil
}
