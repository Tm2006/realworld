package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/tim2006/realworld/private/types"
)

type Tags struct {
}

func (t Tags) Routes(r chi.Router) {
	r.Get("/", t.List)
}

// List возвращает список всех тегов
func (t *Tags) List(w http.ResponseWriter, r *http.Request) {
	log.Println("🏷️ GET /api/tags - получение списка тегов")

	// Создаем список тестовых тегов
	tags := []string{
		"dragons",
		"training",
		"golang",
		"programming",
		"javascript",
		"react",
		"nodejs",
		"frontend",
		"backend",
		"fullstack",
		"api",
		"database",
		"tutorial",
		"beginner",
		"advanced",
	}

	response := types.TagsResponse{
		Tags: tags,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("❌ Ошибка кодирования JSON: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	log.Printf("✅ Возвращено %d тегов", len(tags))
}
