package graph

import (
	"graphblog/internal/article"
	"graphblog/internal/user"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	articleSvc article.Service
	userSvc    user.Service
}

func NewResolver(articleSvc article.Service, userSvc user.Service) *Resolver {

	return &Resolver{
		articleSvc: articleSvc,
		userSvc:    userSvc,
	}

}
