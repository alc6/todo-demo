package store

import (
	"errors"
	"sync"

	"github.com/google/uuid"

	"go.uber.org/zap"
)

var (
	// ErrIDAlreadyExist = errors.New("id already exists")
	ErrIDNotExist    = errors.New("specified id not found in store")
	ErrUnimplemented = errors.New("not implemented")
)

type MapStore struct {
	todosStore map[string]Todo

	logger *zap.Logger

	sync.RWMutex
}

func NewMapStore(logger *zap.Logger) *MapStore {
	mapStore := MapStore{
		todosStore: make(map[string]Todo),
		logger:     logger,
	}

	return &mapStore
}

func (ms *MapStore) CreateTodo(todo *Todo) (string, error) {
	ms.Lock()
	defer ms.Unlock()

	generatedID := uuid.NewString()

	todo.ID = generatedID

	ms.todosStore[generatedID] = *todo

	return generatedID, nil
}

func (ms *MapStore) ReadTodos(ids []string, stopIfMiss bool) ([]*Todo, error) {
	ms.RLock()
	defer ms.RUnlock()

	readTodos := make([]*Todo, 0)

	if len(ids) == 0 {
		// Then read all
		for _, todo := range ms.todosStore {
			// the todo var keeps the same addr thus need for a copy of the object.
			todoCpy := todo

			readTodos = append(readTodos, &todoCpy)
		}

		return readTodos, nil
	}

	for _, id := range ids {
		todo, exists := ms.todosStore[id]

		if !exists {
			if stopIfMiss {
				return nil, ErrIDNotExist
			}

			continue
		}

		readTodos = append(readTodos, &todo)
	}

	return readTodos, nil
}

func (ms *MapStore) UpdateTodosStatus(id string, newStatus TodoStatus) error {
	ms.Lock()
	defer ms.Unlock()

	todo, exists := ms.todosStore[id]
	if !exists {
		return ErrIDNotExist
	}

	todo.Status = newStatus

	ms.todosStore[id] = todo

	return nil
}

func (ms *MapStore) UpdateTodos(_ []*Todo) error {
	return ErrUnimplemented
}

func (ms *MapStore) DeleteTodos(ids []string, stopIfMiss bool) error {
	ms.Lock()
	defer ms.Unlock()

	if len(ids) == 0 {
		for id := range ms.todosStore {
			delete(ms.todosStore, id)
		}

		return nil
	}

	for _, id := range ids {
		_, exists := ms.todosStore[id]

		if !exists {
			if stopIfMiss {
				return ErrIDNotExist
			}

			continue
		}

		delete(ms.todosStore, id)
	}

	return nil
}
