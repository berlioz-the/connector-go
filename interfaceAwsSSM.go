package berlioz

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/service/ssm"
)

// TBD
type SSMAccessor struct {
	info NativeResourceAccessor
}

// TBD
func (x NativeResourceAccessor) SSM() SSMAccessor {
	return SSMAccessor{info: x}
}

// TBD
func (x SSMAccessor) Client(rawPeer interface{}) (*ssm.SSM, *CloudResourceModel, error) {
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
	svc := ssm.New(sess)
	return svc, &peer, nil
}
