package count

import (
	"bufio"
	"os"
	"slices"
	"strings"
)

func CountFilesByExtension(files []string) map[string]int {
	counts := make(map[string]int)
	extensions := getFileExtensions(files)

	for _, ext := range extensions {
		counts[ext] = 0
		for _, file := range files {
			if strings.HasSuffix(file, "."+ext) {
				counts[ext]++
			}
		}
	}
	return counts
}

func getFileExtensions(files []string) []string {
	var extensions []string

	for _, file := range files {
		ext := getFileExtension(file)
		if ext != "" && !slices.Contains(extensions, ext) {
			extensions = append(extensions, ext)
		}
	}
	return extensions
}

func getFileExtension(file string) string {
	parts := strings.Split(file, ".")
	if len(parts) > 1 {
		return parts[len(parts)-1]
	}
	return ""
}

func CountFileExtensionPercentage(counts map[string]int) map[string]float64 {
	total := 0
	percentages := make(map[string]float64)

	for _, count := range counts {
		total += count
	}

	for ext, count := range counts {
		if total > 0 {
			percentages[ext] = (float64(count) / float64(total)) * 100
		} else {
			percentages[ext] = 0
		}
	}
	return percentages
}

func countLines(filename string) int {
	lc := 0
	f, err := os.Open(filename)
	if err != nil {
		return 0
	}
	defer f.Close()
	s := bufio.NewScanner(f) // The split function defaults to ScanLines.
	for s.Scan() {
		lc++
	}
	return lc
}

func CountLinesByExtension(files []string) map[string]int {
	counts := make(map[string]int)
	extensions := getFileExtensions(files)

	for _, ext := range extensions {
		counts[ext] = 0
		for _, file := range files {
			if strings.HasSuffix(file, "."+ext) {
				lineCount := countLines(file)
				counts[ext] += lineCount
			}
		}
	}
	return counts
}

func CountLinePercentageByExtension(lineCounts map[string]int) map[string]float64 {
	total := 0
	percentages := make(map[string]float64)

	for _, count := range lineCounts {
		total += count
	}
	for ext, count := range lineCounts {
		if total > 0 {
			percentages[ext] = (float64(count) / float64(total)) * 100
		} else {
			percentages[ext] = 0
		}
	}
	return percentages
}

func GetTopNPercentages(percentages map[string]float64, n int) map[string]float64 {
	type kv struct {
		Key   string
		Value float64
	}

	var ss []kv
	for k, v := range percentages {
		ss = append(ss, kv{k, v})
	}

	slices.SortFunc(ss, func(a, b kv) int {
		if a.Value < b.Value {
			return 1
		}
		if a.Value > b.Value {
			return -1
		}
		return 0
	})

	topN := make(map[string]float64)
	for i := 0; i < n && i < len(ss); i++ {
		topN[ss[i].Key] = ss[i].Value
	}
	return topN
}
