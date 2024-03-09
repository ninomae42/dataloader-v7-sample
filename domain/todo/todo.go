package todo

import (
	"github.com/ninomae42/dataloader-v7/domain"
)

type (
	Todo struct {
		ID          domain.ID
		Name        domain.String
		Assignees   []domain.ID
		IsCompleted bool
	}

	List []Todo

	TodoOption func(*Todo) error
)

func New(name domain.String) (Todo, error) {
	t := Todo{
		ID:   domain.NewID(),
		Name: name,
	}
	if err := t.validate(); err != nil {
		return Todo{}, err
	}
	return t, nil
}

func (t Todo) Update(opts ...TodoOption) (Todo, error) {
	for _, opt := range opts {
		if err := opt(&t); err != nil {
			return Todo{}, err
		}
	}
	if err := t.validate(); err != nil {
		return Todo{}, err
	}
	return t, nil
}

func WithName(name domain.String) TodoOption {
	return func(t *Todo) error {
		// TODO: Add validation
		t.Name = name
		return nil
	}
}

func WithAssignees(assignees []domain.ID) TodoOption {
	return func(t *Todo) error {
		t.Assignees = assignees
		return nil
	}
}

func WithIsCompleted(isCompleted bool) TodoOption {
	return func(t *Todo) error {
		t.IsCompleted = isCompleted
		return nil
	}
}

func (t Todo) validate() error {
	return nil
}
