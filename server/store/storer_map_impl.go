package store

import "go.uber.org/zap"

type MapStore struct {
	todosStore map[uint64]Todo

	logger *zap.Logger
}

