package skillmodels

import (
	"github.com/criptycpizza7/habit-tracker/common/users"
	usermodels "github.com/criptycpizza7/habit-tracker/common/users/gorm/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Skill struct { // TODO: добавить unique (userID, Name)
	gorm.Model
	ID          uuid.UUID    `gorm:"primary_key;type:uuid;default:gen_random_uuid()"`
	Name        string       `gorm:"type:varchar(255);not null; index:idx_unique_name_userid,unique"`
	UserID      users.UserId `gorm:"index:idx_unique_name_userid,unique"`
	User        usermodels.User
	CurrentTime int `gorm:"default:0"`
	MaxTime     int `gorm:"check:max_time > 0"`
}

type Story struct {
	gorm.Model
	ID      uuid.UUID    `gorm:"primary_key;type:uuid;default:gen_random_uuid()"`
	SkillId uuid.UUID    `gorm:"type:uuid;not null"`
	UserId  users.UserId `gorm:"type:uuid;not null; index:"`
	Time    int
}
