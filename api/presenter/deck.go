package presenter

import "github.com/axseem/learway/model"

type DeckCreateParams struct {
	Title string      `json:"title"`
	Cards model.Cards `json:"cards"`
}
