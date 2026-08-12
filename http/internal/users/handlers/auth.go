package usershandlers

import (
	"log"
	"net/http"

	"github.com/criptycpizza7/habit-tracker/common/db"
	"github.com/criptycpizza7/habit-tracker/common/users"
	userhelpers "github.com/criptycpizza7/habit-tracker/http/internal/users/helpers"
	"github.com/criptycpizza7/habit-tracker/http/internal/web/helpers"
)

func (h *UserHandlers) Register(w http.ResponseWriter, r *http.Request) {
	user_body := helpers.Read[userhelpers.CreateUserInput](r)

	user, err := h.ctrl.Register(user_body.User)
	if err == db.ErrUserAlreadyExists {
		helpers.WriteError(w, http.StatusConflict, "username already taken")
		return
	}
	if err != nil {
		log.Println(err)
		helpers.WriteEmptyError(w, http.StatusInternalServerError)
		return
	}

	helpers.WriteResponse(w, http.StatusCreated, user)
}

func (h *UserHandlers) Login(w http.ResponseWriter, r *http.Request) {
	user_body := helpers.Read[userhelpers.LoginUserInput](r)

	user := users.LoginUserDto(*user_body.User)

	token, err := h.ctrl.Login(&user)
	if err == db.ErrUserNotFound {
		helpers.WriteError(w, http.StatusNotFound, "user not found")
		return
	}
	if err != nil {
		log.Println(err)
		helpers.WriteEmptyError(w, http.StatusInternalServerError)
		return
	}
	token_return := userhelpers.Token{
		JWT: *token,
	}
	helpers.WriteResponse(w, http.StatusOK, token_return)
}
