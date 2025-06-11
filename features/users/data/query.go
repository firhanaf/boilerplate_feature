package data

import "gorm.io/gorm"

type UserQuery struct {
	db *gorm.DB
}

func NewUserQuery(db *gorm.DB) *UserQuery {
	return &UserQuery{
		db: db,
	}
}
