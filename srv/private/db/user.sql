-- user.sql

-- name: GetUser :one
SELECT email, bio, image, username
FROM users
WHERE id = ?;

-- name: UpdateUser :one
UPDATE users
SET email    = ?,
    username = ?,
    password = ?,
    image    = ?,
    bio      = ?
WHERE id = ?
RETURNING email, bio, image, username;

-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE email = ?;

-- name: CreateUser :one
INSERT INTO users (email, username, password)
VALUES (?, ?, ?)
RETURNING *;

-- name: GetUserProfile :one
SELECT
    u.username,
    u.bio,
    u.image,
    CASE WHEN f.follower_id IS NOT NULL THEN 1 ELSE 0 END AS is_following
FROM users u
LEFT JOIN follows f ON u.id = f.followee_id AND f.follower_id = ?
WHERE u.username = ?;

-- name: GetUserProfileById :one
SELECT
    u.username,
    u.bio,
    u.image,
    CASE WHEN f.follower_id IS NOT NULL THEN 1 ELSE 0 END AS is_following
FROM users u
LEFT JOIN follows f ON u.id = f.followee_id AND f.follower_id = ?
WHERE u.id = ?;

-- name: FollowUser :exec
INSERT INTO follows (follower_id, followee_id)
SELECT ?, id FROM users WHERE username = ?;

-- name: UnfollowUser :exec
DELETE FROM follows
WHERE follower_id = ?
  AND followee_id = (SELECT id FROM users WHERE username = ?);