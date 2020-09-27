package service

type Link struct {
	Url      string
	ShortUrl string
}

type Store interface {
	SaveUrlInDB(url string, shortUrl string) error
	FindUrlInDB(shortUrl string) (*string, error)
}
