package article

import (
	"github.com/james-cathcart/golog"
	"graphblog/graph/model"
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
	return model.Article{}, nil
}

func (svc *DefaultService) GetAll() ([]*model.Article, error) {
	return nil, nil
}
