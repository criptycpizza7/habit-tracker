package skillshandlers

import (
	"log"
	"net/http"

	"github.com/google/uuid"

	"github.com/criptycpizza7/habit-tracker/common/db"
	skillsschema "github.com/criptycpizza7/habit-tracker/common/skills/schema"
	"github.com/criptycpizza7/habit-tracker/common/users"
	skillshelpers "github.com/criptycpizza7/habit-tracker/http/internal/skills/helpers"
	"github.com/criptycpizza7/habit-tracker/http/internal/web/helpers"
)

func (h *SkillsHandlers) ListSkills(w http.ResponseWriter, r *http.Request) { // TODO: пагинация
	user_id := r.Context().Value("user_id").(users.UserId)
	skills_list, err := h.ctrl.ListSkills(user_id)
	if err != nil {
		log.Println(err)
		helpers.WriteEmptyError(w, http.StatusInternalServerError)
		return
	}
	helpers.WriteResponse(w, http.StatusOK, skills_list)
}

func (h *SkillsHandlers) AddSkill(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(users.UserId)
	skill_input := helpers.Read[skillshelpers.AddSkillInput](r)
	skill_dto := &skillsschema.AddSkillDto{
		UserId:  user_id,
		Name:    skill_input.Skill.Name,
		MaxTime: skill_input.Skill.MaxTime,
	}
	skill, err := h.ctrl.AddSkill(skill_dto)
	if err != nil {
		log.Println(err)
		helpers.WriteEmptyError(w, http.StatusInternalServerError)
		return
	}
	skill_output := skillsschema.AddSkillOutput{
		Id:      skill.Id,
		Name:    skill.Name,
		MaxTime: skill.MaxTime,
	}
	helpers.WriteResponse(w, http.StatusOK, skill_output)
}

func (h *SkillsHandlers) DeleteSkill(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(users.UserId)
	skill_input := helpers.Read[skillshelpers.DeleteSkill](r)

	skill := &skillsschema.DeleteSkillDto{
		UserId:  user_id,
		SkillId: skill_input.SkillId.SkillId,
	}
	err := h.ctrl.DeleteSkill(skill)
	if err == db.ErrSkillNotFound {
		helpers.WriteError(w, http.StatusNotFound, "skill not found")
		return
	}
	if err != nil {
		log.Println(err)
		helpers.WriteEmptyError(w, http.StatusInternalServerError)
		return
	}
	helpers.WriteResponse(w, http.StatusNoContent, nil)
}

func (h *SkillsHandlers) AddTime(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(users.UserId)
	skill_input := helpers.Read[skillshelpers.AddTime](r).Time
	skill_dto := &skillsschema.AddTimeDto{
		UserID:  user_id,
		SkillId: skill_input.SkillId,
		Time:    skill_input.Time,
	}
	skill, err := h.ctrl.AddTime(skill_dto)
	if err == db.ErrSkillNotFound {
		helpers.WriteError(w, http.StatusNotFound, "skill not found")
		return
	}
	if err != nil {
		log.Println(err)
		helpers.WriteEmptyError(w, http.StatusInternalServerError)
		return
	}
	helpers.WriteResponse(w, http.StatusOK, skill)
}

func (h *SkillsHandlers) AddTimeBulk(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(users.UserId)
	skill_input := helpers.Read[skillshelpers.AddTimeBulk](r).Time
	skill_dto := &skillsschema.AddTimeBulkDto{
		UserId: user_id,
		Time:   skill_input.Time,
		Skills: make([]*uuid.UUID, len(skill_input.SkillIds)),
	}

	copy(skill_dto.Skills, skill_input.SkillIds)

	skill, err := h.ctrl.AddTimeBulk(skill_dto)
	if err == db.ErrSkillNotFound {
		helpers.WriteError(w, http.StatusNotFound, "skill not found")
		return
	}
	if err != nil {
		log.Println(err)
		helpers.WriteEmptyError(w, http.StatusInternalServerError)
		return
	}
	helpers.WriteResponse(w, http.StatusOK, skill)
}
