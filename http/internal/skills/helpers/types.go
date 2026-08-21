package skillshelpers

import "github.com/google/uuid"

type addSkillInput struct {
	Name    string `json:"name"`
	MaxTime int    `json:"max_time"`
}

type AddSkillInput struct {
	Skill *addSkillInput `in:"body"`
}

type deleteSkillInput struct {
	SkillId uuid.UUID `json:"skill_id"`
}

type DeleteSkill struct {
	SkillId *deleteSkillInput `in:"body"`
}

type addTimeInput struct {
	SkillId uuid.UUID `json:"skill_id"`
	Time    int       `json:"time"`
}

type AddTime struct {
	Time *addTimeInput `in:"body"`
}

type addTimeBulkInput struct {
	SkillIds []*uuid.UUID `json:"skill_ids"`
	Time     int          `json:"time"`
}

type AddTimeBulk struct {
	Time *addTimeBulkInput `in:"body"`
}

type DeleteStory struct {
	StoryId uuid.UUID `in:"path=id"`
}

type SkillOutput struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	CurrentTime int       `json:"current_time"`
	MaxTime     int       `json:"max_time"`
}

type SkillsOutput struct {
	Skills []*SkillOutput `json:"skills"`
}
