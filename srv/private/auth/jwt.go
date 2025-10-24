package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Секретный ключ для подписи JWT токенов
// В production это должно быть в переменных окружения!
var jwtSecret = []byte("your-secret-key-change-this-in-production")

// Claims представляет данные JWT токена
type Claims struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}

// GenerateToken создает JWT токен для пользователя
func GenerateToken(userID int64, username, email string) (string, error) {
	// Устанавливаем claims (данные токена)
	claims := Claims{
		UserID:   userID,
		Username: username,
		Email:    email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * 7 * time.Hour)), // 7 дней
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	// Создаем токен с claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Подписываем токен секретным ключом
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateToken проверяет и парсит JWT токен
func ValidateToken(tokenString string) (*Claims, error) {
	// Парсим токен
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	// Извлекаем claims
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrSignatureInvalid
}
