package main

import (
	"context"

	"github.com/alc6/todo-demo/proto/todorpc"
	"go.uber.org/zap"
)

type TodoGRPCServiceHandler struct {
	todorpc.UnimplementedTodosServer

	log *zap.Logger
}

// NewTodosGRPCService returns new instance of TodoGRPCServiceHandler that fully implements
// a 'Todos' gRPC interface.
func NewTodosGRPCService(logger *zap.Logger) *TodoGRPCServiceHandler {
	todoGrpcServer := TodoGRPCServiceHandler{
		log: logger,
	}

	return &todoGrpcServer
}

/* Below functions satisfy the todorpc.TodosClient interface */

func (h *TodoGRPCServiceHandler) CreateTodo(ctx context.Context, in *todorpc.CreateTodoReq) (*todorpc.CreateTodoResp, error) {
	return nil, nil
}

func (h *TodoGRPCServiceHandler) ReadTodo(ctx context.Context, in *todorpc.ReadTodoReq) (*todorpc.ReadTodoResp, error) {
	return nil, nil
}

func (h *TodoGRPCServiceHandler) ReadAllTodo(ctx context.Context, in *todorpc.ReadAllTodoReq) (*todorpc.ReadAllTodoResp, error) {
	return nil, nil
}

func (h *TodoGRPCServiceHandler) UpdateTodo(ctx context.Context, in *todorpc.UpdateTodoReq) (*todorpc.UpdateTodoResp, error) {
	return nil, nil
}

func (h *TodoGRPCServiceHandler) DeleteTodo(ctx context.Context, in *todorpc.DeleteTodoReq) (*todorpc.DeleteTodoResp, error) {
	return nil, nil
}
