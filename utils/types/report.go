package types

type Report struct {
	TotalFiles            int
	TotalByType           map[string]int
	PercentageByType      map[string]float64
	TotalLinesByType      map[string]int
	PercentageLinesByType map[string]float64
	ProjectType           string
}
