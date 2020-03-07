// Package cleaner contains the app with methods to scan and clean folders.
package cleaner

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// App contains the methods to scan and clean paths.
type App struct{}

// Clean scans the given path and removes the files and folders
// that match the given pattern.
// if options.DryRun is set to true, it will only show what it would delete
// without actually deleting anything.
func (a *App) Clean(options *Options) (*Stats, error) {
	stats := &Stats{
		Path:        options.Path,
		RemovePaths: []string{},
		Files:       0,
		Size:        0,
	}

	err := ScanPath(options, stats)
	if err != nil {
		return stats, err
	}

	err = CleanPath(options, stats)

	return stats, err
}

func CleanPath(options *Options, stats *Stats) error {
	for _, p := range stats.RemovePaths {
		if options.Verbose {
			fmt.Printf("Removing path: [%s]\n", p)
		}
		if options.DryRun {
			continue
		}
		err := os.RemoveAll(p)
		if os.IsPermission(err) {
			if options.Verbose {
				fmt.Printf("permission error removing path: \"%s\". skipping...\n", p)
			}
			continue // skip files that can't be deleted due to permission errors
		}
		if err != nil {
			return err
		}
	}

	return nil
}

// AppendFileInfo adds more information to os.FileInfo
func AppendFileInfo(f os.FileInfo, root, path string) *FileInfo {
	dirPath := strings.ReplaceAll(path, root+"/", "")
	file := ""
	if !f.IsDir() {
		file = filepath.Base(dirPath)
		dirPath = filepath.Dir(dirPath)
	}
	return &FileInfo{
		RelativePath: dirPath,
		Filename:     file,
		Level:        len(strings.Split(dirPath, "/")),
	}
}

func ScanPath(options *Options, stats *Stats) error {
	rootPath := options.Path
	maxDepth := options.MaxDepth
	reMatch := regexp.MustCompile(options.Pattern)
	reEndMatch := regexp.MustCompile(options.Pattern + "$")

	return filepath.Walk(rootPath, func(srcPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if srcPath == rootPath {
			return nil
		}

		meta := AppendFileInfo(info, rootPath, srcPath)

		depth := maxDepth
		if !info.IsDir() {
			depth++
		}

		if maxDepth != 0 && meta.Level > depth {
			return nil
		}

		match := reMatch.MatchString(meta.RelativePath)
		if !match {
			return nil
		}

		if info.IsDir() && reEndMatch.MatchString(meta.RelativePath) {
			stats.RemovePaths = append(stats.RemovePaths, srcPath)
		}

		if !info.IsDir() {
			stats.Files += 1
		}
		size := info.Size()
		if !options.CountBytes && size < 4096 {
			size = 4096
		}
		stats.Size += size

		return err
	})
}
