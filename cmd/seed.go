package cmd

import (
	"context"
	"database/sql"

	"github.com/axseem/learway/internal/database"
	"github.com/axseem/learway/internal/model"
	"github.com/axseem/learway/internal/service"
	_ "modernc.org/sqlite"
)

func Seed() error {
	sqlite, err := sql.Open("sqlite", "./dev.db")
	if err != nil {
		return err
	}
	defer sqlite.Close()

	s := service.NewDeckService(database.New(sqlite))

	_, err = s.Create(context.Background(), model.DeckCreateParams{
		Title: "Words for describing things",
		Cards: [][2]string{
			{"Untidy and in bad condition", "Shabby"},
			{"Causing a lot of problems for someone", "Troublesome"},
			{"Annoying and making you lose patience", "Tiresome"},
			{"Likely to upset someone", "Tasteless"},
			{"Extremely clean", "Spotless"},
			{"Attractive because of being unusual and especially old-fashioned", "Quaint"},
			{"Slightly unpleasant or worrying so that you do not want to get involved in any way", "Off-putting"},
			{"Neither small nor large in size, amount, degree, or strength", "Moderate"},
		},
	})

	return err
}
