package utils

import (
	"reporeport/utils/count"
	"reporeport/utils/fileSystem"
	"reporeport/utils/projectType"
	"reporeport/utils/types"
)

func GenerateReport(includeLibs bool) types.Report {

	files := fileSystem.WalkDirectory()

	filteredFiles := FilterStringSlice(files, func(x any) bool {
		if IsToBeCounted(x.(string), includeLibs) {
			return true
		}
		return false

	})

	totalByType := count.CountFilesByExtension(filteredFiles)

	linesByExtension := count.CountLinesByExtension(filteredFiles)

	percentageByType := count.CountFileExtensionPercentage(totalByType)

	percentageLinesByType := count.CountLinePercentageByExtension(linesByExtension)

	projectTypeName := projectType.InferProjectType(files, percentageLinesByType)

	result := types.Report{
		TotalFiles:            len(filteredFiles),
		TotalByType:           totalByType,
		PercentageByType:      percentageByType,
		TotalLinesByType:      linesByExtension,
		PercentageLinesByType: percentageLinesByType,
		ProjectType:           projectTypeName,
	}
	return result
}
