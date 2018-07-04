package berlioz

import (
	"context"
	"log"
	"strings"
	"time"

	"fmt"
	"os"

	opentracing "github.com/opentracing/opentracing-go"
	zipkin "github.com/openzipkin/zipkin-go-opentracing"
	// _ "github.com/apache/thrift"
)

type zipkinInfo struct {
	enabled bool
	url     string

	collector *zipkin.Collector
	tracer    *opentracing.Tracer
}

type TracingSpan struct {
	span *opentracing.Span
}

var myZipkin *zipkinInfo

func initZipkin() {
	myZipkin = newZipkin()
}

func newZipkin() *zipkinInfo {
	x := zipkinInfo{}
	monitorPolicyBool("enable-zipkin", nil, func(value bool) {
		x.enabled = value
		log.Printf("[ZipkinInfo::monitor] Enabled=%s\n", x.enabled)
		x.activateChanges()
	})
	monitorPolicyString("zipkin-endpoint", nil, func(value string) {
		x.url = strings.Replace(value, "v2", "v1", -1)
		log.Printf("[ZipkinInfo::monitor] URL=%s\n", x.url)
		x.activateChanges()
	})
	return &x
}

func (x *zipkinInfo) activateChanges() {
	log.Printf("[ZipkinInfo::activateChanges] \n")

	x.cleanup()
	if x.enabled && len(x.url) > 0 {
		x.setup()
	}
}

func (x *zipkinInfo) cleanup() {
	log.Printf("[ZipkinInfo::cleanup] \n")

	collector := x.collector
	if collector != nil {
		x.collector = nil
		(*collector).Close()
	}
	x.tracer = nil
}

func (x *zipkinInfo) setup() {
	log.Printf("[ZipkinInfo::setup] \n")

	logger1 := log.New(os.Stdout, log.Prefix(), log.Flags())
	logger := zipkin.LogWrapper(logger1)

	collector, err := zipkin.NewHTTPCollector(x.url, zipkin.HTTPLogger(logger))
	if err != nil {
		fmt.Printf("unable to create Zipkin HTTP collector: %+v\n", err)
		return
	}
	log.Printf("[TestZipkin] collector created \n")
	x.collector = &collector

	// create recorder.
	hostPort := os.Getenv("BERLIOZ_ADDRESS")
	serviceName := os.Getenv("BERLIOZ_CLUSTER") + "-" + os.Getenv("BERLIOZ_SERVICE")
	recorder := zipkin.NewRecorder(collector, true, hostPort, serviceName)
	log.Printf("[TestZipkin] recorder created \n")

	// create tracer.
	tracer, err := zipkin.NewTracer(
		recorder,
		zipkin.ClientServerSameSpan(true),
		zipkin.TraceID128Bit(true),
		zipkin.WithLogger(logger),
	)
	if err != nil {
		fmt.Printf("unable to create Zipkin tracer: %+v\n", err)
		return
	}
	x.tracer = &tracer
}

func (x *zipkinInfo) instrument() TracingSpan {
	log.Printf("[ZipkinInfo::instrument] \n")

	tracer := x.tracer
	if tracer == nil {
		return TracingSpan{}
	}
	span := (*tracer).StartSpan("MyOp1")
	log.Printf("[TestZipkin] span started \n")
	opentracing.ContextWithSpan(context.Background(), span)

	return TracingSpan{span: &span}
}

// TBD
func (x TracingSpan) Finish() {
	if x.span == nil {
		return
	}
	(*x.span).Finish()
}

// TBD
func TestZipkin() {

	time.Sleep(500 * time.Millisecond)

	span := myZipkin.instrument()

	time.Sleep(200 * time.Millisecond)

	span.Finish()

	log.Printf("[TestZipkin] end \n")

	// tracer
	time.Sleep(500 * time.Millisecond)
}
