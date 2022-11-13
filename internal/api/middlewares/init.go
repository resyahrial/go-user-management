package middlewares

import (
	"time"

	"github.com/resyahrial/go-user-management/config"
	"github.com/resyahrial/go-user-management/internal/entities"
	"github.com/resyahrial/go-user-management/internal/factory"
	"gorm.io/gorm"
)

type Middleware struct {
	authUsecase entities.AuthUsecase
}

type MiddlewareOpts struct {
	Db  *gorm.DB
	Cfg config.Config
}

type MiddlewareOptionFn func(*Middleware, MiddlewareOpts)

func NewMiddleware(mOpts MiddlewareOpts, opts ...MiddlewareOptionFn) *Middleware {
	m := &Middleware{}
	for _, opt := range opts {
		opt(m, mOpts)
	}
	return m
}

func WithAuthUsecase(m *Middleware, opts MiddlewareOpts) {
	cfg := opts.Cfg
	m.authUsecase = factory.InitAuthUsecase(opts.Db, cfg.Hasher.Cost, time.Duration(cfg.Auth.AccessTimeDuration*int(time.Second)), cfg.Auth.AccessSecretKey)
}
