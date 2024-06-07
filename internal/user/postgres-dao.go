package user

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
		log: golog.NewLogger(golog.NewNativeLogger(`[ user dao ] `)),
		db:  db,
	}
}

func (svc *PostgresDAO) Create(input model.User) (int64, error) {
	return -1, nil
}

func (svc *PostgresDAO) GetAll() ([]*model.User, error) {
	return nil, nil
}
