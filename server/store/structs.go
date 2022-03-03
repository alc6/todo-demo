package store

import (
	"time"
)

type Todo struct {
	Id            uint64        `json:"id,omitempty"`
	Title         string        `json:"title,omitempty"`
	Description   string        `json:"description,omitempty"`
	Deadline      time.Time     `json:"deadline,omitempty"`
	Assignee      string        `json:"assignee,omitempty"`
	TimeAllocated time.Duration `json:"time_allocated,omitempty"`
	Status        TodoStatus    `json:"status,omitempty"`
}

type TodoStatus int32

const (
	TodoStatus_TODO_STATUS_PENDING  TodoStatus = 0
	TodoStatus_TODO_STATUS_DOING    TodoStatus = 1
	TodoStatus_TODO_STATUS_DONE     TodoStatus = 2
	TodoStatus_TODO_STATUS_CANCELED TodoStatus = 3
	TodoStatus_TODO_STATUS_EXPIRED  TodoStatus = 4
)
