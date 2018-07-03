package berlioz

import (
	opentracing "github.com/opentracing/opentracing-go"
	zipkin "github.com/openzipkin/zipkin-go-opentracing"
)

func testZipkin() {
	zipkin.ClientServerSameSpan()
	opentracing.kuku()
}
