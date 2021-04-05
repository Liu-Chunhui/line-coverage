package report

import (
	"fmt"

	"github.com/Liu-Chunhui/line-coverage/pkg/coverage"
	"github.com/Liu-Chunhui/line-coverage/pkg/percentage"
)

func Report(coverprofile string, module string, basePath string) error {
	results, err := coverage.Calculate(coverprofile, module, basePath)
	if err != nil {
		return nil
	}

	for _, r := range results {
		total := r.CoveredLines + r.UncoveredLines
		cover := float64(r.CoveredLines) / float64(total)

		fmt.Printf("%s: %s\n", r.Target, percentage.Display(cover))
	}

	overall := coverage.CalculateOverall(results)

	fmt.Printf("Overall: %s\n", percentage.Display(overall))

	return nil
}
