# 🌟 RealWorld Full-Stack Application

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![React](https://img.shields.io/badge/React-20232A?style=for-the-badge&logo=react&logoColor=61DAFB)
![TypeScript](https://img.shields.io/badge/TypeScript-007ACC?style=for-the-badge&logo=typescript&logoColor=white)
![SQLite](https://img.shields.io/badge/SQLite-07405E?style=for-the-badge&logo=sqlite&logoColor=white)

> Полнофункциональное приложение-блог для изучения современной full-stack разработки

## 📋 Описание проекта

**RealWorld** - это клон Medium.com, созданный для демонстрации и изучения full-stack разработки. Проект включает в себя все функции реального приложения: аутентификацию, CRUD операции, комментарии, лайки, профили пользователей и многое другое.

### 🎯 Цели проекта
- Изучение современного стека технологий
- Практика Clean Architecture
- Понимание взаимодействия frontend и backend
- Освоение DevOps практик

## 🏗️ Архитектура

### Backend (Go)
- **Framework**: Chi router для HTTP обработки
- **Database**: SQLite с SQLC для генерации кода
- **Authentication**: JWT токены
- **Architecture**: Clean Architecture с разделением на слои

### Frontend (React) - _В разработке_
- **Framework**: React 18 + TypeScript
- **Build Tool**: Vite
- **Styling**: TailwindCSS
- **State Management**: React Context + Custom Hooks
- **HTTP Client**: Axios

## 🚀 Быстрый старт

### Предварительные требования
- Go 1.19+ 
- Node.js 18+ (для frontend)
- Git

### Запуск Backend

```bash
# Клонировать репозиторий
git clone https://github.com/YOUR_USERNAME/realworld.git
cd realworld

# Запустить Go сервер
cd srv
go mod tidy
go run cmd/realworld/main.go
```

Сервер будет доступен на `http://localhost:8080`

### API Endpoints

#### Аутентификация
- `POST /users` - Регистрация
- `POST /users/login` - Вход
- `GET /user` - Текущий пользователь
- `PUT /user` - Обновление профиля

#### Статьи
- `GET /articles` - Список статей
- `POST /articles` - Создание статьи
- `GET /articles/:slug` - Получение статьи
- `PUT /articles/:slug` - Обновление статьи
- `DELETE /articles/:slug` - Удаление статьи

#### Профили
- `GET /profiles/:username` - Профиль пользователя
- `POST /profiles/:username/follow` - Подписка
- `DELETE /profiles/:username/follow` - Отписка

## 📊 База данных

Проект использует SQLite с следующими основными таблицами:

- `users` - Пользователи
- `articles` - Статьи
- `comments` - Комментарии
- `tags` - Теги
- `favorites` - Лайки статей
- `follows` - Подписки пользователей

## 🛠️ Разработка

### Структура проекта

```
realworld/
├── srv/                    # Go Backend
│   ├── cmd/realworld/      # Точка входа
│   ├── private/
│   │   ├── api/           # HTTP handlers
│   │   ├── db/            # База данных
│   │   └── types/         # Типы данных
│   └── go.mod
├── web/                   # React Frontend (в разработке)
├── LEARNING_PLAN.md       # План изучения
└── README.md
```

### Запуск в режиме разработки

```bash
# Backend
cd srv
go run cmd/realworld/main.go

# Frontend (когда будет готов)
cd web
npm run dev
```

## 🧪 Тестирование

```bash
# Unit тесты Backend
cd srv
go test ./...

# Frontend тесты (в будущем)
cd web
npm test
```

## 🐳 Docker

```bash
# Запуск всего приложения
docker-compose up -d

# Только backend
docker build -t realworld-api ./srv
docker run -p 8080:8080 realworld-api
```

## 📈 Статус разработки

### ✅ Завершено
- [x] Базовая структура Go API
- [x] Схема базы данных
- [x] HTTP роутинг
- [x] Подготовка к GitHub

### 🔄 В процессе
- [ ] Реализация API endpoints
- [ ] JWT аутентификация
- [ ] Frontend приложение

### 📋 Планируется
- [ ] Комментарии и лайки
- [ ] Система тегов
- [ ] Профили пользователей
- [ ] Real-time уведомления
- [ ] Мобильная версия

## 🤝 Вклад в проект

1. Fork репозиторий
2. Создайте feature branch (`git checkout -b feature/amazing-feature`)
3. Commit изменения (`git commit -m 'Add amazing feature'`)
4. Push в branch (`git push origin feature/amazing-feature`)
5. Откройте Pull Request

## 📚 Обучающие материалы

- [План изучения](LEARNING_PLAN.md) - детальный roadmap
- [RealWorld Spec](https://realworld-docs.netlify.app/) - спецификация API
- [Go Documentation](https://golang.org/doc/)
- [React Documentation](https://react.dev/)

## 📄 Лицензия

Этот проект распространяется под лицензией MIT. См. файл [LICENSE](LICENSE) для деталей.

## 👨‍💻 Автор

**Tim** - [GitHub Profile](https://github.com/YOUR_USERNAME)

## 🙏 Благодарности

- [RealWorld](https://github.com/gothinkster/realworld) за концепцию и спецификацию
- [Chi](https://github.com/go-chi/chi) за отличный HTTP роутер
- [SQLC](https://sqlc.dev/) за генерацию кода из SQL

---

⭐ Если проект был полезен, поставьте звездочку!
