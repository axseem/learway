package cmd

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/axseem/learway/internal/database"
	"github.com/axseem/learway/internal/web"
	_ "modernc.org/sqlite"
)

func Serve(port string) error {
	if port == "" {
		port = "1323"
	} else {
		p, err := strconv.Atoi(port)
		if err != nil || p < 1 || p > 65535 {
			return errors.New("invalid port")
		}
	}

	sqlite, err := sql.Open("sqlite", "./dev.db")
	if err != nil {
		return err
	}

	db := database.New(sqlite)

	fmt.Println("Listening on " + port)
	return http.ListenAndServe(":"+port, web.App(db))
}
