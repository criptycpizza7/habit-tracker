package usershandlers

import (
	userhelpers "github.com/criptycpizza7/habit-tracker/http/internal/users/helpers"
	"github.com/criptycpizza7/habit-tracker/http/internal/web/middlewares"
	"github.com/criptycpizza7/habit-tracker/http/internal/web/router"
	"github.com/ggicci/httpin"
	"github.com/go-chi/chi/v5"
)

const (
	USER_ROUTE  = "/users/{id}"
	ME_ROUTE    = "/users/me"
	USERS_ROUTE = "/users"
	LOGIN_ROUTE = USERS_ROUTE + "/login"
)

func (h *UserHandlers) RegisterHandlers(router *router.Router) {
	router.With(
		httpin.NewInput(userhelpers.CreateUserInput{}),
	).Post(USERS_ROUTE, h.Register)

	router.With(
		httpin.NewInput(userhelpers.LoginUserInput{}),
	).Post(LOGIN_ROUTE, h.Login)

	router.Group(
		func(router chi.Router) {
			router.Use(middlewares.AuthMiddleware)

			router.With(
				httpin.NewInput(userhelpers.GetUser{}),
			).Get(ME_ROUTE, h.Get)

			router.With(
				httpin.NewInput(userhelpers.UpdateUserInput{}),
			).Patch(USER_ROUTE, h.Patch)

			router.With(
				httpin.NewInput(userhelpers.GetUser{}),
			).Delete(USER_ROUTE, h.Delete)
		},
	)
}
