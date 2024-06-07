package article

import (
	"database/sql"
	"github.com/james-cathcart/golog"
	"graphblog/graph/model"
)

type PostgresDAO struct {
	log golog.GoLogger
	db  *sql.DB
}

func NewPostgresDAO(db *sql.DB) DAO {
	return &PostgresDAO{
		log: golog.NewLogger(golog.NewNativeLogger(`[ article dao ] `)),
		db:  db,
	}
}

func (dao *PostgresDAO) Create(input model.Article) (int64, error) {
	return -1, nil
}

func (dao *PostgresDAO) GetAll() ([]*model.Article, error) {
	return nil, nil
}
