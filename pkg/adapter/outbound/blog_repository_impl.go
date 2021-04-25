package outbound

import (
	"database/sql"
	"errors"
	"github.com/canmor/go_ms_clean_arch/pkg/domain/blog"
	"log"
)

type BlogRepositoryImpl struct {
	db *sql.DB
}

func NewBlogRepository(db *sql.DB) blog.BlogRepository {
	return BlogRepositoryImpl{db}
}

func (b BlogRepositoryImpl) Save(blog blog.Blog) (int64, error) {
	tx, err := b.db.Begin()
	if err != nil {
		log.Print(err)
	}
	stmt, err := tx.Prepare("insert into blogs(title, body, created_at) values(?, ?, ?)")
	if err != nil {
		log.Print(err)
	}
	defer func() {
		_ = stmt.Close()
	}()

	res, err := stmt.Exec(blog.Title, blog.Body, blog.CreatedAt)
	if err != nil {
		log.Print(err)
		return 0, err
	}
	err = tx.Commit()
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	return id, err
}

func (b BlogRepositoryImpl) Find(_ int) (*blog.Blog, error) {
	return nil, errors.New("unimplemented")
}
