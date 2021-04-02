package coverage

func calculateOverall(results []*Result) float64 {
	totalCoveredLines := 0
	totalUncoveredLines := 0

	for _, r := range results {
		totalCoveredLines += r.CoveredLines
		totalUncoveredLines += r.UncoveredLines
	}

	return float64(totalCoveredLines) / float64(totalCoveredLines+totalUncoveredLines)
}
