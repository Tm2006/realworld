package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Articles struct {
}

func (a Articles) Routes(r chi.Router) {
	r.Get("/", a.List)
	r.Post("/", a.Create)
	r.Get("/{slug}", a.Get)
	r.Put("/{slug}", a.Update)
	r.Delete("/{slug}", a.Delete)
	r.Post("/{slug}/favorite", a.Favorite)
	r.Delete("/{slug}/favorite", a.Unfavorite)
	r.Get("/{slug}/comments", a.ListComments)
	r.Post("/{slug}/comments", a.CreateComment)
	r.Delete("/{slug}/comments/{id}", a.DeleteComment)
}

func (a *Articles) List(w http.ResponseWriter, r *http.Request) {
}

func (a *Articles) Create(w http.ResponseWriter, r *http.Request) {
}

func (a *Articles) Get(w http.ResponseWriter, r *http.Request) {
}

func (a *Articles) Update(w http.ResponseWriter, r *http.Request) {
}

func (a *Articles) Delete(w http.ResponseWriter, r *http.Request) {
}

func (a *Articles) Favorite(w http.ResponseWriter, r *http.Request) {
}

func (a *Articles) Unfavorite(w http.ResponseWriter, r *http.Request) {
}

func (a *Articles) ListComments(w http.ResponseWriter, r *http.Request) {
}

func (a *Articles) CreateComment(w http.ResponseWriter, r *http.Request) {
}

func (a *Articles) DeleteComment(w http.ResponseWriter, r *http.Request) {
}
