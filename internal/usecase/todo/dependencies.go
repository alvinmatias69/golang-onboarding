package todo

import (
	"context"

	"github.com/alvinmatias69/golang-onboarding/internal/model/todo"
)

//go:generate mockgen -build_flags=-mod=mod -source=dependencies.go -package=todo -destination=dependencies.mock.test.go

type todoRepository interface {
	Get(ctx context.Context) ([]todo.Todo, error)
	Add(ctx context.Context, item todo.Todo) error
	Update(ctx context.Context, id uint64, item todo.Todo) error
}
