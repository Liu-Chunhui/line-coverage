package report

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/Liu-Chunhui/line-coverage/pkg/coverage"
	"github.com/Liu-Chunhui/line-coverage/pkg/fileparser"
	"github.com/Liu-Chunhui/line-coverage/pkg/percentage"
)

func Report(coverprofile string, gomod string) error {
	// load module name from go.mod file
	module, err := loadModule(gomod)
	if err != nil {
		return err
	}
	logrus.Info(fmt.Sprintf("module: %s\n", module))

	// load root path from go.mod file path. normally go.mod is located at the root path
	rootPath := filepath.Dir(gomod)
	logrus.Info(fmt.Sprintf("rootPath: %s\n", rootPath))

	results, err := coverage.Calculate(coverprofile, module, rootPath)
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

func loadModule(gomod string) (string, error) {
	lines, err := fileparser.ReadLines(gomod)
	if err != nil {
		return "", err
	}

	if len(lines) < 1 {
		return "", fmt.Errorf("%s is not a valid go.mod file", gomod)
	}

	// validate the module name
	found, err := fileparser.MatchPattern(lines[0], fileparser.GoModGetModuleNameRule)
	if err != nil {
		return "", err
	}
	if !found {
		return "", fmt.Errorf("failed to retrieve module name from file %s", gomod)
	}
	// "module github.com/Liu-Chunhui/line-coverage\n"
	return strings.TrimLeft(strings.TrimSpace(lines[0]), "module "), nil
}
