package todo

import (
	"context"

	"github.com/alvinmatias69/golang-onboarding/internal/model/todo"
)

// Get todo list
func (r *Repository) Get(ctx context.Context) ([]todo.Todo, error) {
	return r.data, nil
}

// Add new item to todo list
// TODO: implement this
func (r *Repository) Add(ctx context.Context, item todo.Todo) error {
	return nil
}

// Update item in todo list
// TODO: implement this
func (r *Repository) Update(ctx context.Context, id uint64, item todo.Todo) error {
	return nil
}
