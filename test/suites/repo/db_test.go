package repo_test

import "gorm.io/gorm"

type MockDB struct {
	*gorm.DB
}

func (db *MockDB) Create(value interface{}) (tx *gorm.DB) {
	return nil
}
