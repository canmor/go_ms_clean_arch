//go:generate mockgen -package blogmock -destination blogmock/blog_repository.go . BlogRepository
package blog

type BlogRepository interface {
	Find(int) (*Blog, error)
	Save(Blog) (int64, error)
}
