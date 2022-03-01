package main

import (
	"context"

	"github.com/alc6/todo-demo/proto/todorpc"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type TodoGRPCServer struct {
	todorpc.UnimplementedTodosServer

	log *zap.Logger
}

// NewTodoGRPCServer returns new instance of TodoGRPCServer that fully implements
// a 'Todos' gRPC server.
func NewTodoGRPCServer(logger *zap.Logger) *TodoGRPCServer {
	todoGrpcServer := TodoGRPCServer{
		log: logger,
	}

	return &todoGrpcServer
}

/* Below functions satisfy the todorpc.TodosClient interface */

func (server *TodoGRPCServer) CreateTodo(ctx context.Context, in *todorpc.CreateTodoReq, opts ...grpc.CallOption) (*todorpc.CreateTodoResp, error) {
	return nil, nil
}

func (server *TodoGRPCServer) ReadTodo(ctx context.Context, in *todorpc.ReadTodoReq) (*todorpc.ReadTodoResp, error) {
	return nil, nil
}

func (server *TodoGRPCServer) ReadAllTodo(ctx context.Context, in *todorpc.ReadAllTodoReq) (*todorpc.ReadAllTodoResp, error) {
	return nil, nil
}

func (server *TodoGRPCServer) UpdateTodo(ctx context.Context, in *todorpc.UpdateTodoReq) (*todorpc.ReadTodoResp, error) {
	return nil, nil
}

func (server *TodoGRPCServer) DeleteTodo(ctx context.Context, in *todorpc.DeleteTodoReq) (*todorpc.DeleteTodoResp, error) {
	return nil, nil
}

