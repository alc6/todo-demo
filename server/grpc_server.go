package main

import (
	"context"

	"google.golang.org/grpc/codes"

	"google.golang.org/grpc/status"

	"github.com/alc6/todo-demo/server/store"

	"github.com/alc6/todo-demo/proto/todorpc"
	"go.uber.org/zap"
)

type TodoGRPCServiceHandler struct {
	todorpc.UnimplementedTodosServer

	log       *zap.Logger
	todoStore store.Storer
}

// NewTodosGRPCService returns new instance of TodoGRPCServiceHandler that fully implements
// a 'Todos' gRPC interface.
func NewTodosGRPCService(logger *zap.Logger, todoStore store.Storer) *TodoGRPCServiceHandler {
	todoGrpcServer := TodoGRPCServiceHandler{
		log:       logger,
		todoStore: todoStore,
	}

	return &todoGrpcServer
}

/* Below functions satisfy the todorpc.TodosClient interface */

func (h *TodoGRPCServiceHandler) CreateTodo(_ context.Context, in *todorpc.CreateTodoReq) (*todorpc.CreateTodoResp, error) {
	todo := store.TodoFromGRPCStruct(in.GetTodo())

	id, err := h.todoStore.CreateTodo(todo)
	if err != nil {
		return &todorpc.CreateTodoResp{}, status.Errorf(codes.Internal, err.Error())
	}

	resp := todorpc.CreateTodoResp{
		Id: id,
	}

	return &resp, nil
}

func (h *TodoGRPCServiceHandler) ReadTodo(_ context.Context, in *todorpc.ReadTodoReq) (*todorpc.ReadTodoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadAllTodo not implemented")
}

func (h *TodoGRPCServiceHandler) ReadAllTodo(ctx context.Context, in *todorpc.ReadAllTodoReq) (*todorpc.ReadAllTodoResp, error) {
	todos, err := h.todoStore.ReadTodos([]string{}, false)
	if err != nil {
		return &todorpc.ReadAllTodoResp{}, status.Errorf(codes.Unimplemented, err.Error())
	}

	resp := todorpc.ReadAllTodoResp{
		Todo: make([]*todorpc.TodoWithMeta, 0, len(todos)),
	}

	for _, todo := range todos {
		resp.Todo = append(resp.Todo, todo.TodoToGRPCStruct())
	}

	return &resp, nil
}

func (h *TodoGRPCServiceHandler) UpdateTodo(ctx context.Context, in *todorpc.UpdateTodoReq) (*todorpc.UpdateTodoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadAllTodo not implemented")
}

func (h *TodoGRPCServiceHandler) DeleteTodo(ctx context.Context, in *todorpc.DeleteTodoReq) (*todorpc.DeleteTodoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadAllTodo not implemented")
}
