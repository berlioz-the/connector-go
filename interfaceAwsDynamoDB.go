package berlioz

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// TBD
type DynamoDBAccessor struct {
	info NativeResourceAccessor
}

// TBD
func (x NativeResourceAccessor) DynamoDB() DynamoDBAccessor {
	return DynamoDBAccessor{info: x}
}

// TBD
func (x DynamoDBAccessor) Client(rawPeer interface{}) (*dynamodb.DynamoDB, *CloudResourceModel, error) {
	peer, ok := rawPeer.(CloudResourceModel)
	if !ok {
		log.Printf("[TEST] Could not convert peer.")
		return nil, nil, fmt.Errorf("Invalid peer provided.")
	}

	sess, err := x.info.getSession(peer)
	if err != nil {
		log.Printf("[TEST] Could not create session. Reason: %v\n", err)
		return nil, nil, err
	}
	svc := dynamodb.New(sess)
	return svc, &peer, nil
}
