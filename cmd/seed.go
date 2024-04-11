package cmd

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/axseem/learway/model"
	"github.com/axseem/learway/service"
	"github.com/axseem/learway/storage/sqlite"
	"github.com/go-playground/validator/v10"
	_ "modernc.org/sqlite"
)

func Seed() error {
	sqliteDB, err := sql.Open("sqlite", "./dev.db")
	if err != nil {
		return err
	}
	defer sqliteDB.Close()

	v := validator.New(validator.WithRequiredStructEnabled())
	s := service.NewDeckService(sqlite.New(sqliteDB), v)

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

	if err != nil {
		return err
	}

	fmt.Println("database has been seeded")
	return nil
}
