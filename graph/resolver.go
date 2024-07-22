package graph

import (
	"go.uber.org/zap"
	"graphblog/internal/article"
	"graphblog/internal/user"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	log        *zap.Logger
	articleSvc article.Service
	userSvc    user.Service
}

func NewResolver(articleSvc article.Service, userSvc user.Service, logger *zap.Logger) *Resolver {

	return &Resolver{
		log:        logger,
		articleSvc: articleSvc,
		userSvc:    userSvc,
	}

}
