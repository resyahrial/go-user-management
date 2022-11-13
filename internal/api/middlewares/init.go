package middlewares

import (
	"github.com/resyahrial/go-user-management/internal/entities"
)

type Middleware struct {
	authUsecase entities.AuthUsecase
}

type MiddlewareOptionFn func(*Middleware)

func NewMiddleware(opts ...MiddlewareOptionFn) *Middleware {
	m := &Middleware{}
	for _, opt := range opts {
		opt(m)
	}
	return m
}
