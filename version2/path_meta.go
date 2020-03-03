package main

import (
	"path/filepath"
	"strings"
)

type PathMeta struct {
	Dir   string
	File  string
	Level int
}

func getMeta(root, path string, isDir bool) *PathMeta {
	dirPath := strings.ReplaceAll(path, root+"/", "")
	file := ""
	if !isDir {
		file = filepath.Base(dirPath)
		dirPath = filepath.Dir(dirPath)
	}
	return &PathMeta{
		Dir:   dirPath,
		File:  file,
		Level: len(strings.Split(dirPath, "/")),
	}
}
