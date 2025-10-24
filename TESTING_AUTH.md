# 🧪 Тестирование регистрации и входа пользователей

## 📋 Что мы реализовали

### ✅ Готово:
- **Типы данных** для регистрации и логина
- **JWT утилиты** для генерации токенов
- **Endpoints**:
  - `POST /api/users` - регистрация
  - `POST /api/users/login` - вход
- **Валидация** полей (email, username, password)
- **CORS middleware** для фронтенда

### ⚠️ Пока без реализации (TODO):
- Хеширование паролей (bcrypt)
- Сохранение в базу данных
- Генерация реальных JWT токенов
- Проверка уникальности email/username

**Но это не мешает протестировать API!** Endpoints работают и возвращают правильную структуру данных.

---

## 🚀 Шаг 1: Запуск сервера

### В первом терминале:

```bash
# Перейти в папку сервера
cd /home/tim/src/my/realworld/srv

# Запустить сервер
go run cmd/realworld/main.go
```

**Ожидаемый вывод:**
```
✅ Database connected successfully
🚀 Starting server on :8080
```

Сервер будет работать и логировать все запросы.

---

## 🧪 Шаг 2: Тестирование регистрации

### Во втором терминале:

### **Тест 1: Успешная регистрация**

```bash
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{
    "user": {
      "username": "testuser",
      "email": "test@example.com",
      "password": "password123"
    }
  }' | jq .
```

**Ожидаемый ответ:**
```json
{
  "user": {
    "id": 1,
    "email": "test@example.com",
    "username": "testuser",
    "bio": "",
    "image": "",
    "token": "jwt.token.here"
  }
}
```

**Статус код:** `201 Created`

---

### **Тест 2: Регистрация без username**

```bash
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{
    "user": {
      "email": "test@example.com",
      "password": "password123"
    }
  }' | jq .
```

**Ожидаемый ответ:**
```json
{
  "errors": {
    "username": ["не может быть пустым"]
  }
}
```

**Статус код:** `422 Unprocessable Entity`

---

### **Тест 3: Регистрация без email**

```bash
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{
    "user": {
      "username": "testuser",
      "password": "password123"
    }
  }' | jq .
```

**Ожидаемый ответ:**
```json
{
  "errors": {
    "email": ["не может быть пустым"]
  }
}
```

---

### **Тест 4: Регистрация без password**

```bash
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{
    "user": {
      "username": "testuser",
      "email": "test@example.com"
    }
  }' | jq .
```

**Ожидаемый ответ:**
```json
{
  "errors": {
    "password": ["не может быть пустым"]
  }
}
```

---

## 🔐 Шаг 3: Тестирование входа

### **Тест 5: Успешный вход**

```bash
curl -X POST http://localhost:8080/api/users/login \
  -H "Content-Type: application/json" \
  -d '{
    "user": {
      "email": "test@example.com",
      "password": "password123"
    }
  }' | jq .
```

**Ожидаемый ответ:**
```json
{
  "user": {
    "id": 1,
    "email": "test@example.com",
    "username": "demo",
    "bio": "Demo user",
    "image": "https://api.realworld.io/images/demo-avatar.png",
    "token": "jwt.token.here"
  }
}
```

**Статус код:** `200 OK`

---

### **Тест 6: Вход без email**

```bash
curl -X POST http://localhost:8080/api/users/login \
  -H "Content-Type: application/json" \
  -d '{
    "user": {
      "password": "password123"
    }
  }' | jq .
```

**Ожидаемый ответ:**
```json
{
  "errors": {
    "email": ["не может быть пустым"]
  }
}
```

---

### **Тест 7: Вход без password**

```bash
curl -X POST http://localhost:8080/api/users/login \
  -H "Content-Type: application/json" \
  -d '{
    "user": {
      "email": "test@example.com"
    }
  }' | jq .
```

**Ожидаемый ответ:**
```json
{
  "errors": {
    "password": ["не может быть пустым"]
  }
}
```

---

## 📊 Проверка статус кодов

### Показать статус код вместе с ответом:

```bash
# Регистрация с отображением статус кода
curl -i -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{
    "user": {
      "username": "newuser",
      "email": "new@example.com",
      "password": "secure123"
    }
  }'
```

Смотрите на первую строку - должно быть `HTTP/1.1 201 Created`

---

## 🔍 Просмотр логов сервера

В первом терминале (где запущен сервер) вы увидите логи:

```
📝 POST /api/users - регистрация нового пользователя
✅ Пользователь зарегистрирован: testuser
"POST http://localhost:8080/api/users HTTP/1.1" from [::1]:12345 - 201 158B in 1.234ms

🔐 POST /api/users/login - вход пользователя
✅ Пользователь вошел: demo
"POST http://localhost:8080/api/users/login HTTP/1.1" from [::1]:12346 - 200 245B in 0.567ms
```

---

## 🐛 Возможные проблемы

### Проблема: Connection refused
```bash
curl: (7) Failed to connect to localhost port 8080
```

**Решение:** Убедитесь, что сервер запущен в первом терминале.

---

### Проблема: 404 Not Found

**Проверьте URL:**
- ✅ Правильно: `http://localhost:8080/api/users`
- ❌ Неправильно: `http://localhost:8080/users`
- ❌ Неправильно: `http://localhost:8080/api/user` (для регистрации)

---

### Проблема: Сервер падает или зависает

**Перезапустите сервер:**
```bash
# Ctrl+C в первом терминале
# Затем запустите снова
go run cmd/realworld/main.go
```

---

## 📝 Все команды скопом (для быстрого копирования)

### Запуск сервера:
```bash
cd /home/tim/src/my/realworld/srv && go run cmd/realworld/main.go
```

### Регистрация:
```bash
curl -X POST http://localhost:8080/api/users -H "Content-Type: application/json" -d '{"user":{"username":"testuser","email":"test@example.com","password":"password123"}}' | jq .
```

### Вход:
```bash
curl -X POST http://localhost:8080/api/users/login -H "Content-Type: application/json" -d '{"user":{"email":"test@example.com","password":"password123"}}' | jq .
```

### Проверка статус кода:
```bash
curl -i -X POST http://localhost:8080/api/users -H "Content-Type: application/json" -d '{"user":{"username":"user1","email":"user1@test.com","password":"pass123"}}'
```

---

## ✅ Чек-лист проверки

После тестирования убедитесь:

- [ ] Сервер запускается без ошибок
- [ ] Регистрация возвращает статус `201 Created`
- [ ] Регистрация возвращает правильную структуру JSON с user и token
- [ ] Валидация работает (ошибки при пустых полях)
- [ ] Вход возвращает статус `200 OK`
- [ ] Вход возвращает правильную структуру JSON
- [ ] Логи сервера показывают эмодзи и сообщения
- [ ] CORS headers присутствуют (можно проверить с `-i`)

---

## 🎯 Следующие шаги

После успешного тестирования мы добавим:

1. **Реальную работу с БД** - сохранение пользователей
2. **Bcrypt хеширование** - безопасное хранение паролей
3. **JWT токены** - настоящие токены вместо mock
4. **Middleware аутентификации** - защита endpoints
5. **Проверку уникальности** - email и username должны быть уникальными

**Удачного тестирования! 🚀**