-- article.sql

-- name: GetArticleBySlug :one
SELECT 
    a.id,
    a.slug,
    a.title,
    a.description,
    a.body,
    a.created_at,
    a.updated_at,
    a.favorites_count,
    a.author_id,
    u.username,
    u.bio,
    u.image
FROM articles a
JOIN users u ON a.author_id = u.id
WHERE a.slug = ?;


-- name: CreateArticle :one
INSERT INTO articles (slug, title, description, body, author_id)
VALUES (?, ?, ?, ?, ?)
RETURNING id, slug, title, description, body, created_at, updated_at, favorites_count, author_id;

-- name: UpdateArticle :one
UPDATE articles
SET title = ?,
    description = ?,
    body = ?,
    updated_at = CURRENT_TIMESTAMP
WHERE slug = ? AND author_id = ?
RETURNING id, slug, title, description, body, created_at, updated_at, favorites_count, author_id;


-- name: DeleteArticle :exec
DELETE FROM articles
WHERE slug = ? AND author_id = ?;


-- name: CreateFavorite :exec
INSERT INTO favorites (user_id, article_id)
VALUES (?, ?);

-- name: IncrementFavoritesCount :exec
UPDATE articles
SET favorites_count = favorites_count + 1
WHERE id = ?;

-- name: DeleteFavorite :exec
DELETE FROM favorites
WHERE user_id = ? AND article_id = ?;

-- name: DecrementFavoritesCount :exec
UPDATE articles
SET favorites_count = MAX(favorites_count - 1, 0)
WHERE id = ?;

-- name: GetAllTags :many
SELECT id, tag FROM tags;

-- name: CreateTag :one
INSERT INTO tags (tag)
VALUES (?)
RETURNING id, tag;

-- name: GetTagByName :one
SELECT id, tag FROM tags WHERE tag = ?;

-- name: CreateArticleTag :exec
INSERT INTO article_tags (article_id, tag_id)
VALUES (?, ?);

-- name: GetArticleTags :many
SELECT t.id, t.tag
FROM tags t
JOIN article_tags at ON t.id = at.tag_id
WHERE at.article_id = ?;

-- name: ListArticles :many
SELECT 
    a.id,
    a.slug,
    a.title,
    a.description,
    a.body,
    a.created_at,
    a.updated_at,
    a.favorites_count,
    a.author_id,
    u.username,
    u.bio,
    u.image
FROM articles a
JOIN users u ON a.author_id = u.id
ORDER BY a.created_at DESC
LIMIT ? OFFSET ?;

-- name: ListArticlesByAuthor :many
SELECT 
    a.id,
    a.slug,
    a.title,
    a.description,
    a.body,
    a.created_at,
    a.updated_at,
    a.favorites_count,
    a.author_id,
    u.username,
    u.bio,
    u.image
FROM articles a
JOIN users u ON a.author_id = u.id
WHERE u.username = ?
ORDER BY a.created_at DESC
LIMIT ? OFFSET ?;

-- name: CheckIfFavorited :one
SELECT COUNT(*) > 0 AS is_favorited
FROM favorites
WHERE user_id = ? AND article_id = ?;

-- name: CheckIfFollowing :one
SELECT COUNT(*) > 0 AS is_following
FROM follows
WHERE follower_id = ? AND followee_id = ?;

-- name: GetArticleIdBySlug :one
SELECT id FROM articles WHERE slug = ?;

