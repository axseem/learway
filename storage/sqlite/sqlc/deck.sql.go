// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: deck.sql

package sqlc

import (
	"context"
)

const createDeck = `-- name: CreateDeck :exec
INSERT INTO deck (id, title, cards) VALUES (?, ?, ?)
`

type CreateDeckParams struct {
	ID    string
	Title string
	Cards []byte
}

func (q *Queries) CreateDeck(ctx context.Context, arg CreateDeckParams) error {
	_, err := q.db.ExecContext(ctx, createDeck, arg.ID, arg.Title, arg.Cards)
	return err
}

const deleteDeck = `-- name: DeleteDeck :exec
DELETE
FROM deck
WHERE id = ?
`

func (q *Queries) DeleteDeck(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteDeck, id)
	return err
}

const getDeck = `-- name: GetDeck :one
SELECT id, title, cards, created_at, updated_at
FROM deck
WHERE id = ?
LIMIT 1
`

func (q *Queries) GetDeck(ctx context.Context, id string) (Deck, error) {
	row := q.db.QueryRowContext(ctx, getDeck, id)
	var i Deck
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Cards,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listDecks = `-- name: ListDecks :many
SELECT id, title, cards, created_at, updated_at
FROM deck
ORDER BY created_at DESC
`

func (q *Queries) ListDecks(ctx context.Context) ([]Deck, error) {
	rows, err := q.db.QueryContext(ctx, listDecks)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Deck
	for rows.Next() {
		var i Deck
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Cards,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateDeck = `-- name: UpdateDeck :exec
UPDATE deck
SET title = ?, cards = ?, updated_at = CURRENT_TIMESTAMP
WHERE id = ?
`

type UpdateDeckParams struct {
	Title string
	Cards []byte
	ID    string
}

func (q *Queries) UpdateDeck(ctx context.Context, arg UpdateDeckParams) error {
	_, err := q.db.ExecContext(ctx, updateDeck, arg.Title, arg.Cards, arg.ID)
	return err
}
