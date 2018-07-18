package main

import (
	"fmt"
	"os"
	"strings"
)

type CodeWriter struct {
	_file   *os.File
	_indent int
}

func NewCodeWriter(name string) *CodeWriter {
	f, err := os.Create(name)
	if err != nil {
		return nil
	}
	return &CodeWriter{_file: f}
}

func (x *CodeWriter) Close() error {
	return x._file.Close()
}

func (x *CodeWriter) Indent() *CodeWriter {
	x._indent++
	return x
}

func (x *CodeWriter) Unindent() *CodeWriter {
	x._indent--
	return x
}

func (x *CodeWriter) WriteLine(format string, a ...interface{}) *CodeWriter {
	return x.write("", format, a...)
}

func (x *CodeWriter) NewLine() *CodeWriter {
	return x.write("", "")
}

func (x *CodeWriter) Comment(format string, a ...interface{}) *CodeWriter {
	return x.write("// ", format, a...)
}

func (x *CodeWriter) write(prefix string, format string, a ...interface{}) *CodeWriter {
	line := fmt.Sprintf(format, a...)
	line = strings.Repeat("    ", x._indent) + prefix + line + "\n"
	x._file.WriteString(line)
	return x
}
