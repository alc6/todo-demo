package store

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	"google.golang.org/protobuf/types/known/durationpb"

	"github.com/alc6/todo-demo/proto/todorpc"
)

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

const (
	TODO_STATUS_PENDING  TodoStatus = 0
	TODO_STATUS_DOING    TodoStatus = 1
	TODO_STATUS_DONE     TodoStatus = 2
	TODO_STATUS_CANCELED TodoStatus = 3
	TODO_STATUS_EXPIRED  TodoStatus = 4
)

func (t Todo) TodoToGRPCStruct() *todorpc.TodoWithMeta {
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

func TodoFromGRPCStruct(todoRpcStruct *todorpc.TodoWithMeta) *Todo {
	todo := Todo{
		ID:            todoRpcStruct.GetId(),
		Title:         todoRpcStruct.GetTitle(),
		Description:   todoRpcStruct.GetDescription(),
		Deadline:      todoRpcStruct.GetDeadline().AsTime(),
		Assignee:      todoRpcStruct.GetAssignee(),
		TimeAllocated: todoRpcStruct.GetTimeAllocated().AsDuration(),
		Status:        TodoStatus(todoRpcStruct.GetStatus()),
	}

	return &todo
}
