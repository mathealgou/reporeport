package projectCharacteristics

import (
	"reporeport/utils/count"
	"strings"
)

func ProjectCharacteristics(files []string) []string {
	var projectChars []string
	if hasDockerfile(files) {
		projectChars = append(projectChars, "This project contains Docker configuration files.")
	} else {
		projectChars = append(projectChars, "This project does not appear to be Docker compatible.")
	}

	if hasDocumentationFiles(files) {
		projectChars = append(projectChars, "This project has a significant amount of documentation.")
	} else {
		projectChars = append(projectChars, "This project seems to lack sufficient documentation, you're on your own, buddy.")
	}

	return projectChars
}

func hasDockerfile(files []string) bool {
	for _, file := range files {
		if file == "Dockerfile" || file == "docker-compose.yml" {
			return true
		}
	}
	return false
}

func hasDocumentationFiles(files []string) bool {
	docCount := 0
	docLines := 0
	docFileExtensions := []string{".md", ".MD", "LICENCE", ".rst", ".txt"}

	for _, file := range files {
		for _, ext := range docFileExtensions {
			if strings.HasSuffix(file, ext) {
				docCount++
				lines := count.CountLines(file)
				docLines += lines
				break
			}
		}
	}

	hasSignificantDocFiles := docCount > 0
	hasSignificantDocLines := docLines >= 1000

	hasDocs := hasSignificantDocFiles && hasSignificantDocLines

	return hasDocs
}
