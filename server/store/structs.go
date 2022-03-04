package store

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	"google.golang.org/protobuf/types/known/durationpb"

	"github.com/alc6/todo-demo/proto/todorpc"
)

// Todo tasks you expect to... do.
type Todo struct {
	ID            string        `json:"id,omitempty"`
	Title         string        `json:"title,omitempty"`
	Description   string        `json:"description,omitempty"`
	Deadline      time.Time     `json:"deadline,omitempty"`
	Assignee      string        `json:"assignee,omitempty"`
	TimeAllocated time.Duration `json:"time_allocated,omitempty"`
	Status        TodoStatus    `json:"status,omitempty"`
}

type TodoStatus int32

// nolint: revive, stylecheck
const (
	TODO_STATUS_PENDING  TodoStatus = 0
	TODO_STATUS_DOING    TodoStatus = 1
	TODO_STATUS_DONE     TodoStatus = 2
	TODO_STATUS_CANCELED TodoStatus = 3
	TODO_STATUS_EXPIRED  TodoStatus = 4
)

// TodoToGRPCStruct converts to a *todorpc.TodoWithMeta.
func (t *Todo) TodoToGRPCStruct() *todorpc.TodoWithMeta {
	grpcStruct := todorpc.TodoWithMeta{
		Id:            t.ID,
		Title:         t.Title,
		Description:   t.Description,
		Deadline:      timestamppb.New(t.Deadline),
		Assignee:      t.Assignee,
		TimeAllocated: durationpb.New(t.TimeAllocated),
		Status:        todorpc.TodoStatus(t.Status),
	}

	return &grpcStruct
}

// TodoFromGRPCStruct convert a todoGRPCStruct to a Todo.
func TodoFromGRPCStruct(todoGRPCStruct *todorpc.TodoWithMeta) *Todo {
	todo := Todo{
		ID:            todoGRPCStruct.GetId(),
		Title:         todoGRPCStruct.GetTitle(),
		Description:   todoGRPCStruct.GetDescription(),
		Deadline:      todoGRPCStruct.GetDeadline().AsTime(),
		Assignee:      todoGRPCStruct.GetAssignee(),
		TimeAllocated: todoGRPCStruct.GetTimeAllocated().AsDuration(),
		Status:        TodoStatus(todoGRPCStruct.GetStatus()),
	}

	return &todo
}
