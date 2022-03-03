package store

type Storer interface {
	CreateTodo(*Todo) (string, error)
	ReadTodos([]string, bool) ([]*Todo, error)
	UpdateTodos([]*Todo) error
	UpdateTodosStatus(string, TodoStatus) error
	DeleteTodos([]string, bool) error
}
