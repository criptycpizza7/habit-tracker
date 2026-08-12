package skillshandlers

import "github.com/criptycpizza7/habit-tracker/common/skills"

type SkillsHandlers struct {
	ctrl *skills.SkillController
}

func MakeHandlers(ctrl *skills.SkillController) *SkillsHandlers {
	return &SkillsHandlers{ctrl}
}
