package todo

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/alvinmatias69/golang-onboarding/internal/model/todo"
)

func TestRepository_Get(t *testing.T) {
	now := time.Now()
	type fields struct {
		data []todo.Todo
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []todo.Todo
		wantErr bool
	}{
		{
			name: "Success get todo list",
			fields: fields{
				data: []todo.Todo{
					{
						ID:        1,
						Task:      "first",
						IsDone:    true,
						CreatedAt: now,
						UpdatedAt: now,
					},
				},
			},
			args: args{
				ctx: context.Background(),
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
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{
				data: tt.fields.data,
			}
			got, err := r.Get(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Repository.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_Add(t *testing.T) {
	type fields struct {
		data []todo.Todo
	}
	type args struct {
		ctx  context.Context
		item todo.Todo
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
			r := &Repository{
				data: tt.fields.data,
			}
			if err := r.Add(tt.args.ctx, tt.args.item); (err != nil) != tt.wantErr {
				t.Errorf("Repository.Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRepository_Update(t *testing.T) {
	type fields struct {
		data []todo.Todo
	}
	type args struct {
		ctx  context.Context
		id   int64
		item todo.Todo
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
			r := &Repository{
				data: tt.fields.data,
			}
			if err := r.Update(tt.args.ctx, tt.args.id, tt.args.item); (err != nil) != tt.wantErr {
				t.Errorf("Repository.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
