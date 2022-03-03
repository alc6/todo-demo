package main

import (
	"net"
	"time"

	"github.com/alc6/todo-demo/proto/todorpc"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const defaultGRPCPort = ":27015"

func main() {
	logger := NewZapLogger()

	logger.Info("Starting server...")

	todoServices := NewTodosGRPCService(logger)

	grpcServer := grpc.NewServer()

	todorpc.RegisterTodosServer(grpcServer, todoServices)

	stopGRPCServer, errStartServ := startGRPCServer(logger, grpcServer, defaultGRPCPort, true)
	if errStartServ != nil {
		logger.Sugar().Panicw("Failed to start the server", "error", errStartServ)
	}

	time.Sleep(5 * time.Second)

	stopGRPCServer()

	time.Sleep(200 * time.Millisecond)

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
