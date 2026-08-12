package userhelpers

import (
	"github.com/criptycpizza7/habit-tracker/common/users"
)

type GetUser struct {
	Authorization string `in:"header=Authorization"`
}

type CreateUserInput struct {
	User *users.CreateUserDto `in:"body=json"`
}

type LoginUserInput CreateUserInput

type UpdateUserInput struct {
	User          *users.UpdateUserDto `in:"body=json"`
	Authorization string               `in:"header=Authorization"`
}

type Token struct {
	JWT string `json:"jwt"`
}
