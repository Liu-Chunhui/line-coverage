package coverage

import (
	"fmt"

	"github.com/Liu-Chunhui/line-coverage/pkg/fileparser"
	"github.com/davecgh/go-spew/spew"
	log "github.com/sirupsen/logrus"
)

func Calculate(profileFilename string, module string, basePath string) ([]*Result, error) {
	log.Info(fmt.Sprintf("Processing profile file: %s", profileFilename))
	// scan coverage.out to build []*coverageProfile
	profiles, err := loadProfiles(profileFilename, module, basePath)
	if err != nil {
		return nil, err
	}

	log.Info("Building profile targets map")
	// map key: code filename
	maps := make(map[string][]*coverageProfile)

	// loop profiles to build map[string]*coverageProfile
	for _, profile := range profiles {
		if values, ok := maps[profile.TargetFilename]; ok {
			maps[profile.TargetFilename] = append(values, profile)
		} else {
			maps[profile.TargetFilename] = []*coverageProfile{profile}
		}
	}

	log.Info("Building target branches map")
	// code in lines
	targetBranches := make(map[string][]*branch)

	// loop map
	for filename, coverageProfiles := range maps {
		codeInLines, err := fileparser.ReadLines(filename)

		if err != nil {
			return nil, err
		}

		for _, p := range coverageProfiles {
			_, _ = spew.Println("coverageProfile: %v", p)
			target, branches := convertProfileToBranch(p, codeInLines)
			if values, ok := targetBranches[target]; ok {
				targetBranches[target] = append(values, branches...)
			} else {
				targetBranches[target] = branches
			}
		}
	}

	log.Debug("Building results")
	var results []*Result

	for target, branches := range targetBranches {
		_, _ = spew.Println("targetBranches: %s, %v\n", target, branches)
		r, err := calculateTargetResult(target, branches)
		if err != nil {
			return nil, err
		}

		results = append(results, r)
	}

	return results, nil
}

func CalculateOverall(results []*Result) float64 {
	totalCoveredLines := 0
	totalUncoveredLines := 0

	for _, r := range results {
		totalCoveredLines += r.CoveredLines
		totalUncoveredLines += r.UncoveredLines
	}

	return float64(totalCoveredLines) / float64(totalCoveredLines+totalUncoveredLines)
}
