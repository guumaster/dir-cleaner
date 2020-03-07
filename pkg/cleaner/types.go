package cleaner

import (
	"fmt"
	"os"

	"github.com/dustin/go-humanize"
)

// Options contains all the app possible options.
type Options struct {

	// Path start point from where to start a search.
	Path string

	// Pattern base pattern to scan for.
	Patterns []string

	// DryRun indicates if the app should remove files or just show info.
	DryRun bool

	// CountBytes indicates if the app should count exact bytes or count 4K blocks (to match with 'du' defaults)
	CountBytes bool

	// Verbose indicates if the app should print all match and paths to be removed.
	Verbose bool

	// MaxDepth indicates how deep the app should go to scan for the given pattern.
	MaxDepth int
}

// FileInfo extends FileInfo with specific data about relative path and depth level
type FileInfo struct {
	os.FileInfo

	// Filename of the given file (empty if it is a folder)
	Filename string

	// RelativePath partial path to the given path
	RelativePath string

	// Level relative depth level of the given path
	Level int
}

// Stats contains aggregated information of the scanned path
type Stats struct {

	// Path is the path where the scan started
	Path string

	// RemovePaths is a list of all paths that matched the pattern
	RemovePaths []string

	// Files is a counter of the total files inside folders that matched the pattern
	Files int

	// Files is a counter of the total files inside folders that matched the pattern
	FilesMatched int

	// Size is the total size of the files that matched the pattern
	Size int64
}

// String shows a formatted report of all the stats
func (s Stats) String() string {
	if s.FilesMatched == 0 && len(s.RemovePaths) == 0 {
		return fmt.Sprintf("No match found on [%s]\n", s.Path)
	}

	return fmt.Sprintf("Path: [%s]: Scanned %s files. Matched %s in %s directories. [%s]",
		s.Path,
		humanize.FormatInteger("#,###.", s.Files),
		humanize.FormatInteger("#,###.", s.FilesMatched),
		humanize.FormatInteger("#,###.", len(s.RemovePaths)),
		humanize.Bytes(uint64(s.Size)),
	)
}
