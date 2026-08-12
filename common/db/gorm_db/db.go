package gormdb

import (
	"github.com/criptycpizza7/habit-tracker/common/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

func New(cfg config.DBConfig, opts ...gorm.Option) *DB {
	gorm_db, err := gorm.Open(postgres.Open(cfg.DSN()), opts...)
	if err != nil {
		panic(err)
	}

	migrate(gorm_db)

	return &DB{gorm_db}
}
