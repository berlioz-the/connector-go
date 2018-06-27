package berlioz

import (
	"fmt"
	"log"
	"time"
)

// TBD
type execContext struct {
	tryCount int
}

type execActionF func(interface{}) ([]interface{}, error)

func execute(kind string, path []string, action execActionF) ([]interface{}, error) {
	context := new(execContext)
	for {
		res, err := _tryExecute(kind, path, action)
		if err == nil {
			return res, nil
		}

		context.tryCount++
		canRetry := _prepareRetry(context)
		if !canRetry {
			return nil, err
		}
	}
}

func _tryExecute(kind string, path []string, action execActionF) ([]interface{}, error) {
	log.Println("Trying...")

	peersMap := registry.getAsIndexedMap(kind, path)
	peer := peersMap.random()
	if peer == nil {
		return nil, fmt.Errorf("No peer available")
	}

	res, err := action(peer)
	return res, err
}

func _prepareRetry(context *execContext) bool {
	if context.tryCount > 3 {
		return false
	}

	log.Println("Sleeping before retry...")
	time.Sleep(2 * time.Second)
	return true
}
