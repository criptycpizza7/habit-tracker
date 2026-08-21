package usershandlers

import (
	"log"
	"net/http"

	"github.com/criptycpizza7/habit-tracker/common/db"
	"github.com/criptycpizza7/habit-tracker/common/users"
	userhelpers "github.com/criptycpizza7/habit-tracker/http/internal/users/helpers"
	"github.com/criptycpizza7/habit-tracker/http/internal/web/helpers"
)

func (h *UserHandlers) Get(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(users.UserId) // TODO: в идеале переделать под свой контекст с типами, и ResponseWriter и Request тоже
	user, err := h.ctrl.Get(user_id)

	if err == db.ErrUserNotFound {
		helpers.WriteError(w, http.StatusNotFound, "user not found")
		return
	}
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	helpers.WriteResponse(w, http.StatusOK, user)
	w.WriteHeader(http.StatusOK)
}

func (h *UserHandlers) Patch(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(users.UserId)
	user_body := helpers.Read[userhelpers.UpdateUserInput](r)

	user, err := h.ctrl.Update(user_id, user_body.User)
	if err == db.ErrUserNotFound {
		helpers.WriteError(w, http.StatusNotFound, "user not found")
		return
	}
	if err == db.ErrUserAlreadyExists {
		helpers.WriteError(w, http.StatusConflict, "username already taken")
		return
	}
	if err != nil {
		log.Println(err)
		helpers.WriteEmptyError(w, http.StatusInternalServerError)
		return
	}
	helpers.WriteResponse(w, http.StatusOK, user)
}

func (h *UserHandlers) Delete(w http.ResponseWriter, r *http.Request) { // TODO: сделать soft delete
	user_id := r.Context().Value("user_id").(users.UserId)

	err := h.ctrl.Delete(user_id)
	if err == db.ErrUserNotFound {
		helpers.WriteError(w, http.StatusNotFound, "user not found")
		return
	}
	if err != nil {
		log.Println(err)
		helpers.WriteEmptyError(w, http.StatusInternalServerError)
		return
	}

	helpers.WriteResponse(w, http.StatusNoContent, nil)
}
