-- name: GetUserByID :one
SELECT *
FROM user
WHERE id = ?
LIMIT 1;

-- name: GetUserByUsername :one
SELECT *
FROM user
WHERE username = ?
LIMIT 1;

-- name: GetUserByEmail :one
SELECT *
FROM user
WHERE email = ?
LIMIT 1;

-- name: CreateUser :exec
INSERT INTO user (id, username, email, password, name) VALUES (?, ?, ?, ?, ?);

-- name: UpdateUserPassword :exec
UPDATE user
SET password = ?
WHERE id = ?;

-- name: UpdateUserUsername :exec
UPDATE user
SET username = ?
WHERE id = ?;

-- name: UpdateUserProfile :exec
UPDATE user
SET name = ?, description = ?, picture = ?
WHERE id = ?;

-- name: DeleteUser :exec
DELETE
FROM user
WHERE id = ?;