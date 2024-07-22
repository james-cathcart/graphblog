package user

import (
	"database/sql"
	"errors"
	"go.uber.org/zap"
	"graphblog/graph/model"
)

type PostgresDAO struct {
	log *zap.Logger
	db  *sql.DB
}

func NewPostgresDAO(db *sql.DB, logger *zap.Logger) DAO {
	return &PostgresDAO{
		log: logger,
		db:  db,
	}
}

func (dao *PostgresDAO) Create(input model.User) (int64, error) {

	stmt, err := dao.db.Prepare(`INSERT INTO actors (display_name) VALUES ($1) RETURNING id`)
	if err != nil {
		dao.log.Error(err.Error())
		return -1, err
	}
	defer func(closeFunc func() error) {
		err = closeFunc()
		if err != nil {
			dao.log.Error(err.Error())
		}
	}(stmt.Close)

	row := stmt.QueryRow(input.Name)

	var id int64
	err = row.Scan(&id)

	if id == 0 {
		err = errors.New(`failed to create record, possible duplicate`)
		dao.log.Error(err.Error())
		return -1, err
	}

	return id, nil

}

func (dao *PostgresDAO) GetAll() ([]*model.User, error) {

	stmt, err := dao.db.Prepare(`SELECT id, display_name FROM actors ORDER BY id DESC`)
	if err != nil {
		dao.log.Error(err.Error())
		return nil, err
	}
	defer func(closeFunc func() error) {
		err = closeFunc()
		if err != nil {
			dao.log.Error(err.Error())
		}
	}(stmt.Close)

	rows, err := stmt.Query()
	if err != nil {
		dao.log.Error(err.Error())
		return nil, err
	}

	var records []*model.User
	for rows.Next() {
		tmp := model.User{}
		err = rows.Scan(&tmp.ID, &tmp.Name)
		if err != nil {
			dao.log.Error(err.Error())
			continue
		}
		records = append(records, &tmp)
	}

	return records, nil
}
