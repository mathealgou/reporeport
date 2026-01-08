package utils

import (
	"reporeport/utils/count"
	"reporeport/utils/fileSystem"
	"reporeport/utils/projectCharacteristics"
	"reporeport/utils/projectType"
	"reporeport/utils/types"
)

func GenerateReport(includeLibs bool, useGitignore bool) types.Report {

	files := fileSystem.WalkDirectory()

	filteredFiles := FilterStringSlice(files, func(x any) bool {
		if IsToBeCounted(x.(string), includeLibs) {
			return true
		}
		return false

	})

	if useGitignore {
		filteredFiles = fileSystem.ApplyGitignoreFilter(filteredFiles)
	}

	totalByType := count.CountFilesByExtension(filteredFiles)

	linesByExtension := count.CountLinesByExtension(filteredFiles)

	percentageByType := count.CountFileExtensionPercentage(totalByType)

	percentageLinesByType := count.CountLinePercentageByExtension(linesByExtension)

	projectTypeName := projectType.InferProjectType(files, percentageLinesByType)

	projectCharacteristics := projectCharacteristics.ProjectCharacteristics(files)

	result := types.Report{
		TotalFiles:             len(filteredFiles),
		TotalByType:            totalByType,
		PercentageByType:       percentageByType,
		TotalLinesByType:       linesByExtension,
		PercentageLinesByType:  percentageLinesByType,
		ProjectType:            projectTypeName,
		ProjectCharacteristics: projectCharacteristics,
	}
	return result
}
