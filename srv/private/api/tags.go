package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Tags struct {
}

func (t Tags) Routes(r chi.Router) {
	r.Get("/", t.List)
}

func (t *Tags) List(w http.ResponseWriter, r *http.Request) {
}
