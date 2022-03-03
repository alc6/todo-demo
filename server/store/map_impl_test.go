package store_test

import (
	"testing"
	"time"

	"github.com/google/uuid"

	"github.com/alc6/todo-demo/server/store"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"
)

func TestMapStore_CreateReadTodo(t *testing.T) {
	logger := zaptest.NewLogger(t)

	todoStore := store.NewMapStore(logger)

	require.NotNil(t, todoStore)

	task := store.Todo{
		Title:         "Very first todo",
		Description:   "Very first description",
		Deadline:      time.Now().Add(7 * 24 * time.Hour),
		Assignee:      "Me",
		TimeAllocated: time.Duration(7 * time.Second),
		Status:        store.TODO_STATUS_PENDING,
	}

	id, errCreate := todoStore.CreateTodo(&task)

	require.NoError(t, errCreate)
	assert.Len(t, id, 36)

	readTodo, errRead := todoStore.ReadTodos([]string{id}, false)

	require.NoError(t, errRead)
	require.Len(t, readTodo, 1)
	require.NotNil(t, readTodo[0])

	assert.True(t, assert.ObjectsAreEqual(*readTodo[0], task))

	readTodoAll, errReadAll := todoStore.ReadTodos([]string{}, false)

	require.NoError(t, errReadAll)
	require.Len(t, readTodoAll, 1)
	require.NotNil(t, readTodoAll[0])

	assert.True(t, assert.ObjectsAreEqual(*readTodoAll[0], task))
}

func TestMapStore_ReadUnregTodo(t *testing.T) {
	logger := zaptest.NewLogger(t)

	todoStore := store.NewMapStore(logger)

	require.NotNil(t, todoStore)

	unregUUID := uuid.NewString()

	todos, errRead := todoStore.ReadTodos([]string{unregUUID}, true)

	if assert.Error(t, errRead) {
		assert.Equal(t, store.ErrIDNotExist, errRead)
	}

	assert.Nil(t, todos)
}

func TestMapStore_UpdateTodosStatus(t *testing.T) {
	logger := zaptest.NewLogger(t)

	todoStore := store.NewMapStore(logger)

	require.NotNil(t, todoStore)

	task := store.Todo{
		Title:         "Very first todo",
		Description:   "Very first description",
		Deadline:      time.Now().Add(7 * 24 * time.Hour),
		Assignee:      "Me",
		TimeAllocated: time.Duration(7 * time.Second),
		Status:        store.TODO_STATUS_PENDING,
	}

	id, errCreate := todoStore.CreateTodo(&task)

	require.NoError(t, errCreate)
	assert.Len(t, id, 36)

	errUpdateStatus := todoStore.UpdateTodosStatus(id, store.TODO_STATUS_DONE)

	require.NoError(t, errUpdateStatus)

	readTodo, errRead := todoStore.ReadTodos([]string{id}, false)

	require.NoError(t, errRead)
	require.Len(t, readTodo, 1)
	require.NotNil(t, readTodo[0])

	assert.False(t, assert.ObjectsAreEqual(*readTodo[0], task))
	assert.Equal(t, store.TODO_STATUS_DONE, readTodo[0].Status)

	// Now try to modify a task that does not exist.
	readTodoUnreg, errReadUnreg := todoStore.ReadTodos([]string{uuid.NewString()}, true)

	if assert.Error(t, errReadUnreg) {
		assert.Equal(t, store.ErrIDNotExist, errReadUnreg)
	}

	assert.Nil(t, readTodoUnreg)
}

func TestMapStore_CreateDeleteTodo(t *testing.T) {
	logger := zaptest.NewLogger(t)

	todoStore := store.NewMapStore(logger)

	require.NotNil(t, todoStore)

	task := store.Todo{
		Title:         "Very first todo",
		Description:   "Very first description",
		Deadline:      time.Now().Add(7 * 24 * time.Hour),
		Assignee:      "Me",
		TimeAllocated: time.Duration(7 * time.Second),
		Status:        store.TODO_STATUS_PENDING,
	}

	id1, errCreate1 := todoStore.CreateTodo(&task)

	require.NoError(t, errCreate1)
	assert.Len(t, id1, 36)

	task2 := task

	id2, errCreate2 := todoStore.CreateTodo(&task2)

	require.NoError(t, errCreate2)
	assert.Len(t, id2, 36)

	// Make sure they are stored by reading it.
	readTodo, errRead := todoStore.ReadTodos([]string{id1, id2}, false)

	require.NoError(t, errRead)
	require.Len(t, readTodo, 2)
	require.NotNil(t, readTodo[0])
	require.NotNil(t, readTodo[1])

	assert.True(t, assert.ObjectsAreEqual(*readTodo[0], task))
	assert.True(t, assert.ObjectsAreEqual(*readTodo[1], task2))

	assert.NoError(t, todoStore.DeleteTodos([]string{id2}, false))

	remainingTasks, errReadRemaining := todoStore.ReadTodos([]string{}, false)

	require.NoError(t, errReadRemaining)
	require.Len(t, remainingTasks, 1)
	require.NotNil(t, remainingTasks[0])

	assert.True(t, assert.ObjectsAreEqual(*remainingTasks[0], task))

	// Delete last entry using empty slice to delete everything.
	assert.NoError(t, todoStore.DeleteTodos([]string{}, false))

	remainingTasks, errReadRemaining = todoStore.ReadTodos([]string{}, false)
	require.NoError(t, errReadRemaining)
	require.Len(t, remainingTasks, 0)
}
