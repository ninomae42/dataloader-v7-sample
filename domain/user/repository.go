package user

import (
	"context"

	"github.com/ninomae42/dataloader-v7/domain"
	"gorm.io/gorm"
)

// Repository ユーザーリポジトリのインターフェース
type Repository interface {
	// GetByIds IDでユーザーを取得する
	GetByIds(ctx context.Context, db *gorm.DB, IDs []domain.ID) (List, error)
}
