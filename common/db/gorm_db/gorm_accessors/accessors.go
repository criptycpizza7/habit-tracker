package gormaccessors

import (
	gormdb "github.com/criptycpizza7/habit-tracker/common/db/gorm_db"
	"github.com/criptycpizza7/habit-tracker/common/skills"
	gormskills "github.com/criptycpizza7/habit-tracker/common/skills/gorm"
	"github.com/criptycpizza7/habit-tracker/common/users"
	gormusers "github.com/criptycpizza7/habit-tracker/common/users/gorm"
)

type Accessors struct {
	userAccessor  *gormusers.UserAccessor
	skillAccessor *gormskills.SkillAccessor
}

func New(db *gormdb.DB) *Accessors {
	user_accessor := gormusers.NewUserAccessor(db)
	skill_accessor := gormskills.NewSkillAccessor(db)
	return &Accessors{
		userAccessor:  user_accessor,
		skillAccessor: skill_accessor,
	}
}

func (a *Accessors) UserAccessor() users.IUserStore {
	return a.userAccessor
}

func (a *Accessors) SkillAccessor() skills.ISkillStore {
	return a.skillAccessor
}
