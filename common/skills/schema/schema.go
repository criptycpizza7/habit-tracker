package skillsschema

import (
	"time"

	"github.com/google/uuid"

	"github.com/criptycpizza7/habit-tracker/common/users"
)

// TODO: Разобрать файл
type AddSkillDto struct {
	UserId  users.UserId `json:"user_id"`
	Name    string       `json:"name"`
	MaxTime int          `json:"max_time"`
}

type DeleteSkillDto struct {
	UserId  users.UserId `json:"user_id"`
	SkillId uuid.UUID    `json:"skill_id"`
}

type SkillDto struct {
	Id          uuid.UUID    `json:"id"`
	UserId      users.UserId `json:"user_id"`
	Name        string       `json:"name"`
	CurrentTime int          `json:"current_time"`
	MaxTime     int          `json:"max_time"`
}

type SkillsDto struct {
	Skills []*SkillDto `json:"skills"`
}

type StoryItemDto struct {
	Id          uuid.UUID `json:"id"`
	SkillName   string    `json:"skill_name"`
	Time        int       `json:"time"`
	TimePercent float32   `json:"time_percent"`
	CreatedAt   time.Time `json:"created_at"`
}

type StoryDto struct {
	Story []*StoryItemDto `json:"story"`
}

type AddTimeDto struct {
	UserID  users.UserId `json:"user_id"`
	SkillId uuid.UUID    `json:"skill_id"`
	Time    int          `json:"time"`
}

type AddTimeBulkDto struct {
	UserId users.UserId `json:"user_id"`
	Time   int          `json:"time"`
	Skills []*uuid.UUID `json:"skills"`
}

type DeleteStoryDto struct {
	UserId  users.UserId `json:"user_id"`
	StoryId uuid.UUID    `json:"story_id"`
}

type AddSkillOutput struct {
	Id      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	MaxTime int       `json:"max_time"`
}

type AddTimeReturnDto struct {
	CurrentTime int       `json:"current_time"`
	SkillId     uuid.UUID `json:"skill_id"`
}

type AddTimeBulkReturnItemDto struct {
	CurrentTime int        `json:"current_time"`
	SkillId     *uuid.UUID `json:"skill_id"`
}
type AddTimeBulkReturnDto struct {
	Items []*AddTimeBulkReturnItemDto `json:"items"`
}

type StoryJoinDto struct {
	ID        uuid.UUID
	Time      int
	Name      string
	MaxTime   int
	CreatedAt time.Time
}

type ListStoryJoinDto struct {
	Story *[]StoryJoinDto
}
