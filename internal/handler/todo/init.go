package todo

// Handler define main struct for todo handler
type Handler struct {
	usecase todoUseCase
}

// New todo handler instance
func New(usecase todoUseCase) *Handler {
	return &Handler{
		usecase: usecase,
	}
}
