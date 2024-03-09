package resolver

import "github.com/ninomae42/dataloader-v7/domain/user"

type User struct {
	model user.User
}

func NewUser(model user.User) User {
	return User{model: model}
}

func (u User) ID() string {
	return u.model.ID.String()
}

func (u User) Name() string {
	return u.model.Name.String()
}
