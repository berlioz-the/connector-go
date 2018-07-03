package berlioz

// opentracing "github.com/opentracing/opentracing-go"
// zipkin "github.com/openzipkin/zipkin-go-opentracing"
import (
	"context"
	"log"
	"time"

	"fmt"
	"os"

	opentracing "github.com/opentracing/opentracing-go"

	zipkin "github.com/openzipkin/zipkin-go-opentracing"
	// _ "github.com/apache/thrift"
)

func TestZipkin() {

	time.Sleep(1 * time.Second)

	log.Printf("[TestZipkin] init \n")

	logger1 := log.New(os.Stdout, log.Prefix(), log.Flags())
	logger := zipkin.LogWrapper(logger1)
	logger.Log("AAAAAAAAAA")
	logger1.Println("BBBBBBBBB")

	collector, err := zipkin.NewHTTPCollector("http://172.17.0.2:9411/api/v1/spans", zipkin.HTTPLogger(logger))
	if err != nil {
		fmt.Printf("unable to create Zipkin HTTP collector: %+v\n", err)
		os.Exit(-1)
	}
	log.Printf("[TestZipkin] collector created \n")

	// create recorder.
	recorder := zipkin.NewRecorder(collector, true, "127.0.0.1:61001", "serviceName")
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
		os.Exit(-1)
	}

	log.Printf("[TestZipkin] tracer created \n")

	span := tracer.StartSpan("MyOp1")
	log.Printf("[TestZipkin] span started \n")
	opentracing.ContextWithSpan(context.Background(), span)

	time.Sleep(500 * time.Millisecond)
	span.Finish()

	log.Printf("[TestZipkin] end \n")

	// tracer
	time.Sleep(500 * time.Millisecond)

	// collector.Collect(&span)
	collector.Close()

}
