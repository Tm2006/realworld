package types

import "time"

// User представляет пользователя в системе
type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Bio      string `json:"bio"`
	Image    string `json:"image"`
}

// UserResponse - ответ API для пользователя
type UserResponse struct {
	User User `json:"user"`
}

// UserUpdateRequest - запрос на обновление пользователя
type UserUpdateRequest struct {
	User struct {
		Email    string `json:"email,omitempty"`
		Username string `json:"username,omitempty"`
		Bio      string `json:"bio,omitempty"`
		Image    string `json:"image,omitempty"`
	} `json:"user"`
}

// Article представляет статью
type Article struct {
	ID             int64     `json:"id"`
	Slug           string    `json:"slug"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	Body           string    `json:"body"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	Favorited      bool      `json:"favorited"`
	FavoritesCount int       `json:"favoritesCount"`
	Author         Author    `json:"author"`
	TagList        []string  `json:"tagList"`
}

// Author представляет автора статьи
type Author struct {
	Username  string `json:"username"`
	Bio       string `json:"bio"`
	Image     string `json:"image"`
	Following bool   `json:"following"`
}

// ArticleResponse - ответ API для одной статьи
type ArticleResponse struct {
	Article Article `json:"article"`
}

// ArticlesResponse - ответ API для списка статей
type ArticlesResponse struct {
	Articles      []Article `json:"articles"`
	ArticlesCount int       `json:"articlesCount"`
}

// ArticleCreateRequest - запрос на создание статьи
type ArticleCreateRequest struct {
	Article struct {
		Title       string   `json:"title"`
		Description string   `json:"description"`
		Body        string   `json:"body"`
		TagList     []string `json:"tagList"`
	} `json:"article"`
}

// Profile представляет профиль пользователя
type Profile struct {
	Username  string `json:"username"`
	Bio       string `json:"bio"`
	Image     string `json:"image"`
	Following bool   `json:"following"`
}

// ProfileResponse - ответ API для профиля
type ProfileResponse struct {
	Profile Profile `json:"profile"`
}

// TagsResponse - ответ API для списка тегов
type TagsResponse struct {
	Tags []string `json:"tags"`
}

// ErrorResponse - стандартный ответ об ошибке
type ErrorResponse struct {
	Errors map[string][]string `json:"errors"`
}

// UserRegisterRequest - запрос на регистрацию пользователя
type UserRegisterRequest struct {
	User struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	} `json:"user"`
}

// UserLoginRequest - запрос на вход пользователя
type UserLoginRequest struct {
	User struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	} `json:"user"`
}

// UserAuthResponse - ответ API для аутентифицированного пользователя
type UserAuthResponse struct {
	User UserWithToken `json:"user"`
}

// UserWithToken - пользователь с JWT токеном
type UserWithToken struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Bio      string `json:"bio"`
	Image    string `json:"image"`
	Token    string `json:"token"`
}
