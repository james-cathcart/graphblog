package article

import "graphblog/graph/model"

// DAO abstracts the data access layer
type DAO interface {
	Create(input model.Article) (int64, error)
	GetAll() ([]*model.Article, error)
}

// Service abstracts the business logic
type Service interface {
	Create(input model.Article) (model.Article, error)
	GetAll() ([]*model.Article, error)
}
