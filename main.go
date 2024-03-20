package main

import (
	"log"

	"github.com/axseem/learway/cmd"
)

func main() {
	if err := cmd.Serve(); err != nil {
		log.Fatal(err)
	}
}
