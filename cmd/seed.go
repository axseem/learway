package cmd

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/axseem/learway/internal/database"
	nanoid "github.com/matoous/go-nanoid/v2"
	_ "modernc.org/sqlite"
)

func Seed() error {
	CardsJSON, err := json.Marshal([][2]string{
		{"Top of the card 1", "The other side of the card 1"},
		{"Top of the card 2", "The other side of the card 2"},
		{"Top of the card 3", "The other side of the card 3"},
	})
	if err != nil {
		return err
	}

	id, err := nanoid.Generate("0123456789abcdefghijklmnopqrstuvwxyz", 8)
	if err != nil {
		return err
	}

	sqlite, err := sql.Open("sqlite", "./dev.db")
	if err != nil {
		return err
	}
	defer sqlite.Close()

	db := database.New(sqlite)

	return db.CreateDeck(context.Background(), database.CreateDeckParams{
		ID:    id,
		Title: "Demo Deck",
		Cards: CardsJSON,
	})
}
