package db

import (
	"gorm.io/gorm"
)

type Initializer interface {
	InitDb() (*gorm.DB, error, func() error)
}

func New() Initializer {
	return &AuthenticationDbInitializer{}
}
