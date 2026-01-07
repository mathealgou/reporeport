package fileSystem

import (
	"io/fs"
	"os"
	"reporeport/utils/output"
)

func WalkDirectory() []string {
	dir, dirError := os.Getwd()
	if dirError != nil {
		output.PrintErrorAndExit(dirError)
	}

	var files []string
	walkDirError := fs.WalkDir(os.DirFS(dir), ".", func(path string, d fs.DirEntry, err error) error {

		files = append(files, path)
		return nil
	})

	if walkDirError != nil {
		output.PrintErrorAndExit(walkDirError)
	}
	return files
}
