-- name: GetDeck :one
SELECT
  *
FROM
  deck
WHERE
  id = ?
LIMIT
  1;

-- name: ListDecks :many
SELECT
  *
FROM
  deck
ORDER BY
  created_at DESC;

-- name: CreateDeck :exec
INSERT INTO
  deck (id, title, cards)
VALUES
  (?, ?, ?);

-- name: UpdateDeck :exec
UPDATE deck
set
  title = ?,
  cards = ?
WHERE
  id = ?;

-- name: DeleteDeck :exec
DELETE FROM deck
WHERE
  id = ?;