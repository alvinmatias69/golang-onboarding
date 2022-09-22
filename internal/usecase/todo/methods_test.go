package todo

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/alvinmatias69/golang-onboarding/internal/model/todo"
	"github.com/golang/mock/gomock"
)

func TestUseCase_List(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMocktodoRepository(ctrl)
	now := time.Now()

	type args struct {
		ctx    context.Context
		filter todo.TodoFilter
	}
	tests := []struct {
		name    string
		args    args
		want    []todo.Todo
		wantErr bool
		mockFn  func(args)
	}{
		{
			name: "Failure on error get from repo",
			args: args{
				ctx:    context.Background(),
				filter: todo.TodoFilter{},
			},
			want:    nil,
			wantErr: true,
			mockFn: func(a args) {
				mockRepo.
					EXPECT().
					Get(a.ctx).
					Return(nil, errors.New("expected error"))
			},
		},
		{
			name: "Success get simple list from repo",
			args: args{
				ctx:    context.Background(),
				filter: todo.TodoFilter{},
			},
			want: []todo.Todo{
				{
					ID:        1,
					Task:      "first",
					IsDone:    true,
					CreatedAt: now,
					UpdatedAt: now,
				},
			},
			wantErr: false,
			mockFn: func(a args) {
				mockRepo.
					EXPECT().
					Get(a.ctx).
					Return([]todo.Todo{
						{
							ID:        1,
							Task:      "first",
							IsDone:    true,
							CreatedAt: now,
							UpdatedAt: now,
						},
					}, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UseCase{
				repo: mockRepo,
			}
			tt.mockFn(tt.args)
			got, err := u.List(tt.args.ctx, tt.args.filter)
			if (err != nil) != tt.wantErr {
				t.Errorf("UseCase.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UseCase.List() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCase_Add(t *testing.T) {
	type fields struct {
		repo todoRepository
	}
	type args struct {
		ctx     context.Context
		payload todo.TodoCreatePayload
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UseCase{
				repo: tt.fields.repo,
			}
			if err := u.Add(tt.args.ctx, tt.args.payload); (err != nil) != tt.wantErr {
				t.Errorf("UseCase.Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUseCase_MarkAsDone(t *testing.T) {
	type fields struct {
		repo todoRepository
	}
	type args struct {
		ctx    context.Context
		taskID int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UseCase{
				repo: tt.fields.repo,
			}
			if err := u.MarkAsDone(tt.args.ctx, tt.args.taskID); (err != nil) != tt.wantErr {
				t.Errorf("UseCase.MarkAsDone() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
