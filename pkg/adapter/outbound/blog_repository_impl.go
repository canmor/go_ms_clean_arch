package outbound

import (
	"database/sql"
	"fmt"
	"github.com/canmor/go_ms_clean_arch/pkg/domain/blog"
	"github.com/canmor/go_ms_clean_arch/pkg/util"
)

type BlogRepositoryImpl struct {
	db *sql.DB
}

func NewBlogRepository(db *sql.DB) blog.BlogRepository {
	return BlogRepositoryImpl{db}
}

func (b BlogRepositoryImpl) Save(blog blog.Blog) (int64, error) {
	res, err := b.db.Exec("insert into blogs(title, body, created_at) values(?, ?, ?)", blog.Title, blog.Body, blog.CreatedAt)
	if err != nil {
		err = fmt.Errorf("BlogRepositoryImpl.Save failed, err: %s", err)
		util.Log().Error(err)
		return -1, err
	}
	id, err := res.LastInsertId()
	return id, err
}

func (b BlogRepositoryImpl) Find(id int) (*blog.Blog, error) {
	// FIXME: add created_at
	r := b.db.QueryRow("select id, title, body FROM blogs WHERE id=?", id)
	found := blog.Blog{}
	err := r.Scan(&found.Id, &found.Title, &found.Body)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &found, nil
}
