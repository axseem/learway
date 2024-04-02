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
			{"Shabby", "Untidy and in bad condition"},
			{"Troublesome", "Causing a lot of problems for someone"},
			{"Tiresome", "Annoying and making you lose patience"},
			{"Tasteless", "Likely to upset someone"},
			{"Spotless", "Extremely clean"},
			{"Quaint", "Attractive because of being unusual and especially old-fashioned"},
			{"Off-putting", "Slightly unpleasant or worrying so that you do not want to get involved in any way"},
			{"Moderate", "Neither small nor large in size, amount, degree, or strength"},
		},
	})

	return err
}
