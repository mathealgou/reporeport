package main

import (
	"os"
	"reporeport/utils"
	"reporeport/utils/configs"
	"reporeport/utils/output"
	"reporeport/utils/output/help"
	"time"
)

func main() {
	start := time.Now()
	args := os.Args[1:]

	anyArgs := make([]any, len(args))

	for i, a := range args {
		anyArgs[i] = a
	}

	if utils.FindInSlice(anyArgs, func(x any) bool {
		s, _ := x.(string)
		return s == "--help"
	}) != nil {
		help.PrintHelp()
		return
	}

	includeLibs := utils.FindInSlice(anyArgs, func(x any) bool {
		s, _ := x.(string)
		return s == "--include-libs"
	}) != nil

	useGitignore := utils.FindInSlice(anyArgs, func(x any) bool {
		s, _ := x.(string)
		return s == "--use-gitignore"
	}) != nil

	configs.VerboseFlag = utils.FindInSlice(anyArgs, func(x any) bool {
		s, _ := x.(string)
		return s == "--verbose"
	}) != nil
	// Just to show off, really

	report := utils.GenerateReport(includeLibs, useGitignore)

	elapsed := time.Since(start)

	waitTime := 500 * time.Millisecond

	output.PrintProjectType(report)

	time.Sleep(waitTime)

	output.PrintProjectCharacteristics(report.ProjectCharacteristics, waitTime)

	time.Sleep(waitTime)

	output.PrintLinesByPercentage(report, waitTime)

	time.Sleep(waitTime)

	output.PrintTotalLinesOfCode(report)

	output.PrintElapsedTime(elapsed)

}
