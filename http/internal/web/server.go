package web

import (
	"net/http"

	"github.com/criptycpizza7/habit-tracker/http/internal"
	"github.com/criptycpizza7/habit-tracker/http/internal/web/router"
)

func RunServer(config internal.HTTPConfig, hndls *handlers, middlewares []func(http.Handler) http.Handler) error {
	r := router.NewRouter()
	r.Use(
		middlewares...,
	)
	hndls.RegisterHandlers(r)
	return http.ListenAndServe(config.Addr(), r)
}
