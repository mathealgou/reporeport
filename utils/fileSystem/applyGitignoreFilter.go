package fileSystem

import (
	"os"
	"path/filepath"

	ignore "github.com/sabhiram/go-gitignore"
)

func ApplyGitignoreFilter(files []string) []string {
	gitignorePath := ".gitignore"
	ignore, err := ignore.CompileIgnoreFile(gitignorePath)
	if err != nil {
		// If there's no .gitignore file, return the original list
		if os.IsNotExist(err) {
			return files
		}
		// For other errors, you might want to handle them differently
		return files
	}

	var filteredFiles []string
	for _, file := range files {
		relPath, err := filepath.Rel(".", file)
		if err != nil {
			relPath = file // Fallback to original path if there's an error
		}
		if !ignore.MatchesPath(relPath) && !isParentIgnored(relPath, ignore) {
			filteredFiles = append(filteredFiles, file)
		}
	}
	return filteredFiles
}

func isParentIgnored(path string, ignore *ignore.GitIgnore) bool {
	dir := filepath.Dir(path)
	for dir != "." && dir != string(os.PathSeparator) {
		if ignore.MatchesPath(dir + string(os.PathSeparator)) {
			return true
		}
		dir = filepath.Dir(dir)
	}
	return false
}
