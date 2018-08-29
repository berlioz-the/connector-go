package main

import (
	"log"
	"reflect"
	"strconv"
	"strings"
)

var awsMethodCheckCustomHandlers = map[string]func(reflect.Method) bool{
	"*dynamodb.DynamoDB": checkDynamoDbMethod,
	"ssmiface.SSMAPI":    checkSSMMethod,
}

var awsPeerHandler = map[string]func(*CodeWriter, reflect.Method, string){
	"*dynamodb.DynamoDB": setupDynamoDbPeer,
	"ssmiface.SSMAPI":    setupSSMPeer,
}

func generateAws() {
	f := NewCodeWriter("genAws.go")

	f.WriteLine("package berlioz")
	f.NewLine()

	f.WriteLine("import (")
	f.Indent()
	f.WriteLine("\"context\"")
	f.WriteLine("\"github.com/aws/aws-sdk-go/service/dynamodb\"")
	f.WriteLine("\"github.com/aws/aws-sdk-go/service/ssm\"")
	f.Unindent()
	f.WriteLine(")")
	f.NewLine()

	f.Comment("THIS FILE IS GENERATED USING ROBOT")
	f.Comment("DO NOT MODIFY THIS FILE DIRECTLY")
	f.NewLine()

	genDynamoDB(f)
	genSSM(f)

	err := f.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func wrapAwsClient(f *CodeWriter, t reflect.Type, accessorTypeName string) {
	var name string
	if t.Kind() == reflect.Interface {
		name = t.Name()
	} else {
		name = t.Elem().Name()
	}

	log.Printf("[wrapAwsClient] %s\n", name)
	log.Printf("[wrapAwsClient] %s, kind:%s\n", name, t.Kind())
	log.Printf("[wrapAwsClient] %s, methods:%d\n", name, t.NumMethod())
	for i := 0; i < t.NumMethod(); i++ {
		rm := t.Method(i)
		wrapAwsMethod(f, name, accessorTypeName, t, rm)
	}
}

func wrapAwsMethod(f *CodeWriter, name string, accessorTypeName string, t reflect.Type, rm reflect.Method) {
	log.Printf("[wrapAwsMethod] %s::%s\n", name, rm.Name)
	if !shouldWrapAwsMethod(t, rm) {
		return
	}

	f.NewLine()
	f.Comment("Wrapper for %s::%s.", t.String(), rm.Name)

	rmt := rm.Type

	inputArgs := make([]string, 0)
	innerArgs := make([]string, 0)
	outputTypes := make([]string, 0)
	outputReturnSuccess := make([]string, 0)
	outputReturnError := make([]string, 0)

	apiFuncResults := make([]string, 0)
	apiFuncReturnArray := make([]string, 0)

	resultVariableName := "_"
	var inputStructVariable string

	inputArgs = append(inputArgs, "ctx context.Context")

	startIndex := 0
	if t.Kind() == reflect.Ptr {
		if t.Elem().Kind() == reflect.Struct {
			startIndex = 1
		}
	} else if t.Kind() == reflect.Struct {
		startIndex = 1
	}

	for i := startIndex; i < rmt.NumIn(); i++ {
		in := rmt.In(i)
		f.Comment("  In: %v, Type: %v", in.Name(), in.String())
		argName := "in" + strconv.Itoa(i)
		if len(inputStructVariable) == 0 {
			inputStructVariable = argName
		}

		inputArgType := in.String()
		inputArgs = append(inputArgs, argName+" "+inputArgType)
		innerArgs = append(innerArgs, argName)
	}
	for i := 0; i < rmt.NumOut(); i++ {
		out := rmt.Out(i)
		f.Comment("  Out: %v, Type: %v", out.Name(), out.String())
		outputTypes = append(outputTypes, out.String())

		var apiFuncVar string
		var outerFuncVar string
		if out.Name() == "error" {
			apiFuncVar = "err"
			outputReturnSuccess = append(outputReturnSuccess, "nil")
			outputReturnError = append(outputReturnError, "execErr")
		} else {
			apiFuncVar = "res[" + strconv.Itoa(len(apiFuncReturnArray)) + "]"
			apiFuncReturnArray = append(apiFuncReturnArray, apiFuncVar)

			resultVariableName = "execRes"
			outerFuncVar = "execRes[" + strconv.Itoa(len(outputReturnSuccess)) + "].(" + out.String() + ")"
			outputReturnSuccess = append(outputReturnSuccess, outerFuncVar)
			outputReturnError = append(outputReturnError, "nil")
		}
		apiFuncResults = append(apiFuncResults, apiFuncVar)
	}

	f.WriteLine("func (x %s) %s(%s) (%s) {", accessorTypeName, rm.Name, strings.Join(inputArgs, ", "), strings.Join(outputTypes, ", "))
	f.Indent()

	f.WriteLine("action := func(rawPeer interface{}, span *TracingSpan) ([]interface{}, error) {")
	f.Indent()
	f.WriteLine("client, peer, err := x.Client(rawPeer)")
	f.WriteLine("if err != nil {").Indent()
	f.WriteLine("return nil, err").Unindent()
	f.WriteLine("}")
	if peerHandler, ok := awsPeerHandler[t.String()]; ok {
		peerHandler(f, rm, inputStructVariable)
	}
	f.WriteLine("res := make([]interface{}, %d)", len(apiFuncReturnArray))
	f.WriteLine("%s = client.%s(%s)", strings.Join(apiFuncResults, ", "), rm.Name, strings.Join(innerArgs, ", "))
	f.WriteLine("return res, err")
	f.Unindent()
	f.WriteLine("}")

	f.WriteLine("%s, execErr := execute(ctx, x.info.peers, \"%s\", action)", resultVariableName, rm.Name)
	f.WriteLine("if execErr != nil {").Indent()
	f.WriteLine("return %s", strings.Join(outputReturnError, ", ")).Unindent()
	f.WriteLine("}")

	f.WriteLine("return %s", strings.Join(outputReturnSuccess, ", "))

	f.Unindent()
	f.WriteLine("}")
}

func shouldWrapAwsMethod(t reflect.Type, rm reflect.Method) bool {
	rmt := rm.Type

	if rm.Name == "AddDebugHandlers" {
		return false
	}
	if rm.Name == "MaxRetries" {
		return false
	}

	if rmt.NumIn() >= 2 {
		in := rmt.In(1)
		if in.Kind() == reflect.Ptr {
			inElem := in.Elem()
			if inElem.Kind() != reflect.Struct {
				return false
			}
		} else {
			return false
		}
	}

	for i := 0; i < rmt.NumIn(); i++ {
		in := rmt.In(i)
		typeName := in.String()
		if typeName == "aws.Context" {
			return false
		}
		if typeName == "*request.Request" {
			return false
		}
	}
	for i := 0; i < rmt.NumOut(); i++ {
		out := rmt.Out(i)
		typeName := out.String()
		if typeName == "*request.Request" {
			return false
		}
	}

	log.Printf("[shouldWrapAwsMethod] %s...\n", t.String())
	if customChecker, ok := awsMethodCheckCustomHandlers[t.String()]; ok {
		if !customChecker(rm) {
			return false
		}
	}

	return true
}
