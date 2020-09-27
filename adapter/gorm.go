package adapter

import "gorm.io/gorm"

func NewGormStore(db *gorm.DB) *gormStore {
	return &gormStore{
		db,
	}
}

type gormStore struct {
	db     *gorm.DB
}