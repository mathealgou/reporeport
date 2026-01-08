package main

import (
	"fmt"
	"os"
	"reporeport/utils"
	"reporeport/utils/output"
	"reporeport/utils/output/help"
	"time"
)

func main() {
	start := time.Now()
	args := os.Args[1:]
	for i, arg := range args {
		fmt.Printf("arg %d: %s\n", i+1, arg)
	}
	anyArgs := make([]any, len(args))

	for i, a := range args {
		anyArgs[i] = a
	}

	if utils.FindInSlice(anyArgs, func(x any) bool {
		s, _ := x.(string)
		return s == "--help"
	}) != nil {
		help.PrintHelp()
	}

	includeLibs := utils.FindInSlice(anyArgs, func(x any) bool {
		s, _ := x.(string)
		return s == "--include-libs"
	}) != nil

	report := utils.GenerateReport(includeLibs)

	output.PrintProjectType(report)

	output.PrintProjectCharacteristics(report.ProjectCharacteristics)

	output.PrintLinesByPercentage(report)

	// Just to show off, really
	elapsed := time.Since(start)
	output.PrintElapsedTime(elapsed)

}
