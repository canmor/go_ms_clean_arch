package blog

import "time"

type Blog struct {
	Id        int64
	Title     string
	Body      string
	CreatedAt time.Time
}
