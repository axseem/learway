package cmd

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/axseem/learway/database"
	"github.com/axseem/learway/view"
	_ "modernc.org/sqlite"
)

func GetDeck(w http.ResponseWriter, r *http.Request) {

}

func Serve() error {
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
		idWildcard := r.PathValue("id")

		id, err := strconv.Atoi(idWildcard)
		if err != nil {
			w.Write([]byte(err.Error()))
			log.Println(err)
		}

		rawDeck, err := db.GetDeck(context.Background(), int64(id))
		if err != nil {
			w.Write([]byte(err.Error()))
			log.Println(err)
		}

		view.DeckPage(rawDeck).Render(r.Context(), w)
	})

	fmt.Println("Listening on :1323")
	return http.ListenAndServe(":1323", r)
}
