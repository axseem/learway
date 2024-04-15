-- name: GetDeck :one
SELECT *
FROM deck
WHERE id = ?
LIMIT 1;

-- name: GetDeckByUserID :many
SELECT *
FROM deck
WHERE user_id = ?
ORDER BY created_at DESC;

-- name: ListDecks :many
SELECT *
FROM deck
ORDER BY created_at DESC;

-- name: CreateDeck :exec
INSERT INTO deck (id, user_id, title, cards, subject) VALUES (?, ?, ?, ?, ?);

-- name: UpdateDeck :exec
UPDATE deck
SET title = ?, cards = ?, subject = ?, updated_at = CURRENT_TIMESTAMP
WHERE id = ?;

-- name: DeleteDeck :exec
DELETE
FROM deck
WHERE id = ?;

-- name: SearchDeck :many
SELECT *
FROM deck
WHERE title
LIKE ?;