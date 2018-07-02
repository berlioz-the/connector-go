package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	. "github.com/dave/jennifer/jen"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("CWD: %s\n", dir)

	f := generate()
	// fmt.Printf("%#v\n", f)

	err = f.Save("genAws.go")
	if err != nil {
		log.Fatal(err)
	}
}

func generate() *File {
	f := NewFile("berlioz")
	f.Id("import \"github.com/aws/aws-sdk-go/service/dynamodb\"")
	// f.ImportName("github.com/aws/aws-sdk-go/service/dynamodb", "zz")
	// f.Func().Id("main").Params().Block(
	// f.Qual("github.com/aws/aws-sdk-go/service/dynamodb", "kuku"),
	// )

	f.Comment("THIS FILE IS GENERATED USING ROBOT")
	f.Comment("DO NOT MODIFY THIS FILE DIRECTLY")
	genDynamoDB(f)

	return f
}

func genDynamoDB(f *File) {

	// (*dynamodb.DynamoDB)(nil).CreateBackup
	rdynamo := reflect.TypeOf((*dynamodb.DynamoDB)(nil))
	for i := 0; i < rdynamo.NumMethod(); i++ {
		rm := rdynamo.Method(i)
		genDynamoDBMethod(f, rm)
	}
	// x := {}

	// f.Func().Id("main").Params().Block(
	// 	Qual("fmt", "Println").Call(Lit("Hello, world")),
	// )
}

func genDynamoDBMethod(f *File, rm reflect.Method) {
	if !shouldWrapDynamoDBMethod(rm) {
		return
	}

	rmt := rm.Type
	f.Line()
	f.Commentf("Wrapper for DynamoDB %s.", rm.Name)

	inputArgs := make([]Code, 0)
	innerArgs := make([]Code, 0)
	outputTypes := make([]Code, 0)
	outputReturnSuccess := make([]Code, 0)
	outputReturnError := make([]Code, 0)

	apiFuncResults := make([]Code, 0)
	apiFuncReturnArray := make([]Code, 0)

	contents := make([]Code, 0)

	for i := 1; i < rmt.NumIn(); i++ {
		in := rmt.In(i)
		f.Commentf("  In: %v, Type: %v", in.Name(), in.String())
		argName := "in" + strconv.Itoa(i)
		inputArgs = append(inputArgs, Id(argName).Id(in.String()))
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

			outerFuncVar = "execRes[" + strconv.Itoa(len(outputReturnSuccess)) + "]"
			outputReturnSuccess = append(outputReturnSuccess, Id(outerFuncVar))
			outputReturnError = append(outputReturnError, Nil())
		}
		apiFuncResults = append(apiFuncResults, Id(apiFuncVar))
	}

	actualCall := Id("action").Op(":=").Func().Params(Id("peer").Id("interface{}")).Params(Id("[]interface{}"), Error()).Block(
		List(Id("client"), Id("err").Op(":=").Id("x.Client").Call(Id("peer"))),
		If(Id("err").Op("!=").Nil()).Block(
			Return(List(Nil(), Err())),
		),
		Id("res").Op(":=").Make(Id("[]interface{}"), Id(strconv.Itoa(len(apiFuncReturnArray)))),
		List(apiFuncResults...).Op("=").Id("client."+rm.Name).Call(innerArgs...),
		Return(Id("res"), Err()),
	)

	contents = append(contents, actualCall)

	contents = append(contents, List(Id("execRes"), Id("execErr")).Op(":=").Id("execute").Call(
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

func shouldWrapDynamoDBMethod(rm reflect.Method) bool {
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
	return true
}
