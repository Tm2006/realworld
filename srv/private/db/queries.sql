-- name: ListArticles :many
SELECT * FROM articles;

-- name: GetArticle :one
SELECT * FROM articles WHERE slug = ?;

-- name: CreateArticle :one
INSERT INTO articles (slug, title, body) VALUES (?, ?, ?)
RETURNING *;