# 🚀 Быстрая шпаргалка по тестированию API

## Запуск сервера
```bash
cd srv && go run cmd/realworld/main.go
```

## Быстрые тесты с curl

### Базовые проверки
```bash
# Проверка работы сервера
curl http://localhost:8080/api/user

# Проверка CORS
curl -I -X OPTIONS http://localhost:8080/api/user \
  -H "Origin: http://localhost:3000"
```

### User API
```bash
# GET user
curl -i http://localhost:8080/api/user

# UPDATE user  
curl -i -X PUT http://localhost:8080/api/user \
  -H "Content-Type: application/json" \
  -d '{"user": {"email": "test@example.com", "username": "testuser"}}'
```

### Articles API
```bash
# GET articles
curl -i http://localhost:8080/api/articles

# CREATE article
curl -i -X POST http://localhost:8080/api/articles \
  -H "Content-Type: application/json" \
  -d @test-data.json

# GET specific article
curl -i http://localhost:8080/api/articles/test-slug
```

### Profiles API
```bash
# GET profile
curl -i http://localhost:8080/api/profiles/testuser

# FOLLOW user
curl -i -X POST http://localhost:8080/api/profiles/testuser/follow
```

### Tags API
```bash
# GET all tags
curl -i http://localhost:8080/api/tags
```

## Автоматизированное тестирование
```bash
# Запустить все тесты
./test-api.sh

# Проверить только конкретный endpoint
curl -s -w "%{http_code}\n" http://localhost:8080/api/user
```

## Полезные curl опции
```bash
# Показать headers (-i)
curl -i http://localhost:8080/api/user

# Показать только статус код
curl -s -o /dev/null -w "%{http_code}" http://localhost:8080/api/user

# Измерить время ответа
curl -w "Time: %{time_total}s\n" http://localhost:8080/api/user

# Сохранить ответ в файл
curl http://localhost:8080/api/articles > response.json
```

## Ожидаемые результаты (пока)
- Все endpoints: **200 OK** (пустые ответы)
- Несуществующие пути: **404 Not Found**
- CORS headers: присутствуют для localhost:3000