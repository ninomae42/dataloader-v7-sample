package controller

import (
	"context"

	"github.com/graph-gophers/graphql-go"
	"github.com/ninomae42/dataloader-v7/domain"
	"github.com/ninomae42/dataloader-v7/domain/todo"
	"github.com/ninomae42/dataloader-v7/resolver"
)

type Controller struct{}

type GetTodoInput struct {
	Input struct {
		ID graphql.ID
	}
}

func (c Controller) Todo(ctx context.Context, args GetTodoInput) (resolver.Todo, error) {
	m, err := todo.New(domain.String(args.Input.ID))
	if err != nil {
		return resolver.Todo{}, err
	}

	m, err = m.Update(
		todo.WithAssignees([]domain.ID{domain.ID("user_1"), domain.ID("user_2")}),
		todo.WithIsCompleted(false),
	)
	if err != nil {
		return resolver.Todo{}, err
	}

	return resolver.NewTodo(m), nil
}
