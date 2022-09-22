package todo

import "time"

// Todo is struct for todo item
type Todo struct {
	ID        uint64
	Task      string
	IsDone    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

// TodoCreatePayload define payload for creating an item
type TodoCreatePayload struct {
	Task string
}

// TodoFilter define filter for fetching todo list
type TodoFilter struct {
	ID     uint64
	IsDone *bool
}
