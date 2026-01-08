package output

import (
	"fmt"
	"reporeport/utils/configs"
	"reporeport/utils/output/colors"
	"strings"
)

func sortPercentageLinesByType(percentageLinesByType map[string]float64) []string {
	sortedKeys := make([]string, 0, len(percentageLinesByType))
	for key := range percentageLinesByType {
		sortedKeys = append(sortedKeys, key)
	}
	for i := 0; i < len(sortedKeys)-1; i++ {
		for j := i + 1; j < len(sortedKeys); j++ {
			if percentageLinesByType[sortedKeys[i]] < percentageLinesByType[sortedKeys[j]] {
				sortedKeys[i], sortedKeys[j] = sortedKeys[j], sortedKeys[i]
			}
		}
	}
	return sortedKeys
}

func FormatPercentageLinesByType(percentageLinesByType map[string]float64) string {
	sortedKeys := sortPercentageLinesByType(percentageLinesByType)
	// round down
	width := configs.GetTerminalWidth() / 2
	var lines []string
	var biggestPercentExt string
	for ext, percent := range percentageLinesByType {
		if percent > percentageLinesByType[biggestPercentExt] {
			biggestPercentExt = ext
		}
	}
	normalizedWidthMultiplier := float64(width) / percentageLinesByType[biggestPercentExt]
	for i, key := range sortedKeys {
		color := colors.ColorEscapeSequences[i%len(colors.ColorEscapeSequences)]

		barLength := int(percentageLinesByType[key] * normalizedWidthMultiplier)
		ext := key
		percent := percentageLinesByType[ext]
		bar := strings.Repeat("█", barLength)
		line := fmt.Sprintf("%-10s | %-5.2f%% | %s", ext, percent, bar)
		line = color + line
		lines = append(lines, line)
	}
	lines = append(lines, "\033[0m")
	return strings.Join(lines, "\n")
}
