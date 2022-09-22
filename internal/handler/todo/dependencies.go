package todo

import (
	"context"

	"github.com/alvinmatias69/golang-onboarding/internal/model/todo"
)

//go:generate mockgen -build_flags=-mod=mod -source=dependencies.go -package=todo -destination=dependencies.mock.test.go

type todoUseCase interface {
	List(ctx context.Context, filter todo.TodoFilter) ([]todo.Todo, error)
	Add(ctx context.Context, payload todo.TodoCreatePayload) error
	MarkAsDone(ctx context.Context, taskID uint64) error
}
