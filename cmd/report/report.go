package report

import (
	"fmt"

	"github.com/Liu-Chunhui/line-coverage/pkg/coverage"
	"github.com/Liu-Chunhui/line-coverage/pkg/percentage"
)

func Report(coverprofile string, module string, basePath string) error {
	results, err := coverage.Calculate(coverprofile, module, basePath)
	if err != nil {
		return err
	}

	for _, r := range results {
		total := r.CoveredLines + r.UncoveredLines
		cover := float64(r.CoveredLines) / float64(total)

		fmt.Printf("%s: %s(TotalCovered: %d, TotalUncovered: %d)\n", r.Target, percentage.Display(cover), r.CoveredLines, r.UncoveredLines)
	}

	overall := coverage.CalculateOverall(results)

	fmt.Printf("Overall: %s\n", percentage.Display(overall))

	return nil
}
