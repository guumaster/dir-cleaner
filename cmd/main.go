// Package main contains code tu run dir-cleaner as a CLI command.
package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/urfave/cli/v2"

	"github.com/IGZgustavomarin/dir-cleaner/pkg/cleaner"
)

var version = "dev"
var date = time.Now().Format(time.RFC3339)

func main() {
	cmd := buildCLI(&cleaner.App{})

	if err := cmd.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

// buildCLI creates a CLI app
func buildCLI(app *cleaner.App) *cli.App {
	cwd, _ := os.Getwd()
	d, _ := time.Parse(time.RFC3339, date)
	return &cli.App{
		Name:      "dir-cleaner",
		Usage:     "remove your unused files on your system",
		Version:   version,
		Compiled:  d,
		UsageText: "dir-cleaner [--path <path>] [--depth <num>] [--dry-run] [--pattern <some_path> [--pattern <some_path>]]",
		Authors: []*cli.Author{
			{
				Name:  "Gustavo Marin",
				Email: "gustavo.marin@intelygenz.com",
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "path",
				Usage:       "path where to start the search",
				DefaultText: "$PWD",
				Value:       cwd,
				Aliases:     []string{"p"},
			},

			&cli.StringSliceFlag{
				Name:     "pattern",
				Usage:    "pattern to search (can be repeated for multiple patterns)",
				Required: true,
			},

			&cli.BoolFlag{
				Name:  "dry-run",
				Usage: "just check without deleting data",
				Value: false,
			},

			&cli.IntFlag{
				Name:    "max-depth",
				Usage:   "how many levels to check (use 0 for no max depth)",
				Value:   0,
				Aliases: []string{"d"},
			},

			&cli.BoolFlag{
				Name:  "bytes",
				Usage: "count bytes instead of default blocks of 4K to match 'du' reports",
				Value: false,
			},

			&cli.BoolFlag{
				Name:  "verbose",
				Usage: "print more info into the stdout",
				Value: false,
			},
		},
		Action: func(c *cli.Context) error {
			path, _ := filepath.Abs(c.String("path"))

			opts := &cleaner.Options{
				Path:       path,
				Patterns:   c.StringSlice("pattern"),
				MaxDepth:   c.Int("max-depth"),
				DryRun:     c.Bool("dry-run"),
				Verbose:    c.Bool("verbose"),
				CountBytes: c.Bool("bytes"),
			}
			if opts.DryRun {
				fmt.Println("DURING THIS DRY RUN YOUR DATA WON'T BE HARM IN ANY WAY.")
			}
			cleanStats, err := app.Clean(opts)
			if err != nil {
				return cli.Exit(err.Error(), 1)
			}

			fmt.Println(cleanStats)
			return nil
		},
	}
}
