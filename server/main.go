package main

import (
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/alc6/todo-demo/server/store"

	"github.com/alc6/todo-demo/proto/todorpc"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	defaultGRPCPort = ":27015"
	gRPCReflect     = true
)

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	logger := NewZapLogger()

	logger.Info("Starting server...")

	todoStore := store.NewMapStore(logger)

	todoServices := NewTodosGRPCService(logger, todoStore)

	iServer, errNewGrpcServ := newInstrumentedGRPCServer()
	if errNewGrpcServ != nil {
		logger.Sugar().Panicw("Failed to create instrumented grpc server", "error", errNewGrpcServ)
	}

	todorpc.RegisterTodosServer(iServer.grpcServer, todoServices)

	stopGRPCServer, errStartServ := startGRPCServer(logger, iServer.grpcServer, defaultGRPCPort, gRPCReflect)
	if errStartServ != nil {
		logger.Sugar().Panicw("Failed to start the server", "error", errStartServ)
	}

	<-sigs

	stopGRPCServer()
	iServer.openTraceCloser.Close()

	time.Sleep(500 * time.Millisecond) // Give some time to the goroutines to stop properly.
}

// startGRPCServer starts the provided server with previously associated services.
// Reflection can be enabled through reflect flag.
// Function return allows to gracefully shutdown gRPC server and its underlying net.Listener.
func startGRPCServer(logger *zap.Logger, server *grpc.Server, port string, reflect bool) (func(), error) {
	logger.Sugar().Infof("Setting up gRPC server on port %s with reflection=%v", port, reflect)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		return nil, err
	}

	if reflect {
		reflection.Register(server)
	}

	go func(logger *zap.Logger) {
		logger.Info("Starting gRPC server")

		if err := server.Serve(lis); err != nil {
			logger.Sugar().Errorw("Failed to serve gRPC server", "error", err)
		}
	}(logger)

	shutdownFunc := func() {
		server.GracefulStop()
	}

	return shutdownFunc, nil
}
