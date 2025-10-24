package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/tim2006/realworld/private/types"
)

type Profiles struct {
}

func (p Profiles) Routes(r chi.Router) {
	r.Get("/{username}", p.Get)
	r.Post("/{username}/follow", p.Follow)
	r.Delete("/{username}/follow", p.Unfollow)
}

// Get возвращает профиль пользователя
func (p *Profiles) Get(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	log.Printf("👤 GET /api/profiles/%s - получение профиля", username)

	// Создаем тестовый профиль
	profile := types.Profile{
		Username:  username,
		Bio:       "Bio for " + username,
		Image:     "https://api.realworld.io/images/demo-avatar.png",
		Following: false,
	}

	response := types.ProfileResponse{
		Profile: profile,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("❌ Ошибка кодирования JSON: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	log.Printf("✅ Возвращен профиль: %s", username)
}

// Follow подписывается на пользователя
func (p *Profiles) Follow(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	log.Printf("➕ POST /api/profiles/%s/follow - подписка", username)

	// Возвращаем профиль с флагом following = true
	profile := types.Profile{
		Username:  username,
		Bio:       "Bio for " + username,
		Image:     "https://api.realworld.io/images/demo-avatar.png",
		Following: true,
	}

	response := types.ProfileResponse{
		Profile: profile,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response)
	log.Printf("✅ Подписка на пользователя: %s", username)
}

// Unfollow отписывается от пользователя
func (p *Profiles) Unfollow(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	log.Printf("➖ DELETE /api/profiles/%s/follow - отписка", username)

	// Возвращаем профиль с флагом following = false
	profile := types.Profile{
		Username:  username,
		Bio:       "Bio for " + username,
		Image:     "https://api.realworld.io/images/demo-avatar.png",
		Following: false,
	}

	response := types.ProfileResponse{
		Profile: profile,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response)
	log.Printf("✅ Отписка от пользователя: %s", username)
}
