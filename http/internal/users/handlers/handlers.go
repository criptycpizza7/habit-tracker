package usershandlers

import "github.com/criptycpizza7/habit-tracker/common/users"

type UserHandlers struct {
	ctrl *users.UserController
}

func MakeHandler(ctrl *users.UserController) *UserHandlers {
	return &UserHandlers{ctrl}
}
