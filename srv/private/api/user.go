package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/tim2006/realworld/private/types"
)

type User struct {
	DB interface{} // Временно используем interface{}, позже добавим *sql.DB
}

func (u User) Routes(r chi.Router) {
	r.Post("/", u.Register)   // POST /api/users - регистрация
	r.Post("/login", u.Login) // POST /api/users/login - вход
	r.Get("/", u.Get)         // GET /api/user - текущий пользователь
	r.Put("/", u.Put)         // PUT /api/user - обновление
}

// Register регистрирует нового пользователя
func (u *User) Register(w http.ResponseWriter, r *http.Request) {
	log.Println("📝 POST /api/users - регистрация нового пользователя")

	var regRequest types.UserRegisterRequest

	// Парсим JSON из body
	if err := json.NewDecoder(r.Body).Decode(&regRequest); err != nil {
		log.Printf("❌ Ошибка парсинга JSON: %v", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		errorResponse := types.ErrorResponse{
			Errors: map[string][]string{
				"body": {"невозможно распарсить JSON"},
			},
		}
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	// Валидация
	if regRequest.User.Username == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)

		errorResponse := types.ErrorResponse{
			Errors: map[string][]string{
				"username": {"не может быть пустым"},
			},
		}
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	if regRequest.User.Email == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)

		errorResponse := types.ErrorResponse{
			Errors: map[string][]string{
				"email": {"не может быть пустым"},
			},
		}
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	if regRequest.User.Password == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)

		errorResponse := types.ErrorResponse{
			Errors: map[string][]string{
				"password": {"не может быть пустым"},
			},
		}
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	// TODO: Хешировать пароль с bcrypt
	// TODO: Сохранить в базу данных
	// TODO: Сгенерировать JWT токен

	// Пока возвращаем mock ответ
	user := types.UserWithToken{
		ID:       1,
		Email:    regRequest.User.Email,
		Username: regRequest.User.Username,
		Bio:      "",
		Image:    "",
		Token:    "jwt.token.here", // TODO: реальный JWT токен
	}

	response := types.UserAuthResponse{
		User: user,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("❌ Ошибка кодирования JSON: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	log.Printf("✅ Пользователь зарегистрирован: %s", user.Username)
}

// Login выполняет вход пользователя
func (u *User) Login(w http.ResponseWriter, r *http.Request) {
	log.Println("🔐 POST /api/users/login - вход пользователя")

	var loginRequest types.UserLoginRequest

	// Парсим JSON из body
	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		log.Printf("❌ Ошибка парсинга JSON: %v", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		errorResponse := types.ErrorResponse{
			Errors: map[string][]string{
				"body": {"невозможно распарсить JSON"},
			},
		}
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	// Валидация
	if loginRequest.User.Email == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)

		errorResponse := types.ErrorResponse{
			Errors: map[string][]string{
				"email": {"не может быть пустым"},
			},
		}
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	if loginRequest.User.Password == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)

		errorResponse := types.ErrorResponse{
			Errors: map[string][]string{
				"password": {"не может быть пустым"},
			},
		}
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	// TODO: Найти пользователя в БД по email
	// TODO: Проверить пароль с помощью bcrypt
	// TODO: Сгенерировать JWT токен

	// Пока возвращаем mock ответ
	user := types.UserWithToken{
		ID:       1,
		Email:    loginRequest.User.Email,
		Username: "demo",
		Bio:      "Demo user",
		Image:    "https://api.realworld.io/images/demo-avatar.png",
		Token:    "jwt.token.here", // TODO: реальный JWT токен
	}

	response := types.UserAuthResponse{
		User: user,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("❌ Ошибка кодирования JSON: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	log.Printf("✅ Пользователь вошел: %s", user.Username)
}

// Get возвращает информацию о текущем пользователе
func (u *User) Get(w http.ResponseWriter, r *http.Request) {
	log.Println("📥 GET /api/user - получение текущего пользователя")

	// Пока возвращаем тестового пользователя
	user := types.User{
		ID:       1,
		Email:    "demo@realworld.io",
		Username: "demo",
		Bio:      "Demo user for RealWorld API",
		Image:    "https://api.realworld.io/images/demo-avatar.png",
	}

	response := types.UserResponse{
		User: user,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("❌ Ошибка кодирования JSON: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	log.Printf("✅ Пользователь успешно возвращен: %s", user.Username)
}

// Put обновляет информацию о пользователе
func (u *User) Put(w http.ResponseWriter, r *http.Request) {
	log.Println("📝 PUT /api/user - обновление пользователя")

	var updateRequest types.UserUpdateRequest

	// Парсим JSON из body запроса
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

	// Обновляем пользователя (пока просто возвращаем обновленные данные)
	user := types.User{
		ID:       1,
		Email:    updateRequest.User.Email,
		Username: updateRequest.User.Username,
		Bio:      updateRequest.User.Bio,
		Image:    updateRequest.User.Image,
	}

	// Если поля пустые, заполняем значениями по умолчанию
	if user.Email == "" {
		user.Email = "demo@realworld.io"
	}
	if user.Username == "" {
		user.Username = "demo"
	}
	if user.Bio == "" {
		user.Bio = "Updated demo user"
	}
	if user.Image == "" {
		user.Image = "https://api.realworld.io/images/demo-avatar.png"
	}

	response := types.UserResponse{
		User: user,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("❌ Ошибка кодирования JSON: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	log.Printf("✅ Пользователь успешно обновлен: %s", user.Username)
}
