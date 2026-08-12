package router

import (
	"github.com/go-chi/chi/v5"
)

type Router struct {
	chi.Router
}

func NewRouter() *Router {
	return &Router{
		Router: chi.NewRouter(),
	}
}
