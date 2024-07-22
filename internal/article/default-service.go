package article

import (
	"errors"
	"go.uber.org/zap"
	"graphblog/graph/model"
	"strconv"
)

type DefaultService struct {
	log         *zap.Logger
	articleData DAO
}

func NewDefaultService(articleDAO DAO, logger *zap.Logger) Service {
	return &DefaultService{
		log:         logger,
		articleData: articleDAO,
	}
}

func (svc *DefaultService) Create(input model.Article) (model.Article, error) {

	if input.ID != `` {
		return model.Article{}, errors.New(`bad request`)
	}

	id, err := svc.articleData.Create(input)
	if err != nil {
		svc.log.Error(err.Error())
		return model.Article{}, err
	}
	input.ID = strconv.FormatInt(id, 10)

	return input, nil
}

func (svc *DefaultService) GetAll() ([]*model.Article, error) {

	records, err := svc.articleData.GetAll()
	if err != nil {
		svc.log.Error(err.Error())
		return nil, err
	}

	return records, nil
}
