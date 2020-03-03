package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func walkFn(rootPath string, maxDepth int, stats *Stats) error {
	return filepath.Walk(rootPath, func(srcPath string, info os.FileInfo, err error) error {
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
			fmt.Printf("[lvl %2d] path [%s] \n", meta.Level, meta.Dir+meta.File)
		} else {
			stats.Files += 1
			stats.Size += info.Size()
		}

		return err
	})
}

func cleanPath(rootPath string, dryRun bool, maxDepth int) (*Stats, error) {
	var size int64
	var files int
	var removePaths []string

	stats := &Stats{rootPath, removePaths, files, size}

	err := walkFn(rootPath, maxDepth, stats)

	if dryRun {
		return stats, err
	}

	for _, p := range stats.RemovePaths {
		fmt.Println("Removing ", p)
		err = os.RemoveAll(p)
		if os.IsPermission(err) {
			fmt.Printf("err: %s. skipping\n", err)
			continue // skip files that can't be deleted due to permission errors
		}
	}

	return stats, err
}
