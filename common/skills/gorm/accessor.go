package gormskills

import (
	"github.com/criptycpizza7/habit-tracker/common/db"
	gormdb "github.com/criptycpizza7/habit-tracker/common/db/gorm_db"
	skillmodels "github.com/criptycpizza7/habit-tracker/common/skills/gorm/models"
	skillsschema "github.com/criptycpizza7/habit-tracker/common/skills/schema"
	"github.com/criptycpizza7/habit-tracker/common/users"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func fromGormToDtoSkill(skill *skillmodels.Skill) *skillsschema.SkillDto {
	return &skillsschema.SkillDto{
		Id:          skill.ID,
		UserId:      skill.UserID,
		Name:        skill.Name,
		CurrentTime: skill.CurrentTime,
		MaxTime:     skill.MaxTime,
	}
}

type SkillAccessor struct {
	db *gormdb.DB
}

func NewSkillAccessor(db *gormdb.DB) *SkillAccessor {
	return &SkillAccessor{db}
}

func (a *SkillAccessor) ListSkills(user_id users.UserId) (*skillsschema.SkillsDto, error) {
	skills_result := []skillmodels.Skill{}
	result := a.db.Where("user_id = ?", user_id).Find(&skills_result)
	if result.Error != nil {
		return nil, result.Error
	}

	skills_dto := make([]*skillsschema.SkillDto, len(skills_result))
	for i, skill := range skills_result {
		skills_dto[i] = fromGormToDtoSkill(&skill)
	}
	skills_return := &skillsschema.SkillsDto{
		Skills: skills_dto,
	}
	return skills_return, nil
}

func (a *SkillAccessor) AddSkill(skill *skillsschema.AddSkillDto) (*skillsschema.SkillDto, error) {
	skill_model := &skillmodels.Skill{
		Name:    skill.Name,
		UserID:  skill.UserId,
		MaxTime: skill.MaxTime,
	}
	result := a.db.Create(skill_model)
	if result.Error != nil {
		return nil, result.Error
	}
	return fromGormToDtoSkill(skill_model), nil
}

func (a *SkillAccessor) DeleteSkill(skill *skillsschema.DeleteSkillDto) error {
	result := a.db.Where(
		"id = ? and user_id = ?",
		skill.SkillId,
		skill.UserId,
	).Delete(&skillmodels.Skill{})
	if result.RowsAffected == 0 {
		return db.ErrSkillNotFound
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (a *SkillAccessor) AddTime(time *skillsschema.AddTimeDto) (*skillsschema.AddTimeReturnDto, error) {
	var skill_return *skillsschema.AddTimeReturnDto
	err := a.db.Transaction(func(tx *gorm.DB) error {
		err := a.addStory(time)
		if err != nil {
			return err
		}
		skill_return, err = a.addTime(time)
		return err
	})
	return skill_return, err
}

func (a *SkillAccessor) addTime(time *skillsschema.AddTimeDto) (*skillsschema.AddTimeReturnDto, error) {
	skill := &skillmodels.Skill{}
	result := (a.db.Model(skill).Where(
		"id = ? and user_id = ?",
		time.SkillId,
		time.UserID,
	).
		Clauses(clause.Returning{}).
		UpdateColumn("current_time", gorm.Expr("\"current_time\" + ?", time.Time)))

	if result.Error != nil {
		return nil, result.Error
	}
	skill_return := skillsschema.AddTimeReturnDto{
		CurrentTime: skill.CurrentTime,
		SkillId:     skill.ID,
	}
	return &skill_return, nil
}

func (a *SkillAccessor) addStory(time *skillsschema.AddTimeDto) error {
	story := &skillmodels.Story{
		SkillId: time.SkillId,
		UserId:  time.UserID,
		Time:    time.Time,
	}
	result := a.db.Create(story)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (a *SkillAccessor) ListStory(user_id users.UserId) (*skillsschema.ListStoryJoinDto, error) {
	story := &[]skillsschema.StoryJoinDto{}
	result := a.db.Table("stories st").
		Select("st.time as time, skills.name as name, skills.max_time as max_time").
		Joins("JOIN skills on st.skill_id = skills.id AND skills.user_id = ?", user_id).
		Scan(&story)
	if result.Error != nil {
		return nil, result.Error
	}
	return &skillsschema.ListStoryJoinDto{Story: story}, nil
}

func (a *SkillAccessor) DeleteStory(story *skillsschema.DeleteStoryDto) (*skillsschema.AddTimeReturnDto, error) {
	var skill_return *skillsschema.AddTimeReturnDto
	err := a.db.Transaction(func(tx *gorm.DB) error {
		var err error
		skill_return, err = a.deductTime(story)
		if err != nil {
			return err
		}
		return a.deleteStory(story)
	})
	return skill_return, err
}

func (a *SkillAccessor) deleteStory(story *skillsschema.DeleteStoryDto) error {
	result := a.db.Delete(
		&skillmodels.Story{},
		"id = ? and user_id = ?",
		story.StoryId,
		story.UserId,
	)
	if result.RowsAffected == 0 {
		return db.ErrStoryNotFound
	}
	return result.Error
}

func (a *SkillAccessor) deductTime(story *skillsschema.DeleteStoryDto) (*skillsschema.AddTimeReturnDto, error) {
	story_model := &skillmodels.Story{}
	result := a.db.Find(story_model, "id = ? and user_id = ?", story.StoryId, story.UserId)
	if result.Error != nil {
		return nil, result.Error
	}
	skill_model := &skillmodels.Skill{}
	result = a.db.Model(skill_model).
		Where("id = ? and user_id = ?", story_model.SkillId, story.UserId).
		Clauses(clause.Returning{}).
		UpdateColumn("current_time", gorm.Expr("\"current_time\" - ?", story_model.Time))
	if result.Error != nil {
		return nil, result.Error
	}
	return &skillsschema.AddTimeReturnDto{
		CurrentTime: skill_model.CurrentTime,
		SkillId:     skill_model.ID,
	}, nil
}
