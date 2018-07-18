package main

import (
	"reflect"
	"strings"

	"github.com/aws/aws-sdk-go/service/ssm/ssmiface"
)

func genSSM(f *CodeWriter) {
	rSSM := reflect.TypeOf((*ssmiface.SSMAPI)(nil))
	wrapAwsClient(f, rSSM.Elem(), "SSMAccessor")
}

func checkSSMMethod(rm reflect.Method) bool {
	if !strings.Contains(rm.Name, "Parameter") {
		return false
	}

	if strings.Contains(rm.Name, "Parameters") {
		return false
	}

	// rmt := rm.Type
	// if rmt.NumIn() >= 2 {
	// 	in := rmt.In(1)
	// 	if in.Kind() == reflect.Ptr {
	// 		inElem := in.Elem()
	// 		if inElem.Kind() == reflect.Struct {
	// 			if _, ok := inElem.FieldByName("Name"); !ok {
	// 				return false
	// 			}
	// 		} else {
	// 			return false
	// 		}
	// 	} else {
	// 		return false
	// 	}
	// }

	return true
}

func setupSSMPeer(f *CodeWriter, rm reflect.Method, inputVarName string) {
	f.WriteLine("%s.Name = &peer.Name", inputVarName)
}
