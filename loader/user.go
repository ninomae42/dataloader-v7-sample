package loader

import (
	"context"

	"github.com/graph-gophers/dataloader/v7"
	"github.com/ninomae42/dataloader-v7/domain"
	"github.com/ninomae42/dataloader-v7/domain/user"
	"github.com/ninomae42/dataloader-v7/registry"
	"gorm.io/gorm"
)

type userLoader struct {
	db   *gorm.DB
	repo user.Repository
}

func NewUserLoader(r registry.Registry) userLoader {
	return userLoader{
		db:   r.Repository.DB(),
		repo: r.Repository.NewUserRepostiory(),
	}
}

func (u userLoader) BatchedLoader(ctx context.Context, ids []domain.ID) []*dataloader.Result[user.User] {
	result := make([]*dataloader.Result[user.User], len(ids))

	users, err := u.repo.GetByIds(ctx, u.db, ids)
	if err != nil {
		for i := range ids {
			result[i] = &dataloader.Result[user.User]{Error: err}
		}
		return result
	}

	userByID := make(map[domain.ID]user.User, len(users))
	for _, user := range users {
		userByID[user.ID] = user
	}

	for i, id := range ids {
		if u, ok := userByID[id]; ok {
			result[i] = &dataloader.Result[user.User]{Data: u}
		} else {
			result[i] = &dataloader.Result[user.User]{Error: user.ErrNotFound}
		}
	}
	return result
}
