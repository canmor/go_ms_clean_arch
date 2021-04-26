package usecase

type ShortURL interface {
	Create(url string) (string, error)
}
