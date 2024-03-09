package middleware

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/ninomae42/dataloader-v7/loader"
	"github.com/ninomae42/dataloader-v7/registry"
)

type middleware struct {
	registry registry.Registry
}

func newMiddleware(r registry.Registry) *middleware {
	return &middleware{
		registry: r,
	}
}

func (m middleware) withDataLoader(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		l := loader.New(m.registry)

		ctx = loader.WithContext(ctx, l)

		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}

func Mount(ctx context.Context, r *chi.Mux) (*chi.Mux, error) {
	registry, err := registry.New(ctx)
	if err != nil {
		return nil, err
	}
	m := newMiddleware(registry)
	r.Use(m.withDataLoader)
	return r, nil
}
