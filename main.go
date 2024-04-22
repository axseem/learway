package main

import (
	"log"
	"os"

	"github.com/axseem/learway/cmd"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.App{
		Name:  "learway",
		Usage: "cli to manage learway server",
		Commands: []*cli.Command{
			{
				Name:  "dev",
				Usage: "run in development suitable environment",
				Action: func(ctx *cli.Context) error {
					return cmd.Dev()
				},
			},
			{
				Name:  "seed",
				Usage: "seed database with mock data",
				Action: func(ctx *cli.Context) error {
					return cmd.Seed()
				},
			},
			{
				Name:  "serve",
				Usage: "run server in production ready mode",
				Action: func(ctx *cli.Context) error {
					return cmd.Serve()
				},
			},
			{
				Name:  "migrate",
				Usage: "apply migration queries to database",
				Action: func(ctx *cli.Context) error {
					return cmd.Migrate(ctx.Args().First())
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
