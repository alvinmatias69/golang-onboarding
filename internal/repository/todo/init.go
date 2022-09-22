package todo

import (
	"time"

	"github.com/alvinmatias69/golang-onboarding/internal/model/todo"
)

// Repository define main struct for todo repo
type Repository struct {
	data []todo.Todo
}

// New todo repository instance
func New() *Repository {
	return &Repository{
		data: []todo.Todo{
			{
				ID:        1,
				Task:      "First",
				IsDone:    true,
				CreatedAt: time.Now(),
			},
		},
	}
}
