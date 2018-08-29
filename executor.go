package berlioz

import (
	"context"
	"fmt"
	"log"
	"math"
	"time"
)

// TBD
type execInfo struct {
	tryCount int
}

type execActionF func(interface{}, *TracingSpan) ([]interface{}, error)

func execute(ctx context.Context, peers PeerAccessor, operationName string, action execActionF) ([]interface{}, error) {
	// log.Printf("[EXECUTE]: kind: %s, path: %v\n", kind, path)

	info := new(execInfo)
	for {
		res, err := _tryExecute(ctx, peers, operationName, action)
		if err == nil {
			return res, nil
		}

		info.tryCount++
		log.Printf("There was error: %s\n", err)
		canRetry := _prepareRetry(info)
		if !canRetry {
			return nil, err
		}
	}
}

func _tryExecute(ctx context.Context, peers PeerAccessor, operationName string, action execActionF) ([]interface{}, error) {
	log.Printf("Trying %s...\n", peers.serviceID)
	// log.Printf("EnableZipkin = %s\n", resolvePolicy("enable-zipkin", nil))
	// log.Printf("ZipkinURL = %s\n", resolvePolicy("zipkin-endpoint", nil))

	peer := peers.Random()
	if peer == nil {
		// log.Printf("REGISTRY: %#v\n", registry)
		return nil, fmt.Errorf("No peer available")
	}

	// naming := []string{}
	// if _, ok := peer.(EndpointModel); ok {
	// 	switch kind {
	// 	case "service":
	// 		naming = append(naming, os.Getenv("BERLIOZ_CLUSTER"))
	// 		naming = append(naming, path[0])
	// 	case "cluster":
	// 		naming = append(naming, path[0])
	// 		naming = append(naming, path[1])
	// 	}
	// } else if cloudPeer, ok := peer.(CloudResourceModel); ok {
	// 	naming = append(naming, os.Getenv("BERLIOZ_CLUSTER"))
	// 	naming = append(naming, cloudPeer.SubClass)
	// 	naming = append(naming, path[0])
	// } else {
	// 	naming = append(naming, kind)
	// 	naming = append(naming, path...)
	// }

	// remoteServiceName := strings.Join(naming, "-")
	// TODO
	remoteServiceName := "MYSERVICE"

	span := myZipkin.instrument(ctx, remoteServiceName, operationName)
	defer span.Finish()
	res, err := action(peer, &span)
	return res, err
}

func _prepareRetry(info *execInfo) bool {
	if info.tryCount > resolvePolicyInt("retry-count", nil) {
		return false
	}
	timeout := resolvePolicyInt("retry-initial-delay", nil)
	timeout = timeout * int(math.Pow(resolvePolicyFloat("retry-delay-multiplier", nil), float64(info.tryCount-1)))
	maxDelay := resolvePolicyInt("retry-max-delay", nil)
	if timeout > maxDelay {
		timeout = maxDelay
	}
	timeout = 500
	if timeout > 0 {
		log.Printf("Sleeping %dms before retry...\n", timeout)
		time.Sleep(time.Duration(timeout) * time.Millisecond)
	}
	return true
}
