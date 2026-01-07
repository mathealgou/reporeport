package projectType

import (
	"reporeport/utils/count"
	"strings"
)

func InferProjectType(files []string, percentagesByType map[string]float64) string {
	// Placeholder logic for inferring project type
	// In a real implementation, this would analyze the files to determine the project type

	top1percentage := count.GetTopNPercentages(percentagesByType, 1)

	switch {
	case isFileTypeInTopN("py", top1percentage):
		return "Python Project"
	case isFileTypeInTopN("java", top1percentage):
		return "Java Project"
	case isFileTypeInTopN("go", top1percentage):
		return "Go Project"
	}

	top3percentages := count.GetTopNPercentages(percentagesByType, 3)

	if len(files) == 0 {
		return "Unknown"
	}
	if hasPackageJSON(files) && isFileTypeInTopN("jsx", top3percentages) || isFileTypeInTopN("tsx", top3percentages) {
		return "React Project"
	}
	if hasPackageJSON(files) && isFileTypeInTopN("ts", top3percentages) {
		return "TypeScript Project"
	}
	return "Generic Project"
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

func hasTsxFiles(files []string) bool {
	for _, file := range files {
		if len(file) > 4 && file[len(file)-4:] == ".tsx" {
			return true
		}
	}
	return false
}
