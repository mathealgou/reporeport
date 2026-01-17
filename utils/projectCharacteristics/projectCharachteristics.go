package projectCharacteristics

import (
	"fmt"
	"os/exec"
	"reporeport/utils/configs"
	"reporeport/utils/count"
	"reporeport/utils/output/colors"
	"strings"
)

func ProjectCharacteristics(files []string) []string {
	var projectChars []string

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

	if hasDockerfile(files) {
		projectChars = append(projectChars,
			"This project contains "+colors.ColorEscapeSequencesMap["cyan"]+"Docker configuration files."+colors.ColorReset)
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
	command := "git log --reverse --pretty=format:%cd --date=iso-local | head -n 1"
	output, err := exec.Command("bash", "-c", command).Output()
	if err != nil {
		return "Unknown"
	}

	firstCommitDate := strings.TrimSpace(string(output))
	if firstCommitDate == "" {
		return "Unknown"
	}

	return formatGitIsoDate(firstCommitDate)

}

func tryGetLastCommitDate() string {
	command := "git log -1 --pretty=format:%cd --date=iso-local"
	output, err := exec.Command("bash", "-c", command).Output()
	if err != nil {
		return "Unknown"
	}

	lastCommitDate := strings.TrimSpace(string(output))
	if lastCommitDate == "" {
		return "Unknown"
	}

	return formatGitIsoDate(lastCommitDate)
}

func formatGitIsoDate(dateString string) string {
	splitDatestring := strings.Split(dateString, " ")
	date := splitDatestring[0]
	time := splitDatestring[1]

	yearMonthDay := strings.Split(date, "-")
	year := yearMonthDay[0]
	month := yearMonthDay[1]
	day := yearMonthDay[2]

	result := fmt.Sprintf("%s/%s/%s @ %s", day, month, year, time)

	return result
}
