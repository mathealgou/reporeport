package fileSystem

import (
	"bufio"
	"os"
	"strings"
)

func SearchForWordInFile(filePath string, word string) bool {
	file, err := os.Open(filePath)
	if err != nil {
		return false
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), word) {
			closingErr := file.Close()
			if closingErr != nil {
				return false
			}
			return true
		}
	}
	closingErr := file.Close()
	if closingErr != nil {
		return false
	}
	return false
}
