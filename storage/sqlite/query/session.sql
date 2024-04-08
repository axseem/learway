-- name: GetSessionByID :one
SELECT *
FROM session
WHERE id = ?
LIMIT 1; 

-- name: GetSessionsByUserID :many
SELECT *
FROM session
WHERE user_id = ?
ORDER BY created_at DESC;

-- name: CreateSession :exec
INSERT INTO session (id, user_id, fingerprint, ip, expires_at) VALUES (?, ?, ?, ?, ?);

-- name: UpdateSession :exec
UPDATE session
SET fingerprint = ?, ip = ?, expires_at = ?
WHERE id = ?;

-- name: DeleteSession :exec
DELETE
FROM session
WHERE id = ?;