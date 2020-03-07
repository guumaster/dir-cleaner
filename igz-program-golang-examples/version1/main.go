package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"

	"github.com/urfave/cli/v2"
)

var Version = "dev"

var reMatch = regexp.MustCompile("node_modules")
var reEndMatch = regexp.MustCompile("node_modules$")

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
			rootPath, _ := filepath.Abs(path)

			var size int64
			var files int
			var removePaths []string

			stats := &Stats{rootPath, removePaths, files, size}
			err := filepath.Walk(rootPath, func(srcPath string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}

				if srcPath == rootPath {
					return nil
				}

				meta := getMeta(rootPath, srcPath, info.IsDir())

				depth := maxDepth
				if !info.IsDir() {
					depth++
				}

				if maxDepth != 0 && meta.Level > depth {
					return nil
				}

				match := reMatch.MatchString(meta.Dir)
				if !match {
					return nil
				}

				if info.IsDir() && reEndMatch.MatchString(meta.Dir) {
					stats.RemovePaths = append(stats.RemovePaths, srcPath)
					// fmt.Printf("[lvl %2d] path [%s] \n", meta.Level, meta.Dir+meta.File)
				} else {
					stats.Files += 1
					stats.Size += info.Size()
				}

				return err
			})
			if err != nil {
				return err
			}

			if dryRun {
				fmt.Println(stats)
				return nil
			}

			for _, p := range stats.RemovePaths {
				fmt.Println("Removing ", p)
				err = os.RemoveAll(p)
				if os.IsPermission(err) {
					fmt.Printf("err: %s. skipping\n", err)
					continue // skip files that can't be deleted due to permission errors
				}
			}

			fmt.Println(stats)
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
