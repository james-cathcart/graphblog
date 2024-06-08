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

	stmt, err := dao.db.Prepare(`INSERT INTO articles (title, content, status, user_id) VALUES (?,?,?,?)`)
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

	result, err := stmt.Exec(input.Title, input.Content, input.Status, input.User.ID)
	if err != nil {
		dao.log.Error(err)
		return -1, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		dao.log.Error(err)
		return -1, err
	}

	return id, nil
}

func (dao *PostgresDAO) GetAll() ([]*model.Article, error) {

	stmt, err := dao.db.Prepare(`SELECT articles.id, articles.title, articles.content, articles.status, users.id, users.display_name FROM articles INNER JOIN users ON articles.user_id=users.id ORDER BY articles.id DESC`)
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
