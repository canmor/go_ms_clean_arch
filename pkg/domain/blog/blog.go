package blog

import (
	"fmt"
	"time"
)

type Blog struct {
	Id        int64
	Title     string
	Body      string
	CreatedAt time.Time
}

func (b Blog) URL() string {
	return fmt.Sprintf("https://www.blog.com/blogs/%d", b.Id)
}
