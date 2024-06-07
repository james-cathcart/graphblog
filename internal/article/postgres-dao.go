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

	stmt, err := dao.db.Prepare(`SELECT id, title, content, status FROM articles ORDER BY id DESC`)
	if err != nil {
		dao.log.Error(err)
		return nil, err
	}

	rows, err := stmt.Query()
	if err != nil {
		dao.log.Error(err)
		return nil, err
	}

	var records []*model.Article
	for rows.Next() {
		tmp := model.Article{}
		err = rows.Scan(&tmp.ID, &tmp.Content)
		if err != nil {
			continue
		}

		records = append(records, &tmp)

	}

	return records, nil
}
