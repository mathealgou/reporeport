package utils

import (
	"fmt"
	"reporeport/utils/configs"
	"reporeport/utils/count"
	"reporeport/utils/fileSystem"
	"reporeport/utils/projectCharacteristics"
	"reporeport/utils/projectType"
	"reporeport/utils/types"
)

func GenerateReport(includeLibs bool, useGitignore bool) types.Report {

	if configs.VerboseFlag {
		fmt.Println("Generating report with the following settings:")
		fmt.Printf("  Include library files: %v\n", includeLibs)
		fmt.Printf("  Use .gitignore: %v\n", useGitignore)
	}

	if configs.VerboseFlag {
		fmt.Println("Walking through the directory to collect files...")
	}

	files := fileSystem.WalkDirectory()

	if configs.VerboseFlag {
		fmt.Printf("Total files found: %d\n", len(files))
		fmt.Println("Filtering files based on the provided settings...")
	}

	filteredFiles := FilterStringSlice(files, func(x any) bool {
		if IsToBeCounted(x.(string), includeLibs) {
			return true
		}
		return false

	})

	if useGitignore {
		if configs.VerboseFlag {
			fmt.Println("Applying .gitignore filters...")
		}
		filteredFiles = fileSystem.ApplyGitignoreFilter(filteredFiles)
	}

	if configs.VerboseFlag {
		fmt.Printf("Total files after filtering: %d\n", len(filteredFiles))
		fmt.Println("Generating full report...")
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

	if configs.VerboseFlag {
		fmt.Println("Report generation completed.")
	}
	return result
}
