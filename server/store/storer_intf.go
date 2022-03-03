package store

type Storer interface {
	CreateTodo(Todo) (uint64, error)
	ReadTodos([]uint64) ([]Todo, error)
	UpdateTodos([]Todo) (error)
	DeleteTodos([]uint64) (bool, error)
}
