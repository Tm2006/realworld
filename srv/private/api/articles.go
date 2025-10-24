package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/tim2006/realworld/private/types"
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

// List возвращает список статей
func (a *Articles) List(w http.ResponseWriter, r *http.Request) {
	log.Println("📥 GET /api/articles - получение списка статей")

	// Создаем тестовые статьи
	articles := []types.Article{
		{
			ID:             1,
			Slug:           "how-to-train-your-dragon",
			Title:          "How to train your dragon",
			Description:    "Ever wonder how?",
			Body:           "Very carefully. This is a comprehensive guide...",
			CreatedAt:      time.Now().Add(-24 * time.Hour),
			UpdatedAt:      time.Now().Add(-24 * time.Hour),
			Favorited:      false,
			FavoritesCount: 5,
			Author: types.Author{
				Username:  "demo",
				Bio:       "Demo author",
				Image:     "https://api.realworld.io/images/demo-avatar.png",
				Following: false,
			},
			TagList: []string{"dragons", "training"},
		},
		{
			ID:             2,
			Slug:           "go-programming-basics",
			Title:          "Go Programming Basics",
			Description:    "Learn Go programming language",
			Body:           "Go is a great language for backend development...",
			CreatedAt:      time.Now().Add(-12 * time.Hour),
			UpdatedAt:      time.Now().Add(-12 * time.Hour),
			Favorited:      true,
			FavoritesCount: 12,
			Author: types.Author{
				Username:  "gopher",
				Bio:       "Go enthusiast",
				Image:     "https://api.realworld.io/images/gopher.png",
				Following: true,
			},
			TagList: []string{"golang", "programming"},
		},
	}

	response := types.ArticlesResponse{
		Articles:      articles,
		ArticlesCount: len(articles),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("❌ Ошибка кодирования JSON: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	log.Printf("✅ Возвращено %d статей", len(articles))
}

// Create создает новую статью
func (a *Articles) Create(w http.ResponseWriter, r *http.Request) {
	log.Println("📝 POST /api/articles - создание новой статьи")

	var createRequest types.ArticleCreateRequest

	if err := json.NewDecoder(r.Body).Decode(&createRequest); err != nil {
		log.Printf("❌ Ошибка парсинга JSON: %v", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		errorResponse := types.ErrorResponse{
			Errors: map[string][]string{
				"body": {"can't be parsed"},
			},
		}
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	// Создаем slug из заголовка
	slug := strings.ToLower(strings.ReplaceAll(createRequest.Article.Title, " ", "-"))
	slug = strings.ReplaceAll(slug, "'", "")

	// Создаем новую статью
	article := types.Article{
		ID:             3, // В реальности будет из БД
		Slug:           slug,
		Title:          createRequest.Article.Title,
		Description:    createRequest.Article.Description,
		Body:           createRequest.Article.Body,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
		Favorited:      false,
		FavoritesCount: 0,
		Author: types.Author{
			Username:  "demo",
			Bio:       "Demo author",
			Image:     "https://api.realworld.io/images/demo-avatar.png",
			Following: false,
		},
		TagList: createRequest.Article.TagList,
	}

	response := types.ArticleResponse{
		Article: article,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("❌ Ошибка кодирования JSON: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	log.Printf("✅ Создана статья: %s", article.Title)
}

// Get возвращает статью по slug
func (a *Articles) Get(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")
	log.Printf("📥 GET /api/articles/%s - получение статьи", slug)

	// Тестовая статья
	article := types.Article{
		ID:             1,
		Slug:           slug,
		Title:          "Article: " + slug,
		Description:    "Description for " + slug,
		Body:           "This is the body content for article: " + slug,
		CreatedAt:      time.Now().Add(-24 * time.Hour),
		UpdatedAt:      time.Now().Add(-12 * time.Hour),
		Favorited:      false,
		FavoritesCount: 7,
		Author: types.Author{
			Username:  "demo",
			Bio:       "Demo author",
			Image:     "https://api.realworld.io/images/demo-avatar.png",
			Following: false,
		},
		TagList: []string{"demo", "test"},
	}

	response := types.ArticleResponse{
		Article: article,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("❌ Ошибка кодирования JSON: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	log.Printf("✅ Возвращена статья: %s", article.Title)
}

// Update обновляет статью
func (a *Articles) Update(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")
	log.Printf("📝 PUT /api/articles/%s - обновление статьи", slug)

	var updateRequest types.ArticleCreateRequest

	if err := json.NewDecoder(r.Body).Decode(&updateRequest); err != nil {
		log.Printf("❌ Ошибка парсинга JSON: %v", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		errorResponse := types.ErrorResponse{
			Errors: map[string][]string{
				"body": {"can't be parsed"},
			},
		}
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	// Обновляем статью
	article := types.Article{
		ID:             1,
		Slug:           slug,
		Title:          updateRequest.Article.Title,
		Description:    updateRequest.Article.Description,
		Body:           updateRequest.Article.Body,
		CreatedAt:      time.Now().Add(-24 * time.Hour),
		UpdatedAt:      time.Now(), // Обновляем время
		Favorited:      false,
		FavoritesCount: 7,
		Author: types.Author{
			Username:  "demo",
			Bio:       "Demo author",
			Image:     "https://api.realworld.io/images/demo-avatar.png",
			Following: false,
		},
		TagList: updateRequest.Article.TagList,
	}

	response := types.ArticleResponse{
		Article: article,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("❌ Ошибка кодирования JSON: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	log.Printf("✅ Обновлена статья: %s", article.Title)
}

// Delete удаляет статью
func (a *Articles) Delete(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")
	log.Printf("🗑️ DELETE /api/articles/%s - удаление статьи", slug)

	w.WriteHeader(http.StatusNoContent)
	log.Printf("✅ Статья удалена: %s", slug)
}

// Favorite добавляет статью в избранное
func (a *Articles) Favorite(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")
	log.Printf("❤️ POST /api/articles/%s/favorite - добавление в избранное", slug)

	// Возвращаем статью с обновленным статусом избранного
	article := types.Article{
		ID:             1,
		Slug:           slug,
		Title:          "Article: " + slug,
		Description:    "Description for " + slug,
		Body:           "Body content for " + slug,
		CreatedAt:      time.Now().Add(-24 * time.Hour),
		UpdatedAt:      time.Now().Add(-12 * time.Hour),
		Favorited:      true, // Добавлено в избранное
		FavoritesCount: 8,    // Увеличиваем счетчик
		Author: types.Author{
			Username:  "demo",
			Bio:       "Demo author",
			Image:     "https://api.realworld.io/images/demo-avatar.png",
			Following: false,
		},
		TagList: []string{"demo", "test"},
	}

	response := types.ArticleResponse{
		Article: article,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response)
	log.Printf("✅ Статья добавлена в избранное: %s", slug)
}

// Unfavorite убирает статью из избранного
func (a *Articles) Unfavorite(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")
	log.Printf("💔 DELETE /api/articles/%s/favorite - удаление из избранного", slug)

	// Возвращаем статью с обновленным статусом избранного
	article := types.Article{
		ID:             1,
		Slug:           slug,
		Title:          "Article: " + slug,
		Description:    "Description for " + slug,
		Body:           "Body content for " + slug,
		CreatedAt:      time.Now().Add(-24 * time.Hour),
		UpdatedAt:      time.Now().Add(-12 * time.Hour),
		Favorited:      false, // Убрано из избранного
		FavoritesCount: 6,     // Уменьшаем счетчик
		Author: types.Author{
			Username:  "demo",
			Bio:       "Demo author",
			Image:     "https://api.realworld.io/images/demo-avatar.png",
			Following: false,
		},
		TagList: []string{"demo", "test"},
	}

	response := types.ArticleResponse{
		Article: article,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response)
	log.Printf("✅ Статья убрана из избранного: %s", slug)
}

// ListComments возвращает комментарии к статье (заглушка)
func (a *Articles) ListComments(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")
	log.Printf("💬 GET /api/articles/%s/comments - получение комментариев", slug)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"comments": []}`))
}

// CreateComment создает комментарий (заглушка)
func (a *Articles) CreateComment(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")
	log.Printf("💬 POST /api/articles/%s/comments - создание комментария", slug)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"comment": {"id": 1, "body": "Test comment", "createdAt": "2025-10-02T12:00:00Z"}}`))
}

// DeleteComment удаляет комментарий (заглушка)
func (a *Articles) DeleteComment(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")
	id := chi.URLParam(r, "id")
	log.Printf("🗑️ DELETE /api/articles/%s/comments/%s - удаление комментария", slug, id)

	w.WriteHeader(http.StatusNoContent)
}
