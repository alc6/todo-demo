package main

import (
	"io"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegermetrics "github.com/uber/jaeger-lib/metrics"
	"google.golang.org/grpc"
)

const serviceName = "todo_server"

type instrumentedGRPCServer struct {
	grpcServer      *grpc.Server
	openTrace       opentracing.Tracer
	openTraceCloser io.Closer
}

func newInstrumentedGRPCServer() (*instrumentedGRPCServer, error) {
	var igrpc instrumentedGRPCServer

	cfg := jaegercfg.Configuration{
		ServiceName: serviceName,
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans: true,
		},
	}

	var err error

	igrpc.openTrace, igrpc.openTraceCloser, err = cfg.NewTracer(jaegercfg.Metrics(jaegermetrics.NullFactory))
	if err != nil {
		return nil, err
	}

	igrpc.grpcServer = grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_opentracing.UnaryServerInterceptor(grpc_opentracing.WithTracer(igrpc.openTrace)),
		)),
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_opentracing.StreamServerInterceptor(),
		)),
	)

	return &igrpc, nil
}
