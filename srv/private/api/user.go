package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type User struct {
}

func (u User) Routes(r chi.Router) {
	r.Get("/", u.Get)
	r.Put("/", u.Put)
}

func (u *User) Get(w http.ResponseWriter, r *http.Request) {

}

func (u *User) Put(w http.ResponseWriter, r *http.Request) {

}
