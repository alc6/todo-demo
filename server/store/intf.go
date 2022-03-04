package store

// Storer interface provides function to satisfy to implement a storage for Todo.
type Storer interface {
	CreateTodo(*Todo) (string, error)
	ReadTodos([]string, bool) ([]*Todo, error)
	UpdateTodos([]*Todo) error
	UpdateTodosStatus(string, TodoStatus) error
	DeleteTodos([]string, bool) error
}
