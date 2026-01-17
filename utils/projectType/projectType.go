package projectType

import (
	"reporeport/utils/count"
	"reporeport/utils/fileSystem"
	"strings"
)

func InferProjectType(files []string, percentagesByType map[string]float64) string {

	if len(files) == 0 {
		return "Empty Repository"
	}

	top1percentage := count.GetTopNPercentages(percentagesByType, 1)

	var projectType string

	// simple detection by language
	switch {
	case isFileTypeInTopN("py", top1percentage):
		return "Python Project"
	case isFileTypeInTopN("java", top1percentage):
		return "Java Project"
	case isFileTypeInTopN("go", top1percentage):
		return "Go Project"
	case isFileTypeInTopN("cpp", top1percentage) || isFileTypeInTopN("hpp", top1percentage):
		return "C++ Project"
	case isFileTypeInTopN("c", top1percentage) || isFileTypeInTopN("h", top1percentage):
		return "C Project"
	}

	top3percentages := count.GetTopNPercentages(percentagesByType, 3)

	if hasPackageJSON(files) && (isFileTypeInTopN("jsx", top3percentages) || isFileTypeInTopN("js", top3percentages)) {
		projectType = "Javascript React Project"
	}
	if hasPackageJSON(files) && isFileTypeInTopN("tsx", top3percentages) || isFileTypeInTopN("ts", top3percentages) {
		projectType = "Typescript React Project"
	}
	if hasPackageJSON(files) && isFileTypeInTopN("ts", top1percentage) {
		projectType = "Typescript backend/script Project"
	}
	if hasPackageJSON(files) && isFileTypeInTopN("js", top1percentage) {
		projectType = "Javascript backend/script Project"
	}
	if isVtexAppOrStore(files) {
		projectType = "VTEX IO App or Storefront"
	}
	if isFileTypeInTopN("html", top3percentages) && !isFileTypeInTopN("js", top3percentages) {
		projectType = "Simple web page"
	}
	if isFileTypeInTopN("css", top1percentage) && !isFileTypeInTopN("html", top3percentages) {
		projectType = "CSS library or framework"
	}

	// fallback
	if projectType == "" {
		projectType = "Unable to determine project type"
	}
	return projectType
}

func hasPackageJSON(files []string) bool {
	for _, file := range files {
		if strings.Contains(file, "package.json") {
			return true
		}
	}
	return false
}

func isFileTypeInTopN(fileType string, topN map[string]float64) bool {
	_, exists := topN[fileType]
	return exists
}

func isVtexAppOrStore(files []string) bool {
	for _, file := range files {
		if strings.Contains(file, "manifest.json") {
			if fileSystem.SearchForWordInFile(file, "vtex") {
				return true
			}

		}
	}
	return false
}
