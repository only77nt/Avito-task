package adapter

import (
	"errors"
	"github.com/only77nt/avito-task/service"
	"gorm.io/gorm"
)

func NewGormStore(db *gorm.DB) *gormStore {
	return &gormStore{
		db,
	}
}

type gormStore struct {
	db *gorm.DB
}

func (g gormStore) SaveUrlInDB(url string, shortUrl string) error {
	result := service.Link{
		Url:      url,
		ShortUrl: shortUrl,
	}

	if err := g.db.Create(result).Error; err != nil {
		return err
	}

	return nil
}

func (g gormStore) FindUrlInDB(shortUrl string) (*string, error) {
	if shortUrl == "" {
		return nil, errors.New("Empty url")
	}

	result := service.Link{}

	if err := g.db.Where("short_url = ?", shortUrl).Find(&result).Error; err != nil {
		return nil, err
	}

	return &result.Url, nil
}
