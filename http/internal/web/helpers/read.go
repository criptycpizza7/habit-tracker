package helpers

import (
	"net/http"

	"github.com/ggicci/httpin"
)

func Read[T any](r *http.Request) *T { // TODO: валидировать
	return r.Context().Value(httpin.Input).(*T)
}
