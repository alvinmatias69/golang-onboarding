package todo

import (
	"errors"
	"testing"
	"time"

	"github.com/alvinmatias69/golang-onboarding/internal/model/todo"
	"github.com/golang/mock/gomock"
)

func TestHandler_Parse(t *testing.T) {
	type fields struct {
		usecase todoUseCase
	}
	type args struct {
		params string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Failure on empty params",
			args: args{
				params: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				usecase: tt.fields.usecase,
			}
			if err := h.Parse(tt.args.params); (err != nil) != tt.wantErr {
				t.Errorf("Handler.Parse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHandler_list(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUseCase := NewMocktodoUseCase(ctrl)

	type args struct {
		params string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		mockFn  func(args)
	}{
		{
			name:    "Failure on error usecase",
			args:    args{},
			wantErr: true,
			mockFn: func(a args) {
				mockUseCase.
					EXPECT().
					List(gomock.Any(), todo.TodoFilter{}).
					Return(nil, errors.New("expected error"))
			},
		},
		{
			name:    "Success return simple filter",
			args:    args{},
			wantErr: false,
			mockFn: func(a args) {
				mockUseCase.
					EXPECT().
					List(gomock.Any(), todo.TodoFilter{}).
					Return([]todo.Todo{
						{
							ID:        1,
							Task:      "First",
							IsDone:    false,
							CreatedAt: time.Now(),
							UpdatedAt: time.Now(),
						},
					}, nil)
			},
		},
		{
			name: "Success return ID filter",
			args: args{
				params: "1",
			},
			wantErr: false,
			mockFn: func(a args) {
				mockUseCase.
					EXPECT().
					List(gomock.Any(), todo.TodoFilter{
						ID:     1,
						IsDone: nil,
					}).
					Return([]todo.Todo{
						{
							ID:        1,
							Task:      "First",
							IsDone:    false,
							CreatedAt: time.Now(),
							UpdatedAt: time.Now(),
						},
					}, nil)
			},
		},
		{
			name: "Success return IsDone filter",
			args: args{
				params: "false",
			},
			wantErr: false,
			mockFn: func(a args) {
				mockUseCase.
					EXPECT().
					List(gomock.Any(), gomock.Any()).
					Return([]todo.Todo{
						{
							ID:        1,
							Task:      "First",
							IsDone:    true,
							CreatedAt: time.Now(),
							UpdatedAt: time.Now(),
						},
					}, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				usecase: mockUseCase,
			}
			tt.mockFn(tt.args)
			if err := h.list(tt.args.params); (err != nil) != tt.wantErr {
				t.Errorf("Handler.list() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
