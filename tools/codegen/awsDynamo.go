package main

import (
	"reflect"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	. "github.com/dave/jennifer/jen"
)

func genDynamoDB(f *File) {
	rdynamo := reflect.TypeOf((*dynamodb.DynamoDB)(nil))
	wrapAwsClient(f, rdynamo)
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

func setupDynamoDbPeer(f *File, rm reflect.Method, inputVarName string) Code {
	return Id(inputVarName).Id(".TableName").Op("=").Id("&peer.Name")
}
