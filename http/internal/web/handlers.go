package web

import (
	"github.com/criptycpizza7/habit-tracker/common"
	skillshandlers "github.com/criptycpizza7/habit-tracker/http/internal/skills/handlers"
	usershandlers "github.com/criptycpizza7/habit-tracker/http/internal/users/handlers"
	"github.com/criptycpizza7/habit-tracker/http/internal/web/router"
)

type handlers struct {
	userHandlers   *usershandlers.UserHandlers
	skillsHandlers *skillshandlers.SkillsHandlers
}

func MakeHandlers(ctrl *common.Controllers) *handlers {
	return &handlers{
		userHandlers:   usershandlers.MakeHandler(ctrl.UserController),
		skillsHandlers: skillshandlers.MakeHandlers(ctrl.SkillController),
	}
}

func (h *handlers) RegisterHandlers(router *router.Router) {
	h.userHandlers.RegisterHandlers(router)
	h.skillsHandlers.RegisterHandlers(router)
}
