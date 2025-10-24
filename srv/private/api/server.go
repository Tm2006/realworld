package api

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
)

type Server struct {
	Address string
	DB      *sql.DB
}

func (s Server) Run() error {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)

	// CORS настройки для разработки
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "http://localhost:5173"}, // React и Vite dev servers
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Максимальное время кеширования preflight запроса
	})
	r.Use(c.Handler)

	// API Routes
	r.Route("/api", func(r chi.Router) {
		r.Route("/users", User{}.Routes) // POST /api/users, POST /api/users/login
		r.Route("/user", User{}.Routes)  // GET /api/user, PUT /api/user
		r.Route("/articles", Articles{}.Routes)
		r.Route("/profiles", Profiles{}.Routes)
		r.Route("/tags", Tags{}.Routes)
	})

	return http.ListenAndServe(s.Address, r)
}
