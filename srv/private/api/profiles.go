package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Profiles struct {
}

func (p Profiles) Routes(r chi.Router) {
	r.Get("/{username}", p.Get)
	r.Post("/{username}/follow", p.Follow)
	r.Delete("/{username}/follow", p.Unfollow)
}

func (p *Profiles) Get(w http.ResponseWriter, r *http.Request) {
}

func (p *Profiles) Follow(w http.ResponseWriter, r *http.Request) {
}

func (p *Profiles) Unfollow(w http.ResponseWriter, r *http.Request) {
}
