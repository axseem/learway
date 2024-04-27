package service_test

import (
	"context"
	"net"
	"reflect"
	"testing"
	"time"

	"github.com/axseem/learway/core"
	"github.com/axseem/learway/model"
	"github.com/axseem/learway/storage"
	"github.com/axseem/learway/test"
)

func TestDeckGet(t *testing.T) {
	app, clear := test.App()
	defer clear()

	decks := seedDecks(app)

	got, err := app.Service.Deck.Get(context.Background(), "d1")
	if err != nil {
		t.Error(err)
	}
	if !isDeckCreatedCorrectly(decks[0], got) {
		t.Errorf("expected: %v, got: %v", decks[0], got)
	}
}

func TestDeckGetByUsername(t *testing.T) {
	app, clear := test.App()
	defer clear()

	decks := seedDecks(app)

	got, err := app.Service.Deck.GetByUsername(context.Background(), "user1")
	if err != nil {
		t.Error(err)
	}
	if len(got) != 2 {
		t.Error("expected to get 2 decks")
	}
	if !isDeckCreatedCorrectly(decks[0], got[0]) {
		t.Errorf("expected: %v, got: %v", decks[0], got[0])
	}
	if !isDeckCreatedCorrectly(decks[2], got[1]) {
		t.Errorf("expected: %v, got: %v", decks[2], got[1])
	}
}

func TestDeckList(t *testing.T) {
	app, clear := test.App()
	defer clear()

	decks := seedDecks(app)

	got, err := app.Service.Deck.List(context.Background())
	if err != nil {
		t.Error(err)
	}
	if len(got) != 3 {
		t.Error("expected to get 3 decks")
	}
	if !isDeckCreatedCorrectly(decks[0], got[0]) {
		t.Errorf("expected: %v, got: %v", decks[0], got[0])
	}
	if !isDeckCreatedCorrectly(decks[1], got[1]) {
		t.Errorf("expected: %v, got: %v", decks[1], got[1])
	}
	if !isDeckCreatedCorrectly(decks[2], got[2]) {
		t.Errorf("expected: %v, got: %v", decks[2], got[2])
	}
}

func TestDeckGetDelete(t *testing.T) {
	app, clear := test.App()
	defer clear()

	decks := seedDecks(app)

	err := app.Service.Deck.Delete(context.Background(), "d1")
	if err != nil {
		t.Error(err)
	}

	got, err := app.Service.Deck.List(context.Background())
	if err != nil {
		t.Error(err)
	}

	if len(got) != 2 {
		t.Error("expected to get 2 decks")
	}
	if !isDeckCreatedCorrectly(decks[1], got[0]) {
		t.Errorf("expected: %v, got: %v", decks[1], got[0])
	}
	if !isDeckCreatedCorrectly(decks[2], got[1]) {
		t.Errorf("expected: %v, got: %v", decks[2], got[1])
	}
}

func TestDeckUpdated(t *testing.T) {
	app, clear := test.App()
	defer clear()

	_ = seedDecks(app)

	updatedDeck := model.DeckUpdateParams{
		Title: "Test deck 2 updated",
		Cards: model.Cards{
			{"card 1 front updated", "card 1 back updated"},
			{"card 2 front updated", "card 2 back updated"},
			{"card 3 front updated", "card 3 back updated"},
		},
		Subject: "not general",
	}

	updated, err := app.Service.Deck.Update(context.Background(), "d2", updatedDeck)
	if err != nil {
		t.Error(err)
	}

	if !(updatedDeck.Title == updated.Title &&
		reflect.DeepEqual(updatedDeck.Cards, updated.Cards) &&
		updatedDeck.Subject == updated.Subject) {
		t.Errorf("expected: %v, got: %v", updated, updated)
	}
}

func TestDeckSearch(t *testing.T) {
	app, clear := test.App()
	defer clear()

	decks := seedDecks(app)

	found, err := app.Service.Deck.Search(context.Background(), "eck 1")
	if err != nil {
		t.Error(err)
	}

	if !isDeckCreatedCorrectly(decks[0], found[0]) {
		t.Errorf("expected: %v, got: %v", decks[0], found[0])
	}
}

