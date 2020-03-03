package main

import (
	"fmt"

	"github.com/dustin/go-humanize"
)

type Stats struct {
	Path        string
	RemovePaths []string
	Files       int
	Size        int64
}

func (s Stats) String() string {
	if s.Files == 0 {
		return fmt.Sprintf("No match found on [%s]\n", s.Path)
	}

	return fmt.Sprintf("Path: [%s]: Found %s files [%s]\n",
		s.Path,
		humanize.FormatInteger("#,###.", s.Files),
		humanize.Bytes(uint64(s.Size)),
	)
}
