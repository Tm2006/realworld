# 🧪 Инструкция по тестированию RealWorld API

## 📋 Подготовка к тестированию

### 1. Запуск сервера

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

Сервер будет работать на `http://localhost:8080`

### 2. Инструменты для тестирования

#### A) **curl** (командная строка)
```bash
# Проверка доступности
curl --version
```

#### B) **HTTPie** (более удобный)
```bash
# Установка (если нет)
sudo apt install httpie

# Использование
http GET localhost:8080/api/user
```

#### C) **Postman** (графический интерфейс)
- Скачать с [postman.com](https://www.postman.com/)
- Импортировать коллекцию API

---

## 🔍 Тестирование endpoints

### **1. Базовые проверки**

#### Проверка работоспособности сервера
```bash
# Простой запрос к корню (должен вернуть 404 - это нормально)
curl -i http://localhost:8080/

# Проверка API префикса
curl -i http://localhost:8080/api/
```

#### Проверка CORS headers
```bash
# OPTIONS запрос для проверки CORS
curl -i -X OPTIONS http://localhost:8080/api/user \
  -H "Origin: http://localhost:3000" \
  -H "Access-Control-Request-Method: GET"
```

**Ожидаемые headers:**
```
Access-Control-Allow-Origin: http://localhost:3000
Access-Control-Allow-Methods: GET, POST, PUT, DELETE, OPTIONS
```

---

### **2. User endpoints**

#### GET /api/user
```bash
# Получение текущего пользователя (пока пустой handler)
curl -i http://localhost:8080/api/user
```

**Ожидаемый ответ:**
```
HTTP/1.1 200 OK
Content-Length: 0
```

#### PUT /api/user
```bash
# Обновление пользователя
curl -i -X PUT http://localhost:8080/api/user \
  -H "Content-Type: application/json" \
  -d '{"user": {"email": "test@example.com", "username": "testuser"}}'
```

---

### **3. Articles endpoints**

#### GET /api/articles
```bash
# Получение списка статей
curl -i http://localhost:8080/api/articles
```

#### POST /api/articles
```bash
# Создание новой статьи
curl -i -X POST http://localhost:8080/api/articles \
  -H "Content-Type: application/json" \
  -d '{
    "article": {
      "title": "Test Article",
      "description": "Test description",
      "body": "Test body content",
      "tagList": ["test", "demo"]
    }
  }'
```

#### GET /api/articles/test-article
```bash
# Получение статьи по slug
curl -i http://localhost:8080/api/articles/test-article
```

---

### **4. Profiles endpoints**

#### GET /api/profiles/username
```bash
# Получение профиля пользователя
curl -i http://localhost:8080/api/profiles/testuser
```

#### POST /api/profiles/username/follow
```bash
# Подписка на пользователя
curl -i -X POST http://localhost:8080/api/profiles/testuser/follow \
  -H "Authorization: Token jwt.token.here"
```

---

### **5. Tags endpoints**

#### GET /api/tags
```bash
# Получение всех тегов
curl -i http://localhost:8080/api/tags
```

---

## 📊 Проверка ответов API

### **Успешные ответы:**
- `200 OK` - Успешный запрос
- `201 Created` - Ресурс создан
- `204 No Content` - Успешно, но нет контента

### **Ошибки клиента:**
- `400 Bad Request` - Неверный запрос
- `401 Unauthorized` - Не авторизован
- `403 Forbidden` - Доступ запрещен
- `404 Not Found` - Ресурс не найден
- `422 Unprocessable Entity` - Ошибка валидации

### **Ошибки сервера:**
- `500 Internal Server Error` - Ошибка сервера

---

## 🧩 Автоматизированные тесты

### **1. Создание test script**

Создайте файл `test-api.sh`:

```bash
#!/bin/bash

# Цвета для вывода
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

API_URL="http://localhost:8080"

echo -e "${YELLOW}🧪 Testing RealWorld API${NC}"
echo "================================"

# Функция для тестирования endpoint
test_endpoint() {
    local method=$1
    local endpoint=$2
    local expected_status=$3
    local description=$4
    
    echo -n "Testing $method $endpoint... "
    
    response=$(curl -s -w "%{http_code}" -X $method "$API_URL$endpoint")
    status_code="${response: -3}"
    
    if [ "$status_code" -eq "$expected_status" ]; then
        echo -e "${GREEN}✅ PASS${NC} ($status_code)"
    else
        echo -e "${RED}❌ FAIL${NC} (expected $expected_status, got $status_code)"
    fi
}

# Базовые тесты
echo -e "\n${YELLOW}📋 Basic Tests${NC}"
test_endpoint "GET" "/api/user" 200 "Get current user"
test_endpoint "PUT" "/api/user" 200 "Update user"
test_endpoint "GET" "/api/articles" 200 "Get articles"
test_endpoint "GET" "/api/profiles/testuser" 200 "Get profile"
test_endpoint "GET" "/api/tags" 200 "Get tags"

# CORS тесты
echo -e "\n${YELLOW}🌐 CORS Tests${NC}"
cors_response=$(curl -s -I -X OPTIONS "$API_URL/api/user" \
  -H "Origin: http://localhost:3000" \
  -H "Access-Control-Request-Method: GET")

if echo "$cors_response" | grep -q "Access-Control-Allow-Origin"; then
    echo -e "CORS headers: ${GREEN}✅ PASS${NC}"
else
    echo -e "CORS headers: ${RED}❌ FAIL${NC}"
fi

echo -e "\n${YELLOW}🏁 Testing completed!${NC}"
```

**Запуск тестов:**
```bash
chmod +x test-api.sh
./test-api.sh
```

---

### **2. Тестирование с JSON данными**

Создайте файл `test-data.json`:

```json
{
  "user": {
    "email": "test@example.com",
    "username": "testuser",
    "password": "password123",
    "bio": "Test user bio",
    "image": "https://example.com/avatar.jpg"
  },
  "article": {
    "title": "How to train your dragon",
    "description": "Ever wonder how?",
    "body": "Very carefully.",
    "tagList": ["dragons", "training"]
  }
}
```

**Тестирование с файлом:**
```bash
# POST запрос с данными из файла
curl -i -X POST http://localhost:8080/api/articles \
  -H "Content-Type: application/json" \
  -d @test-data.json
```

---

## 🐛 Отладка проблем

### **Проблема: Connection refused**
```bash
# Проверить, запущен ли сервер
ps aux | grep realworld

# Проверить порт
netstat -tlnp | grep :8080
# или
lsof -i :8080
```

### **Проблема: CORS ошибки**
```bash
# Проверить CORS headers
curl -I -X OPTIONS http://localhost:8080/api/user \
  -H "Origin: http://localhost:3000"
```

### **Проблема: 404 ошибки**
```bash
# Убедиться, что используете правильный префикс /api/
curl -i http://localhost:8080/api/user  # ✅ Правильно
curl -i http://localhost:8080/user      # ❌ Неправильно
```

### **Проблема: Медленные ответы**
```bash
# Измерить время ответа
curl -w "@curl-format.txt" http://localhost:8080/api/user
```

Создайте `curl-format.txt`:
```
     time_namelookup:  %{time_namelookup}\n
        time_connect:  %{time_connect}\n
     time_appconnect:  %{time_appconnect}\n
    time_pretransfer:  %{time_pretransfer}\n
       time_redirect:  %{time_redirect}\n
  time_starttransfer:  %{time_starttransfer}\n
                     ----------\n
          time_total:  %{time_total}\n
```

---

## 📈 Мониторинг API

### **Проверка логов сервера**
Сервер выводит логи в консоль:
```
2025/09/29 14:36:59 "GET /api/user HTTP/1.1" from [::1]:48350 - 200 0B in 125µs
```

### **Проверка базы данных**
```bash
# Подключиться к SQLite
sqlite3 realworld.db

# Посмотреть таблицы
.tables

# Проверить пользователей
SELECT * FROM users;

# Выход
.exit
```

---

## ✅ Чек-лист тестирования

- [ ] Сервер запускается без ошибок
- [ ] База данных инициализируется
- [ ] CORS headers настроены правильно
- [ ] Все endpoints отвечают (пусть и пустыми ответами)
- [ ] JSON запросы обрабатываются
- [ ] Ошибки обрабатываются корректно
- [ ] Логирование работает
- [ ] База данных создается и подключается

---

## 🚀 Следующие шаги

После базового тестирования можно:

1. **Реализовать handlers** - добавить логику в пустые методы
2. **Добавить аутентификацию** - JWT токены
3. **Написать unit тесты** - автоматические тесты Go кода
4. **Создать Postman коллекцию** - для удобного тестирования
5. **Добавить интеграционные тесты** - полный цикл тестирования

**Удачного тестирования! 🎯**