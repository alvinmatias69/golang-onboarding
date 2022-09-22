package todo

import (
	"context"

	"github.com/alvinmatias69/golang-onboarding/internal/model/todo"
)

// List saved todo list with given filter
func (u *UseCase) List(ctx context.Context, filter todo.TodoFilter) ([]todo.Todo, error) {
	todoList, err := u.repo.Get(ctx)
	if err != nil {
		return nil, err
	}

	// TODO: implement filter here

	return todoList, nil
}

// Add given parameter into current todo list
// TODO: implement this. Remember to add the timestamp
func (u *UseCase) Add(ctx context.Context, payload todo.TodoCreatePayload) error {
	return nil
}

// MarkAsDone for given taskID
// TODO: implement this. Remember to add the updated timestamp
func (u *UseCase) MarkAsDone(ctx context.Context, taskID uint64) error {
	return nil
}
