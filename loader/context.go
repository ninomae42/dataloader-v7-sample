package loader

import (
	"context"
	"errors"
)

type loaderKey struct{}

var loaderContextKey = loaderKey{}

var ErrLoaderNotFound = errors.New("loader: loader not found in context")

// WithContext loadlerをcontextにセットする
func WithContext(ctx context.Context, loader Loader) context.Context {
	return context.WithValue(ctx, loaderContextKey, loader)
}

// FromContext contextからloaderを取得する
func FromContext(ctx context.Context) (Loader, error) {
	loader, ok := ctx.Value(loaderContextKey).(Loader)
	if !ok {
		return Loader{}, ErrLoaderNotFound
	}
	return loader, nil
}
