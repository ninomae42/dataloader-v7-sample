package user

import (
	"errors"

	"github.com/ninomae42/dataloader-v7/domain"
)

var ErrNotFound = errors.New("user: user not found")

type (
	User struct {
		ID   domain.ID
		Name domain.String
	}

	List []User

	UserOption func(*User) error
)

func New(name domain.String) (User, error) {
	u := User{
		ID:   domain.NewID(),
		Name: name,
	}
	if err := u.validate(); err != nil {
		return User{}, err
	}
	return u, nil
}

func (u User) Update(opts ...func(*User) error) (User, error) {
	for _, opt := range opts {
		if err := opt(&u); err != nil {
			return User{}, err
		}
	}

	if err := u.validate(); err != nil {
		return User{}, err
	}

	return u, nil
}

func WithName(name domain.String) UserOption {
	return func(u *User) error {
		if err := name.ValidateLength(0, 20); err != nil {
			return err
		}
		u.Name = name
		return nil
	}
}

func (u User) validate() error {
	return nil
}
