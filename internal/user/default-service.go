package user

import (
	"errors"
	"github.com/james-cathcart/golog"
	"graphblog/graph/model"
	"strconv"
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

	if input.ID != `` {
		return model.User{}, errors.New(`bad request`)
	}
	if input.Name == `` {
		return model.User{}, errors.New(`bad request`)
	}

	id, err := svc.userData.Create(input)
	if err != nil {
		svc.log.Error(err)
		return model.User{}, err
	}

	input.ID = strconv.FormatInt(id, 10)

	return input, nil
}

func (svc *DefaultService) GetAll() ([]*model.User, error) {

	records, err := svc.userData.GetAll()
	if err != nil {
		svc.log.Error(err)
		return nil, err
	}

	return records, nil
}
