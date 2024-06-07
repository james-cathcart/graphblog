package user

import (
	"github.com/james-cathcart/golog"
	"graphblog/graph/model"
)

type DefaultService struct {
	log      golog.GoLogger
	userData DAO
}

func NewDefaultService(userDAO DAO) Service {
	return &DefaultService{
		log:      golog.NewLogger(golog.NewNativeLogger(`[ user svc ] `)),
		userData: userDAO,
	}
}

func (svc *DefaultService) Create(input model.User) (model.User, error) {
	return model.User{}, nil
}

func (svc *DefaultService) GetAll() ([]*model.User, error) {
	return nil, nil
}
