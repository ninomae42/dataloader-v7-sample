package repository

import (
	"context"

	"github.com/ninomae42/dataloader-v7/domain"
	"github.com/ninomae42/dataloader-v7/domain/user"
	"gorm.io/gorm"
)

// User ユーザーリポジトリの実装
type User struct{}

// GetByIds IDでユーザーを取得する
func (u User) GetByIds(ctx context.Context, db *gorm.DB, IDs []domain.ID) (user.List, error) {
	// NOTE: This is a stub implementation.
	return user.List{}, nil
}