func seedDecks(app *core.App) []storage.DeckCreateParams {
	app.Queries.User.Create(context.Background(), storage.UserCreateParams{
		ID:       "u1",
		Username: "user1",
		Email:    "user1@email.com",
		Password: []byte{},
		Name:     "User",
	})

	decks := []storage.DeckCreateParams{
		{
			ID:     "d1",
			UserID: "u1",
			Title:  "Test deck 1",
			Cards: model.Cards{
				{"card 1 front", "card 1 back"},
				{"card 2 front", "card 2 back"},
				{"card 3 front", "card 3 back"},
			},
			Subject: "general",
		},
		{
			ID:     "d2",
			UserID: "u2",
			Title:  "Test deck 2",
			Cards: model.Cards{
				{"card 1 front", "card 1 back"},
				{"card 2 front", "card 2 back"},
				{"card 3 front", "card 3 back"},
			},
			Subject: "general",
		},
		{
			ID:     "d3",
			UserID: "u1",
			Title:  "Test deck 3",
			Cards: model.Cards{
				{"card 1 front", "card 1 back"},
				{"card 2 front", "card 2 back"},
				{"card 3 front", "card 3 back"},
			},
			Subject: "general",
		},
	}

	for _, deck := range decks {
		err := app.Queries.Deck.Create(context.Background(), deck)
		if err != nil {
			panic(err)
		}
	}

	return decks
}

func isDeckCreatedCorrectly(args storage.DeckCreateParams, deck model.Deck) bool {
	return args.ID == deck.ID &&
		args.UserID == deck.UserID &&
		args.Title == deck.Title &&
		reflect.DeepEqual(args.Cards, deck.Cards) &&
		args.Subject == deck.Subject
}

func TestDeckCreate(t *testing.T) {
	testCases := []struct {
		desc  string
		input model.DeckCreateParams
		err   bool
	}{
		{
			desc: "valid deck",
			input: model.DeckCreateParams{
				SessionID: "a",
				Title:     "Test deck",
				Cards: model.Cards{
					{"card 1 front", "card 1 back"},
					{"card 2 front", "card 2 back"},
					{"card 3 front", "card 3 back"},
				},
				Subject: "general",
			},
			err: false,
		},
		{
			desc: "too short deck",
			input: model.DeckCreateParams{
				SessionID: "a",
				Title:     "Test deck",
				Cards: model.Cards{
					{"card 1 front", "card 1 back"},
				},
				Subject: "general",
			},
			err: true,
		},
		{
			desc: "empty title",
			input: model.DeckCreateParams{
				SessionID: "a",
				Title:     "",
				Cards: model.Cards{
					{"card 1 front", "card 1 back"},
					{"card 2 front", "card 2 back"},
				},
				Subject: "general",
			},
			err: true,
		},
		{
			desc: "empty card side",
			input: model.DeckCreateParams{
				SessionID: "a",
				Title:     "Test deck",
				Cards: model.Cards{
					{"", "card 1 back"},
					{"card 2 front", "card 2 back"},
				},
				Subject: "general",
			},
			err: true,
		},
		{
			desc: "no cards",
			input: model.DeckCreateParams{
				SessionID: "a",
				Title:     "Test deck",
				Cards:     model.Cards{},
				Subject:   "general",
			},
			err: true,
		},
		{
			desc: "invalid session",
			input: model.DeckCreateParams{
				SessionID: "b",
				Title:     "Test deck",
				Cards: model.Cards{
					{"card 1 front", "card 1 back"},
					{"card 2 front", "card 2 back"},
				},
				Subject: "general",
			},
			err: true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			app, clear := test.App()
			defer clear()

			err := app.Queries.Session.Create(context.Background(), storage.SessionCreateParams{
				ID:          "a",
				UserID:      "a0",
				Fingerprint: []byte{},
				IP:          net.IP{},
				ExpiresAt:   time.Now().Add(time.Hour),
			})
			if err != nil {
				panic(err)
			}

			output, err := app.Service.Deck.Create(context.Background(), tC.input)
			if err != nil {
				if tC.err {
					return
				} else {
					t.Error(err)
				}
			}

			if tC.input.Title != output.Title ||
				!reflect.DeepEqual(tC.input.Cards, output.Cards) ||
				tC.input.Subject != output.Subject {
				t.Errorf("expected: %v, got: %v", tC.input, output)
			}
		})
	}
}
