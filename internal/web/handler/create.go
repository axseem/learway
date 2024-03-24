package handler

import (
	"context"
	"encoding/json"
	"io"

	"github.com/axseem/learway/internal/database"
	"github.com/axseem/learway/internal/web/view"
	"github.com/labstack/echo/v4"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

func CreatePage(c echo.Context) error {
	return render(c, view.CreatePage())
}

func CreateDeck(db *database.Queries) echo.HandlerFunc {
	return func(c echo.Context) error {
		b, err := io.ReadAll(c.Request().Body)
		if err != nil {
			return err
		}

		deckParams := struct {
			Title string
			Cards [][2]string
		}{}
		if err = json.Unmarshal(b, &deckParams); err != nil {
			return err
		}

		cards, err := json.Marshal(deckParams.Cards)
		if err != nil {
			return err
		}

		id, err := gonanoid.Generate("0123456789abcdefghijklmnopqrstuvwxyz", 8)
		if err != nil {
			return err
		}

		return db.CreateDeck(context.Background(), database.CreateDeckParams{
			ID:    id,
			Title: deckParams.Title,
			Cards: cards,
		})
	}
}
