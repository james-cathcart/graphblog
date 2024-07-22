package user

import (
	"errors"
	"go.uber.org/zap"
	"graphblog/graph/model"
	"strconv"
)

type DefaultService struct {
	log      *zap.Logger
	userData DAO
}

func NewDefaultService(userDAO DAO, logger *zap.Logger) Service {
	return &DefaultService{
		log:      logger,
		userData: userDAO,
	}
}

func (svc *DefaultService) Create(input model.User) (model.User, error) {

	if input.ID != `` {
		return model.User{}, errors.New(`bad request`)
	}
	if input.Name == `` {
		return model.User{}, errors.New(`bad request`)
	}

	id, err := svc.userData.Create(input)
	if err != nil {
		svc.log.Error(err.Error())
		return model.User{}, err
	}

	input.ID = strconv.FormatInt(id, 10)

	return input, nil
}

func (svc *DefaultService) GetAll() ([]*model.User, error) {

	records, err := svc.userData.GetAll()
	if err != nil {
		svc.log.Error(err.Error())
		return nil, err
	}

	return records, nil
}
