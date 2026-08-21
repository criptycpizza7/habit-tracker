package skills

import (
	skillsschema "github.com/criptycpizza7/habit-tracker/common/skills/schema"
	"github.com/criptycpizza7/habit-tracker/common/users"
)

type SkillController struct {
	skillStore ISkillStore
}

func NewController(skillStore ISkillStore) *SkillController {
	return &SkillController{skillStore}
}

func (c *SkillController) ListSkills(user_id users.UserId) (*skillsschema.SkillsDto, error) {
	return c.skillStore.ListSkills(user_id)
}

func (c *SkillController) AddSkill(skill *skillsschema.AddSkillDto) (*skillsschema.SkillDto, error) {
	return c.skillStore.AddSkill(skill)
}

func (c *SkillController) DeleteSkill(skill *skillsschema.DeleteSkillDto) error {
	return c.skillStore.DeleteSkill(skill)
}

func (c *SkillController) AddTime(time *skillsschema.AddTimeDto) (*skillsschema.AddTimeReturnDto, error) {
	return c.skillStore.AddTime(time)
}

func (c *SkillController) AddTimeBulk(time *skillsschema.AddTimeBulkDto) (*skillsschema.AddTimeBulkReturnDto, error) {
	return c.skillStore.AddTimeBulk(time)
}

func (c *SkillController) ListStory(user_id users.UserId) (*skillsschema.StoryDto, error) {
	story_join, err := c.skillStore.ListStory(user_id)
	story := &skillsschema.StoryDto{Story: make([]*skillsschema.StoryItemDto, 0)}
	if err != nil {
		return nil, err
	}
	for _, item := range *story_join.Story {
		story.Story = append(story.Story, &skillsschema.StoryItemDto{
			Id:          item.ID,
			SkillName:   item.Name,
			Time:        item.Time,
			TimePercent: float32(item.Time) * 100.0 / float32(item.MaxTime),
			CreatedAt:   item.CreatedAt,
		})
	}
	return story, nil
}

func (c *SkillController) DeleteStory(story *skillsschema.DeleteStoryDto) (*skillsschema.AddTimeReturnDto, error) {
	return c.skillStore.DeleteStory(story)
}
