package skillshandlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/criptycpizza7/habit-tracker/common/db"
	skillsschema "github.com/criptycpizza7/habit-tracker/common/skills/schema"
	"github.com/criptycpizza7/habit-tracker/common/users"
	skillshelpers "github.com/criptycpizza7/habit-tracker/http/internal/skills/helpers"
	"github.com/criptycpizza7/habit-tracker/http/internal/web/helpers"
)

func (h *SkillsHandlers) ListStory(w http.ResponseWriter, r *http.Request) { // TODO: пагинация
	user_id := r.Context().Value("user_id").(users.UserId)
	fmt.Println(h.ctrl)
	story, err := h.ctrl.ListStory(user_id)
	if err != nil {
		log.Println(err)
		helpers.WriteEmptyError(w, http.StatusInternalServerError)
		return
	}
	helpers.WriteResponse(w, http.StatusOK, story)
}

func (h *SkillsHandlers) DeleteStory(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(users.UserId)
	story_input := helpers.Read[skillshelpers.DeleteStory](r)
	story := &skillsschema.DeleteStoryDto{
		UserId:  user_id,
		StoryId: story_input.StoryId,
	}
	skill, err := h.ctrl.DeleteStory(story)
	if err == db.ErrStoryNotFound {
		helpers.WriteError(w, http.StatusNotFound, "story not found")
		return
	}
	if err != nil {
		log.Println(err)
		helpers.WriteEmptyError(w, http.StatusInternalServerError)
		return
	}
	helpers.WriteResponse(w, http.StatusOK, skill)
}
