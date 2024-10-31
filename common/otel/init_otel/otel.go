package init_otel

import (
	"context"
	"fmt"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/zipkin"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
	"time"
)

func InitOTEL() func(ctx context.Context) {
	res, err := newResource("pulse_community", "v0.0.1")
	if err != nil {
		fmt.Println("Error creating resource:", err)
		panic(err)
	}
	fmt.Println("Resource created successfully")

	prop := newPropagator()
	otel.SetTextMapPropagator(prop)

	tp, err := newTraceProvider(res)
	if err != nil {
		fmt.Println("Error creating trace provider:", err)
		panic(err)
	}
	otel.SetTracerProvider(tp)
	fmt.Println("Tracer provider initialized successfully")

	return func(ctx context.Context) {
		if shutdownErr := tp.Shutdown(ctx); shutdownErr != nil {
			fmt.Println("Error shutting down trace provider:", shutdownErr)
		}
	}
}

func newTraceProvider(res *resource.Resource) (*trace.TracerProvider, error) {
	exporter, err := zipkin.New("http://localhost:9411/api/v2/spans")
	if err != nil {
		fmt.Println("Error creating Zipkin exporter:", err)
		return nil, err
	}
	fmt.Println("Zipkin exporter created successfully")

	traceProvider := trace.NewTracerProvider(
		trace.WithBatcher(exporter, trace.WithBatchTimeout(time.Second)),
		trace.WithResource(res),
	)
	fmt.Println("Trace provider created successfully")
	return traceProvider, nil
}

// InitOTEL 返回一个关闭函数，并且让调用者关闭的时候来决定这个 ctx

func newResource(serviceName, serviceVersion string) (*resource.Resource, error) {
	return resource.Merge(resource.Default(),
		resource.NewWithAttributes(semconv.SchemaURL,
			semconv.ServiceName(serviceName),
			semconv.ServiceVersion(serviceVersion),
		))
}

func newPropagator() propagation.TextMapPropagator {
	return propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	)
}
