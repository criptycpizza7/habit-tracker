package skills

import (
	skillsschema "github.com/criptycpizza7/habit-tracker/common/skills/schema"
	"github.com/criptycpizza7/habit-tracker/common/users"
)

type ISkillStore interface {
	ListSkills(user_id users.UserId) (*skillsschema.SkillsDto, error)
	AddSkill(skill *skillsschema.AddSkillDto) (*skillsschema.SkillDto, error)
	DeleteSkill(skill *skillsschema.DeleteSkillDto) error
	AddTime(time *skillsschema.AddTimeDto) (*skillsschema.AddTimeReturnDto, error)
	AddTimeBulk(time *skillsschema.AddTimeBulkDto) (*skillsschema.AddTimeBulkReturnDto, error)

	ListStory(user_id users.UserId) (*skillsschema.ListStoryJoinDto, error)
	DeleteStory(story *skillsschema.DeleteStoryDto) (*skillsschema.AddTimeReturnDto, error)
}
