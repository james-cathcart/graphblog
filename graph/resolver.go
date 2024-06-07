package graph

import (
	"github.com/james-cathcart/golog"
	"graphblog/internal/article"
	"graphblog/internal/user"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	log        golog.GoLogger
	articleSvc article.Service
	userSvc    user.Service
}

func NewResolver(articleSvc article.Service, userSvc user.Service) *Resolver {

	return &Resolver{
		log:        golog.NewLogger(golog.NewNativeLogger(`[ resolver ] `)),
		articleSvc: articleSvc,
		userSvc:    userSvc,
	}

}
