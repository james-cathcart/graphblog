package article

import (
	"errors"
	"github.com/james-cathcart/golog"
	"graphblog/graph/model"
	"strconv"
)

type DefaultService struct {
	log         golog.GoLogger
	articleData DAO
}

func NewDefaultService(articleDAO DAO) Service {
	return &DefaultService{
		log:         golog.NewLogger(golog.NewNativeLogger(`[ article svc ] '`)),
		articleData: articleDAO,
	}
}

func (svc *DefaultService) Create(input model.Article) (model.Article, error) {

	if input.ID != `` {
		return model.Article{}, errors.New(`bad request`)
	}

	id, err := svc.articleData.Create(input)
	if err != nil {
		svc.log.Error(err)
		return model.Article{}, err
	}
	input.ID = strconv.FormatInt(id, 10)

	return input, nil
}

func (svc *DefaultService) GetAll() ([]*model.Article, error) {

	records, err := svc.articleData.GetAll()
	if err != nil {
		svc.log.Error(err)
		return nil, err
	}

	return records, nil
}
