package cmd

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/axseem/learway/database"
	"github.com/axseem/learway/view"
	nanoid "github.com/matoous/go-nanoid/v2"
	_ "modernc.org/sqlite"
)

func GetDeck(w http.ResponseWriter, r *http.Request) {

}

func Serve(port string) error {
	sqlite, err := sql.Open("sqlite", "./dev.db")
	if err != nil {
		log.Fatal(err)
	}
	db := database.New(sqlite)

	r := http.NewServeMux()

	fs := http.FileServer(http.Dir("./assets"))
	r.Handle("/assets/", http.StripPrefix("/assets/", fs))

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		rawDecks, err := db.ListDecks(context.Background())
		if err != nil {
			log.Println(err)
		}

		view.IndexPage(rawDecks).Render(r.Context(), w)
	})

	r.HandleFunc("GET /deck/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		rawDeck, err := db.GetDeck(context.Background(), id)
		if err != nil {
			w.Write([]byte(err.Error()))
			log.Println(err)
		}

		view.DeckPage(rawDeck).Render(r.Context(), w)
	})

	r.HandleFunc("GET /create", func(w http.ResponseWriter, r *http.Request) {
		view.CreatePage().Render(r.Context(), w)
	})

	r.HandleFunc("POST /create", func(w http.ResponseWriter, r *http.Request) {
		b, err := io.ReadAll(r.Body)
		if err != nil {
			w.Write([]byte(err.Error()))
			log.Println(err)
		}

		deckParams := struct {
			Title string
			Cards [][2]string
		}{}

		err = json.Unmarshal(b, &deckParams)
		if err != nil {
			w.Write([]byte(err.Error()))
			log.Println(err)
		}

		c, err := json.Marshal(deckParams.Cards)
		if err != nil {
			w.Write([]byte(err.Error()))
			log.Println(err)
		}

		id, err := nanoid.Generate("0123456789abcdefghijklmnopqrstuvwxyz", 8)
		if err != nil {
			w.Write([]byte(err.Error()))
			log.Println(err)
		}

		err = db.CreateDeck(context.Background(), database.CreateDeckParams{
			ID:    id,
			Title: deckParams.Title,
			Cards: string(c),
		})
		if err != nil {
			w.Write([]byte(err.Error()))
			log.Println(err)
		}
	})

	if port == "" {
		port = "1323"
	}

	fmt.Println("Listening on " + port)
	return http.ListenAndServe(":"+port, r)
}
