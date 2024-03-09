package registry

import (
	"context"

	"github.com/ninomae42/dataloader-v7/domain/user"
	"github.com/ninomae42/dataloader-v7/repository"
	"gorm.io/gorm"
)

type Registry struct {
	Repository Repository
}

func New(ctx context.Context) (Registry, error) {
	return Registry{}, nil
}

type Repository struct {
	db *gorm.DB
}

func (r Repository) DB() *gorm.DB {
	return r.db
}

func (r Repository) NewUserRepostiory() user.Repository {
	return repository.User{}
}
