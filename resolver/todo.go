package resolver

import (
	"context"

	"github.com/ninomae42/dataloader-v7/domain/todo"
	"github.com/ninomae42/dataloader-v7/loader"
)

type Todo struct {
	model todo.Todo
}

func NewTodo(model todo.Todo) Todo {
	return Todo{model: model}
}

func (t Todo) ID() string {
	return t.model.ID.String()
}

func (t Todo) Name() string {
	return t.model.Name.String()
}

func (t Todo) Assignees(ctx context.Context) ([]User, error) {
	loader, err := loader.FromContext(ctx)
	if err != nil {
		return []User{}, err
	}

	thunkMany := loader.UserLoader.LoadMany(ctx, t.model.Assignees)

	results, errs := thunkMany()
	for _, err := range errs {
		if err != nil {
			return []User{}, err
		}
	}

	users := make([]User, len(results))
	for i, u := range results {
		users[i] = NewUser(u)
	}

	return users, nil
}
