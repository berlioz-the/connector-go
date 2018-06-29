package berlioz

import (
	"fmt"
	"log"
	"math"
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
		log.Printf("There was error: %s\n", err)
		canRetry := _prepareRetry(context)
		if !canRetry {
			return nil, err
		}
	}
}

func _tryExecute(kind string, path []string, action execActionF) ([]interface{}, error) {
	log.Println("Trying...")
	log.Printf("EnableZipkin = %s\n", resolvePolicy("enable-zipkin", nil))
	log.Printf("ZipkinURL = %s\n", resolvePolicy("zipkin-endpoint", nil))

	peersMap := registry.getAsIndexedMap(kind, path)
	peer := peersMap.random()
	if peer == nil {
		return nil, fmt.Errorf("No peer available")
	}

	res, err := action(peer)
	return res, err
}

func _prepareRetry(context *execContext) bool {
	if context.tryCount > resolvePolicyInt("retry-count", nil) {
		return false
	}
	timeout := resolvePolicyInt("retry-initial-delay", nil)
	timeout = timeout * int(math.Pow(resolvePolicyFloat("retry-delay-multiplier", nil), float64(context.tryCount-1)))
	maxDelay := resolvePolicyInt("retry-max-delay", nil)
	if timeout > maxDelay {
		timeout = maxDelay
	}
	if timeout > 0 {
		log.Printf("Sleeping %dms before retry...\n", timeout)
		time.Sleep(time.Duration(timeout) * time.Millisecond)
	}
	return true
}
