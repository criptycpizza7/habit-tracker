package common

import (
	"github.com/criptycpizza7/habit-tracker/common/skills"
	"github.com/criptycpizza7/habit-tracker/common/users"
)

type Controllers struct {
	UserController  *users.UserController
	SkillController *skills.SkillController
}

func MakeControllers(accessors IAccessors) *Controllers {
	return &Controllers{
		UserController:  users.NewController(accessors.UserAccessor()),
		SkillController: skills.NewController(accessors.SkillAccessor()),
	}
}
