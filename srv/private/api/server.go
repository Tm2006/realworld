package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	Address string
}

func (s Server) Run() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Route("/user", User{}.Routes)
	r.Route("/articles", Articles{}.Routes)
	r.Route("/profiles", Profiles{}.Routes)
	r.Route("/tags", Tags{}.Routes)

	http.ListenAndServe(s.Address, r)
}
