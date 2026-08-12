package users

import (
	"time"
)

type UserDto struct {
	ID        UserId    `json:"id"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateUserDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginUserDto CreateUserDto

type LoginUserReturnDto struct {
	ID        UserId    `json:"id"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Password  string    `json:"password"`
}

type UpdateUserDto struct {
	Username *string `json:"username"`
	Password *string `json:"password"`
}
