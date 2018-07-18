package main

import (
	"reflect"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func genDynamoDB(f *CodeWriter) {
	rdynamo := reflect.TypeOf((*dynamodb.DynamoDB)(nil))
	wrapAwsClient(f, rdynamo, "DynamoDBAccessor")
}

func checkDynamoDbMethod(rm reflect.Method) bool {
	rmt := rm.Type
	if rmt.NumIn() >= 2 {
		in := rmt.In(1)
		if in.Kind() == reflect.Ptr {
			inElem := in.Elem()
			if inElem.Kind() == reflect.Struct {
				if _, ok := inElem.FieldByName("TableName"); !ok {
					return false
				}
			} else {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

func setupDynamoDbPeer(f *CodeWriter, rm reflect.Method, inputVarName string) {
	f.WriteLine("%s.TableName = &peer.Name", inputVarName)
}
