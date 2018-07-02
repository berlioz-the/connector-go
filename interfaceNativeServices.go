package berlioz

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// TBD
type NativeResourceAccessor struct {
	kind string
	path []string
}

func getNativeResource(kind string, name string) NativeResourceAccessor {
	path := make([]string, 1)
	path[0] = name
	return NativeResourceAccessor{kind: kind, path: path}
}

// TBD
func Database(name string) NativeResourceAccessor {
	return getNativeResource("database", name)
}

func (x NativeResourceAccessor) getMap() indexedMap {
	return registry.getAsIndexedMap(x.kind, x.path)
}

// TBD
func (x NativeResourceAccessor) Monitor(callback func(NativeResourceAccessor)) {
	registry.subscribe(x.kind, x.path, func(interface{}) {
		callback(x)
	})
}

// TBD
func (x NativeResourceAccessor) All() CloudResourcesModel {
	y := x.getMap()
	result := CloudResourcesModel{}
	for k, v := range y.all() {
		result[k] = v.(CloudResourceModel)
	}
	return result
}

// TBD
func (x NativeResourceAccessor) Get(identity string) (CloudResourceModel, bool) {
	y := x.getMap()
	val := y.get(identity)
	if val == nil {
		return CloudResourceModel{}, false
	}
	return val.(CloudResourceModel), true
}

// TBD
func (x NativeResourceAccessor) Random() (CloudResourceModel, bool) {
	y := x.getMap()
	val := y.random()
	if val == nil {
		return CloudResourceModel{}, false
	}
	return val.(CloudResourceModel), true
}

// TBD
func (x NativeResourceAccessor) Test() {
	log.Printf("[TEST]...\n")

	sess, err := x.getSession()
	if err != nil {
		log.Printf("[TEST] Could not create session. Reason: %v\n", err)
		return
	}
	svc := dynamodb.New(sess)
	params := &dynamodb.ScanInput{
		TableName: aws.String("localHomePC-berliozgo-contacts"),
	}
	result, err := svc.Scan(params)
	if err != nil {
		log.Printf("[TEST] Scan Error: %v\n", err)
	} else {
		log.Printf("[TEST] Scan Result: %v\n", result)
	}
}

// TBD
type DynamoDBAccessor struct {
	info NativeResourceAccessor
}

// TBD
func (x NativeResourceAccessor) DynamoDB() DynamoDBAccessor {
	return DynamoDBAccessor{info: x}
}

// TBD
func (x DynamoDBAccessor) Client(peer interface{}) (*dynamodb.DynamoDB, error) {
	sess, err := x.info.getSession()
	if err != nil {
		log.Printf("[TEST] Could not create session. Reason: %v\n", err)
		return nil, err
	}
	svc := dynamodb.New(sess)
	return svc, nil
}

func kuku() {
	params := &dynamodb.ScanInput{
		TableName: aws.String("localHomePC-berliozgo-contacts"),
	}
	result, err := Database("").DynamoDB().Scan(params)
	if err != nil {
		log.Printf("[TEST] Scan Error: %v\n", err)
	} else {
		log.Printf("[TEST] Scan Result: %v\n", result)
	}
}

func (x NativeResourceAccessor) getSession() (*session.Session, error) {
	peer, ok := x.Random()
	if !ok {
		return nil, fmt.Errorf("No PEER PRESENT")
	}
	credentials := credentials.NewStaticCredentials(peer.Config.Credentials.AccessKeyID,
		peer.Config.Credentials.SecretAccessKey,
		"")
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(peer.Config.Region),
		Credentials: credentials,
	})
	return sess, err
}
