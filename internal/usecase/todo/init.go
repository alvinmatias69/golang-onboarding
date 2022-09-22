package todo

// UseCase define main todo struct
type UseCase struct {
	repo todoRepository
}

// New usecase instance
func New(repo todoRepository) *UseCase {
	return &UseCase{
		repo: repo,
	}
}
