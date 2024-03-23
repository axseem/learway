package main

import (
	"log"
	"os"

	"github.com/axseem/learway/cmd"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.App{
		Name:  "lerway",
		Usage: "cli to manage lerway server",
		Commands: []*cli.Command{
			{
				Name:  "serve",
				Usage: "run srever",
				Action: func(ctx *cli.Context) error {
					return cmd.Serve(ctx.Args().First())
				},
			},
			{
				Name:  "seed",
				Usage: "seed database with test data",
				Action: func(ctx *cli.Context) error {
					return cmd.Seed()
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
