package loader

import (
	"github.com/graph-gophers/dataloader/v7"
	"github.com/ninomae42/dataloader-v7/domain"
	"github.com/ninomae42/dataloader-v7/domain/user"
	"github.com/ninomae42/dataloader-v7/registry"
)

// Loader DataLoaderの集合
type Loader struct {
	UserLoader *dataloader.Loader[domain.ID, user.User]
}

// New Loaderのコンストラクタ
func New(r registry.Registry) Loader {
	userLoader := NewUserLoader(r)

	return Loader{
		UserLoader: dataloader.NewBatchedLoader(userLoader.BatchedLoader),
	}
}
