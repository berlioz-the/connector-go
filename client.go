package berlioz

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
)

type messageHandler func([]byte) error

func initClient(wsURL string, handler messageHandler) {
	log.Printf("[initClient] WS URL: %s\n", wsURL)

	go func() {
		for {
			err := _clientTryProcess(wsURL, handler)
			if err != nil {
				log.Printf("[initClient] Error:", err)
			}

			time.Sleep(2 * time.Second)
		}
	}()
}

func _clientTryProcess(wsURL string, handler messageHandler) error {

	// log.Printf("[_clientTryProcess] Start \n")

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
