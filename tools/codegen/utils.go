package main

import (
	"reflect"
	"strings"
)

type MyTypeInfo struct {
	packageName string
	name        string
	isPtr       bool
}

func (x MyTypeInfo) String() string {
	str := ""
	if x.isPtr {
		str = str + "*"
	}
	if len(x.packageName) > 0 {
		str = str + x.packageName + "."
	}
	str = str + x.name
	return str
}

func parseType(t reflect.Type) MyTypeInfo {
	isPtr := false
	tt := t
	if t.Kind() == reflect.Ptr {
		isPtr = true
		tt = t.Elem()
	}
	arr := strings.Split(tt.String(), ".")
	if len(arr) == 1 {
		return MyTypeInfo{packageName: "", name: arr[0], isPtr: isPtr}
	}
	return MyTypeInfo{packageName: arr[0], name: arr[1], isPtr: isPtr}
}
