package usermodels

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uuid.UUID `gorm:"primary_key;type:uuid;default:gen_random_uuid()"`
	Username string    `gorm:"unique"`
	Password string    `gorm:""`
}
