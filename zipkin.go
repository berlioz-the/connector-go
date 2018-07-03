package berlioz

// opentracing "github.com/opentracing/opentracing-go"
// zipkin "github.com/openzipkin/zipkin-go-opentracing"
import (
	_ "github.com/opentracing/opentracing-go"
	// _ "github.com/apache/thrift"
)

func testZipkin() {
	// collector, err := zipkin.NewHTTPCollector("http://localhost:9411/api/v1/spans")
	// if err != nil {
	// 	fmt.Printf("unable to create Zipkin HTTP collector: %+v\n", err)
	// 	os.Exit(-1)
	// }

	// // create recorder.
	// recorder := zipkin.NewRecorder(collector, false, "127.0.0.1:61001", "serviceName")

	// // create tracer.
	// ///tracer
	// trader, err := zipkin.NewTracer(
	// 	recorder,
	// 	zipkin.ClientServerSameSpan(true),
	// 	zipkin.TraceID128Bit(true),
	// )
	// if err != nil {
	// 	fmt.Printf("unable to create Zipkin tracer: %+v\n", err)
	// 	os.Exit(-1)
	// }
}
