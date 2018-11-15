package berlioz

import (
	"context"
	"log"
	"net/http"
	"strings"

	"fmt"
	"os"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	zipkin "github.com/openzipkin/zipkin-go-opentracing"
	"github.com/openzipkin/zipkin-go-opentracing/thrift/gen-go/zipkincore"
	// _ "github.com/apache/thrift"
)

type zipkinInfo struct {
	localServiceName string

	enabled bool
	url     string

	collector *zipkin.Collector
	tracers   map[string]*opentracing.Tracer
}

// TBD
type TracingSpan struct {
	span   *opentracing.Span
	tracer *opentracing.Tracer
}

var myZipkin *zipkinInfo

func initZipkin() {
	myZipkin = newZipkin()
}

func newZipkin() *zipkinInfo {
	x := zipkinInfo{
		localServiceName: os.Getenv("BERLIOZ_CLUSTER") + "-" + os.Getenv("BERLIOZ_SECTOR") + "-" + os.Getenv("BERLIOZ_SERVICE"),
	}
	monitorPolicyBool("enable-zipkin", nil, func(value bool) {
		x.enabled = value
		log.Printf("[ZipkinInfo::monitor] Enabled=%s\n", x.enabled)
		x.activateChanges()
	})
	// TODO: use zipkin-service-id instead
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
	x.tracers = make(map[string]*opentracing.Tracer)
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
}

func (x *zipkinInfo) getTracer(name string) *opentracing.Tracer {
	if tracer, ok := x.tracers[name]; ok {
		return tracer
	}
	collector := x.collector
	if collector == nil {
		return nil
	}

	recorder := zipkin.NewRecorder(*collector, true, "", name)

	tracer, err := zipkin.NewTracer(
		recorder,
		zipkin.ClientServerSameSpan(true),
		zipkin.TraceID128Bit(true),
		// zipkin.WithLogger(logger),
	)
	if err != nil {
		fmt.Printf("unable to create Zipkin tracer: %+v\n", err)
		return nil
	}

	x.tracers[name] = &tracer
	return &tracer
}

func (x *zipkinInfo) instrument(context context.Context, name string, operation string) TracingSpan {
	log.Printf("[ZipkinInfo::instrument] %s\n", name)

	tracer := x.getTracer(name)
	if tracer == nil {
		return TracingSpan{}
	}

	options := make([]opentracing.StartSpanOption, 0)

	var parentCtx opentracing.SpanContext
	parentSpan := opentracing.SpanFromContext(context)
	if parentSpan != nil {
		log.Printf("[ZipkinInfo::instrument] %s. Parent span present.\n", name)

		parentCtx = parentSpan.Context()
		options = append(options, opentracing.ChildOf(parentCtx))
	}

	span := (*tracer).StartSpan(operation, options...)
	log.Printf("[ZipkinInfo::instrument] %s. END\n", name)

	// opentracing.ContextWithSpan(context, span)

	return TracingSpan{span: &span, tracer: tracer}
}

// TBD
func (x TracingSpan) Finish() {
	if x.span == nil {
		return
	}
	(*x.span).Finish()
}

func (x *zipkinInfo) instrumentServerRequest(req *http.Request) (*http.Request, TracingSpan) {
	log.Printf("[ZipkinInfo::instrumentServerRequest] \n")

	tracer := x.getTracer(x.localServiceName)
	if tracer == nil {
		return req, TracingSpan{}
	}
	wireContext, err := (*tracer).Extract(
		opentracing.TextMap,
		opentracing.HTTPHeadersCarrier(req.Header),
	)
	var span opentracing.Span
	if err == nil {
		span = (*tracer).StartSpan(req.Method, ext.RPCServerOption(wireContext))
	} else {
		span = (*tracer).StartSpan(req.Method)
	}

	span.SetTag(zipkincore.HTTP_METHOD, req.Method)
	span.SetTag(zipkincore.HTTP_URL, req.URL.Path)

	ctx := opentracing.ContextWithSpan(req.Context(), span)

	return req.WithContext(ctx), TracingSpan{span: &span}
}
