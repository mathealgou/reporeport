package projectCharacteristics

import (
	"bufio"
	"fmt"
	"os"
	"reporeport/utils/configs"
	"reporeport/utils/count"
	"reporeport/utils/output/colors"
	"strconv"
	"strings"
	"time"
)

func ProjectCharacteristics(files []string) []string {
	var projectChars []string
	if hasDockerfile(files) {
		projectChars = append(projectChars,
			"This project contains "+colors.ColorEscapeSequencesMap["cyan"]+"Docker configuration files."+colors.ColorReset)
	} else {
		projectChars = append(projectChars,
			"This project does "+colors.ColorEscapeSequencesMap["cyan"]+"not appear to be Docker compatible."+colors.ColorReset)
	}

	if hasDocumentationFiles(files) {
		projectChars = append(projectChars,
			"This project has a "+(colors.ColorEscapeSequencesMap["cyan"]+"significant amount of documentation."+colors.ColorReset))
	} else {
		projectChars = append(projectChars, "This project seems to"+colors.ColorEscapeSequencesMap["cyan"]+" lack sufficient documentation, you're on your own, buddy."+colors.ColorReset)
	}

	if configs.VerboseFlag {
		fmt.Println("Attempting to retrieve the date of the first commit...")
	}
	firstCommitDate := tryGetFirstCommitDate()
	if firstCommitDate != "Unknown" {
		projectChars = append(projectChars, "The first commit in this repository was made on "+(colors.ColorEscapeSequencesMap["cyan"]+firstCommitDate+colors.ColorReset+"."))
	} else {
		projectChars = append(projectChars, "Could not determine the date of the first commit.")
	}
	lastCommitDate := tryGetLastCommitDate()
	if lastCommitDate != "Unknown" {
		projectChars = append(projectChars, "The last commit in this repository was made on "+(colors.ColorEscapeSequencesMap["cyan"]+lastCommitDate+colors.ColorReset+"."))
	} else {
		projectChars = append(projectChars, "Could not determine the date of the last commit.")
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
	hasSignificantDocLines := docLines >= 100

	hasDocs := hasSignificantDocFiles && hasSignificantDocLines

	return hasDocs
}

func tryGetFirstCommitDate() string {
	headFilePath := ".git/logs/HEAD"

	file, err := os.Open(headFilePath)
	if err != nil {
		return "Unknown"
	}
	defer file.Close()

	var firstCommitDate string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		if len(parts) >= 3 {
			timestampStr := parts[4]
			timestampInt, _ := strconv.ParseInt(timestampStr, 10, 64)

			commitTime := time.Unix(timestampInt, 0)

			firstCommitDate = commitTime.Format("02 Jan 2006")

			break

		}
	}

	if firstCommitDate == "" {
		return "Unknown"
	}

	return firstCommitDate

}

func tryGetLastCommitDate() string {
	headFilePath := ".git/logs/HEAD"

	file, err := os.Open(headFilePath)
	if err != nil {
		return "Unknown"
	}
	defer file.Close()

	var lastCommitDate string
	scanner := bufio.NewScanner(file)
	var lastLine string
	// There has to be a better way to get the last line, right?
	for scanner.Scan() {
		lastLine = scanner.Text()
	}

	parts := strings.Split(lastLine, " ")
	if len(parts) >= 3 {
		timestampStr := parts[4]
		timestampInt, _ := strconv.ParseInt(timestampStr, 10, 64)

		commitTime := time.Unix(timestampInt, 0)
		lastCommitDate = commitTime.Format("02 Jan 2006")
	}

	if lastCommitDate == "" {
		return "Unknown"
	}

	return lastCommitDate
}
