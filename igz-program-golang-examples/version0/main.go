package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v2"
)

var Version = "dev"

func main() {
	app := &cli.App{
		Name:      "dir-cleaner",
		Usage:     "remove some unused files on your system",
		Version:   Version,
		UsageText: "dir-cleaner [--path <path>] [--depth <num>] [--dry-run]",
		Authors: []*cli.Author{
			{
				Name:  "Gustavo Marin",
				Email: "gustavo.marin@intelygenz.com",
			},
		},
		Flags: []cli.Flag{

			&cli.StringFlag{
				Name:    "path",
				Usage:   "path where to start the search",
				Aliases: []string{"p"},
			},

			&cli.BoolFlag{
				Name:  "dry-run",
				Usage: "just check without deleting data",
				Value: false,
			},

			&cli.IntFlag{
				Name:    "max-depth",
				Usage:   "how many levels to check",
				Value:   0,
				Aliases: []string{"d"},
			},
		},
		Action: func(c *cli.Context) error {
			dryRun := c.Bool("dry-run")
			maxDepth := c.Int("max-depth")

			cwd, _ := os.Getwd()
			path := c.String("path")

			if path == "" {
				path = cwd
			}
			path, _ = filepath.Abs(path)

			fmt.Printf("Scanning path: %s\n  params:  [dry-run: %t] [max-depth: %d]\n", path, dryRun, maxDepth)

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
