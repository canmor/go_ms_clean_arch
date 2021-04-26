package usecase

import (
	"github.com/canmor/go_ms_clean_arch/pkg/domain/blog"
	"time"
)

type BlogUseCase struct {
	repo            blog.BlogRepository
	shortURLService ShortURL
}

func NewBlogUseCase(repo blog.BlogRepository, shortURLService ShortURL) *BlogUseCase {
	return &BlogUseCase{repo, shortURLService}
}

func (b BlogUseCase) Create(title string, body string) *blog.Blog {
	res := blog.Blog{Title: title, Body: body, CreatedAt: time.Now()}
	id, err := b.repo.Save(res)
	if err != nil {
		return nil
	}
	res.Id = id
	return &res
}

func (b BlogUseCase) Share(blog blog.Blog) (string, error) {
	return b.shortURLService.Create(blog.URL())
}
