package cmd

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"

	"github.com/axseem/learway/database"
)

func Seed() {
	sqlite, err := sql.Open("sqlite", "./dev.db")
	if err != nil {
		log.Fatal(err)
	}
	db := database.New(sqlite)

	CardsJSON, err := json.Marshal([][2]string{
		{"Top of the card 1", "The other side of the card 1"},
		{"Top of the card 2", "The other side of the card 2"},
		{"Top of the card 3", "The other side of the card 3"},
	})
	if err != nil {
		log.Fatal(err)
	}

	err = db.CreateDeck(context.Background(), database.CreateDeckParams{
		Title: "Demo Deck",
		Cards: string(CardsJSON),
	})
	if err != nil {
		log.Fatal(err)
	}
}
