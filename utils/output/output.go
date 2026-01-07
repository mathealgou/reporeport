package output

import (
	"encoding/json"
	"fmt"
	"os"
	"reporeport/utils/output/colors"
	"reporeport/utils/types"
	"time"
)

func PrintReport(report types.Report) error {
	JSONReport, _ := json.MarshalIndent(report, "", "  ")
	fmt.Printf("%s\n", JSONReport)
	return nil
}

func PrintErrorAndExit(err error) {
	fmt.Println("Error:", err)
	os.Exit(1)
}

func PrintLinesByPercentage(report types.Report) {
	fmt.Println("Percentage of Lines by " + colors.ColorEscapeSequencesMap["cyan"] + "File Type:" + colors.ColorReset)
	percentageLinesOutput := FormatPercentageLinesByType(report.PercentageLinesByType)
	fmt.Println(percentageLinesOutput)
}

func PrintElapsedTime(elapsedTime time.Duration) {
	message := "Report generated in: " + colors.ColorEscapeSequencesMap["cyan"] + elapsedTime.String() + colors.ColorReset
	fmt.Println(message)
}

func PrintProjectType(report types.Report) {
	message := "Inferred Project Type: " + colors.ColorEscapeSequencesMap["cyan"] + report.ProjectType + colors.ColorReset
	fmt.Println(message)
}
