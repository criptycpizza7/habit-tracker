package common

import (
	"github.com/criptycpizza7/habit-tracker/common/skills"
	"github.com/criptycpizza7/habit-tracker/common/users"
)

type IAccessors interface {
	UserAccessor() users.IUserStore
	SkillAccessor() skills.ISkillStore
}
