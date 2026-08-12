package skillshandlers

import (
	skillshelpers "github.com/criptycpizza7/habit-tracker/http/internal/skills/helpers"
	"github.com/criptycpizza7/habit-tracker/http/internal/web/middlewares"
	"github.com/criptycpizza7/habit-tracker/http/internal/web/router"
	"github.com/ggicci/httpin"
	"github.com/go-chi/chi/v5"
)

const (
	SKILLS_ROUTE     = "/skills"
	TIME_ROUTE       = SKILLS_ROUTE + "/time"
	STORY_ROUTE      = "/story"
	STORY_ITEM_ROUTE = "/story/{id}"
)

func (h *SkillsHandlers) RegisterHandlers(router *router.Router) {
	router.Group(
		func(r chi.Router) {
			r.Use(middlewares.AuthMiddleware)

			r.Get(SKILLS_ROUTE, h.ListSkills)
			r.With(
				httpin.NewInput(skillshelpers.AddSkillInput{}),
			).Post(SKILLS_ROUTE, h.AddSkill)
			r.With(
				httpin.NewInput(skillshelpers.DeleteSkill{}),
			).Delete(SKILLS_ROUTE, h.DeleteSkill)
			r.With(
				httpin.NewInput(skillshelpers.AddTime{}),
			).Post(TIME_ROUTE, h.AddTime)
		},
	)

	router.Group(
		func(r chi.Router) {
			r.Use(middlewares.AuthMiddleware)

			r.Get(STORY_ROUTE, h.ListStory)
			r.With(
				httpin.NewInput(skillshelpers.DeleteStory{}),
			).Delete(STORY_ITEM_ROUTE, h.DeleteStory)
		},
	)
}
