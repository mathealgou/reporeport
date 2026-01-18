package output

import (
	"encoding/json"
	"fmt"
	"os"
	"reporeport/utils/configs"
	"reporeport/utils/output/colors"
	"reporeport/utils/types"
	"strconv"
	"strings"
	"time"
)

func wrapText(text string, width int) string {
	if len(text) <= width {
		return text
	}

	var wrappedText strings.Builder
	for len(text) > width {
		splitAt := width
		for i := width; i > 0; i-- {
			if text[i] == ' ' {
				splitAt = i
				break
			}
		}
		wrappedText.WriteString(text[:splitAt] + "\n")
		text = text[splitAt+1:]
	}
	wrappedText.WriteString(text)
	return wrappedText.String()
}

func PrintReport(report types.Report) error {
	JSONReport, _ := json.MarshalIndent(report, "", "  ")
	fmt.Printf("%s\n", JSONReport)
	return nil
}

func PrintErrorAndExit(err error) {
	fmt.Println("Error:", err)
	os.Exit(1)
}

func PrintLinesByPercentage(report types.Report, waitTime time.Duration) {
	fmt.Println("Percentage of Lines by " + colors.ColorEscapeSequencesMap["cyan"] + "File Type:" + colors.ColorReset)
	percentageLinesOutput := strings.Split(FormatPercentageLinesByType(report.PercentageLinesByType), "\n")

	for _, line := range percentageLinesOutput {
		fmt.Println(line)
		time.Sleep(waitTime)
	}
}

func PrintElapsedTime(elapsedTime time.Duration) {
	message := "Report generated in: " + colors.ColorEscapeSequencesMap["cyan"] + elapsedTime.String() + colors.ColorReset
	fmt.Println(message)
}

func PrintProjectType(report types.Report) {
	message := "Inferred Project Type: " + colors.ColorEscapeSequencesMap["cyan"] + report.ProjectType + colors.ColorReset
	fmt.Println(message)
}

func PrintProjectCharacteristics(characteristics []string, waitTime time.Duration) {
	terminalWidth := configs.GetTerminalWidth()
	barGraphLegendWidth := 27 // Adjust for bullet points and spacing
	textWidth := terminalWidth/2 + barGraphLegendWidth
	fmt.Println("\nProject Characteristics:")
	for _, char := range characteristics {
		if len(char) > textWidth {
			char = wrapText(char, textWidth)
		}
		// hehehehehe ■
		fmt.Println(colors.ColorEscapeSequencesMap["cyan"] + " ■ " + colors.ColorReset + char + "\n")
		time.Sleep(waitTime)
	}
}

func PrintTotalLinesOfCode(report types.Report) {
	sum := 0

	for _, count := range report.TotalLinesByType {
		sum += count
	}

	fmt.Println(colors.ColorEscapeSequencesMap["cyan"] + " ■ " + strconv.Itoa(sum) + " ■ " + colors.ColorReset + " lines of code in total\n")
}
