package article

import (
	"database/sql"
	"errors"
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

	stmt, err := dao.db.Prepare(`INSERT INTO articles (title, content, status, actor_id) VALUES ($1,$2,$3,$4) RETURNING id`)
	if err != nil {
		dao.log.Error(err)
		return -1, err
	}
	defer func(closeFunc func() error) {
		err = closeFunc()
		if err != nil {
			dao.log.Error(err)
		}
	}(stmt.Close)

	row := stmt.QueryRow(input.Title, input.Content, input.Status, input.User.ID)

	var id int64
	err = row.Scan(&id)
	if err != nil {
		dao.log.Error(err)
		return -1, err
	}

	if id == 0 {
		err = errors.New(`failed to create record, possible duplicate`)
		dao.log.Error(err)
		return -1, err
	}

	return id, nil
}

func (dao *PostgresDAO) GetAll() ([]*model.Article, error) {

	stmt, err := dao.db.Prepare(`SELECT articles.id, articles.title, articles.content, articles.status, actors.id, actors.display_name FROM articles INNER JOIN actors ON articles.actor_id=actors.id ORDER BY articles.id DESC`)
	if err != nil {
		dao.log.Error(err)
		return nil, err
	}
	defer func(closeFunc func() error) {
		err = closeFunc()
		if err != nil {
			dao.log.Error(err)
		}
	}(stmt.Close)

	rows, err := stmt.Query()
	if err != nil {
		dao.log.Error(err)
		return nil, err
	}

	var records []*model.Article
	for rows.Next() {
		tmp := model.Article{
			User: &model.User{},
		}
		err = rows.Scan(&tmp.ID, &tmp.Title, &tmp.Content, &tmp.Status, &tmp.User.ID, &tmp.User.Name)
		if err != nil {
			dao.log.Error(err)
			continue
		}

		records = append(records, &tmp)

	}

	return records, nil
}
