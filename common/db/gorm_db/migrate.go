package gormdb

import (
	skillmodels "github.com/criptycpizza7/habit-tracker/common/skills/gorm/models"
	usermodels "github.com/criptycpizza7/habit-tracker/common/users/gorm/models"
	"gorm.io/gorm"
)

func migrate(db *gorm.DB) { // TODO: переделать под нормальные миграции
	db.AutoMigrate(
		&usermodels.User{},
		&skillmodels.Skill{},
		&skillmodels.Story{},
	)
}
