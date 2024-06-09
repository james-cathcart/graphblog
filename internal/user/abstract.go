package user

//go:generate mockgen -destination=dao_mock.go -package=user . DAO
//go:generate mockgen -destination=service_mock.go -package=user . Service

import (
	"graphblog/graph/model"
)

// DAO abstracts the data access layer
type DAO interface {
	Create(input model.User) (int64, error)
	GetAll() ([]*model.User, error)
}

// Service abstracts the business logic layer
type Service interface {
	Create(input model.User) (model.User, error)
	GetAll() ([]*model.User, error)
}
