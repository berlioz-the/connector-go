package main

import (
	"log"
	"reflect"
	"strconv"

	. "github.com/dave/jennifer/jen"
)

var awsMethodCheckCustomHandlers = map[string]func(reflect.Method) bool{
	"*dynamodb.DynamoDB": checkDynamoDbMethod,
}

var awsPackageNames = map[string]string{
	"dynamodb": "github.com/aws/aws-sdk-go/service/dynamodb",
}

var awsPeerHandler = map[string]func(*File, reflect.Method, string) Code{
	"*dynamodb.DynamoDB": setupDynamoDbPeer,
}

func generateAws() {
	f := NewFile("berlioz")

	f.Comment("THIS FILE IS GENERATED USING ROBOT")
	f.Comment("DO NOT MODIFY THIS FILE DIRECTLY")
	genDynamoDB(f)

	err := f.Save("genAws.go")
	if err != nil {
		log.Fatal(err)
	}
}

func wrapAwsClient(f *File, t reflect.Type) {
	for i := 0; i < t.NumMethod(); i++ {
		rm := t.Method(i)
		wrapAwsMethod(f, t, rm)
	}
}

func wrapAwsMethod(f *File, t reflect.Type, rm reflect.Method) {
	if !shouldWrapAwsMethod(t, rm) {
		return
	}

	rmt := rm.Type
	f.Line()
	f.Commentf("Wrapper for %s::%s.", t.String(), rm.Name)

	inputArgs := make([]Code, 0)
	innerArgs := make([]Code, 0)
	outputTypes := make([]Code, 0)
	outputReturnSuccess := make([]Code, 0)
	outputReturnError := make([]Code, 0)

	apiFuncResults := make([]Code, 0)
	apiFuncReturnArray := make([]Code, 0)

	contents := make([]Code, 0)

	resultVariableName := "_"
	var inputStructVariable string

	for i := 1; i < rmt.NumIn(); i++ {
		in := rmt.In(i)
		f.Commentf("  In: %v, Type: %v", in.Name(), in.String())
		argName := "in" + strconv.Itoa(i)
		if len(inputStructVariable) == 0 {
			inputStructVariable = argName
		}

		inputArgType := resolveGeneratorType(in)
		inputArgs = append(inputArgs, Id(argName).Add(inputArgType))

		innerArgs = append(innerArgs, Id(argName))
	}
	for i := 0; i < rmt.NumOut(); i++ {
		out := rmt.Out(i)
		f.Commentf("  Out: %v, Type: %v", out.Name(), out.String())
		outputTypes = append(outputTypes, Id(out.String()))

		var apiFuncVar string
		var outerFuncVar string
		if out.Name() == "error" {
			apiFuncVar = "err"
			outputReturnSuccess = append(outputReturnSuccess, Nil())
			outputReturnError = append(outputReturnError, Id("execErr"))
		} else {
			apiFuncVar = "res[" + strconv.Itoa(len(apiFuncReturnArray)) + "]"
			apiFuncReturnArray = append(apiFuncReturnArray, Id(apiFuncVar))

			resultVariableName = "execRes"
			outerFuncVar = "execRes[" + strconv.Itoa(len(outputReturnSuccess)) + "].(" + out.String() + ")"
			outputReturnSuccess = append(outputReturnSuccess, Id(outerFuncVar))
			outputReturnError = append(outputReturnError, Nil())
		}
		apiFuncResults = append(apiFuncResults, Id(apiFuncVar))
	}

	var peerCustomCode Code
	if peerHandler, ok := awsPeerHandler[t.String()]; ok {
		peerCustomCode = peerHandler(f, rm, inputStructVariable)
	}

	actualCall := Id("action").Op(":=").Func().Params(Id("rawPeer").Id("interface{}")).Params(Id("[]interface{}"), Error()).Block(
		List(Id("client"), Id("peer"), Id("err").Op(":=").Id("x.Client").Call(Id("rawPeer"))),
		If(Id("err").Op("!=").Nil()).Block(
			Return(List(Nil(), Err())),
		),
		peerCustomCode,
		Id("res").Op(":=").Make(Id("[]interface{}"), Id(strconv.Itoa(len(apiFuncReturnArray)))),
		List(apiFuncResults...).Op("=").Id("client."+rm.Name).Call(innerArgs...),
		Return(Id("res"), Err()),
	)

	contents = append(contents, actualCall)

	contents = append(contents, List(Id(resultVariableName), Id("execErr")).Op(":=").Id("execute").Call(
		Id("x.info.kind"),
		Id("x.info.path"),
		Id("action"),
	))

	contents = append(contents, If(Id("execErr").Op("!=").Nil()).Block(
		Return(outputReturnError...),
	))

	contents = append(contents, Return(outputReturnSuccess...))

	f.Func().Params(
		Id("x").Id("DynamoDBAccessor"),
	).Id(rm.Name).Params(inputArgs...).Params(outputTypes...).Block(contents...)
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

	if customChecker, ok := awsMethodCheckCustomHandlers[t.String()]; ok {
		if !customChecker(rm) {
			return false
		}
	}

	return true
}
