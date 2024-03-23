package cmd

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"

	"github.com/axseem/learway/internal/database"
	nanoid "github.com/matoous/go-nanoid/v2"
)

func Seed() error {
	sqlite, err := sql.Open("sqlite", "./dev.db")
	if err != nil {
		return err
	}
	db := database.New(sqlite)

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
		log.Println(err)
	}

	return db.CreateDeck(context.Background(), database.CreateDeckParams{
		ID:    id,
		Title: "Demo Deck",
		Cards: string(CardsJSON),
	})
}
