package coverage

import (
	"github.com/Liu-Chunhui/line-coverage/pkg/fileparser"
)

func Calculate(profileFilename string, module string, basePath string) ([]*Result, error) {
	// scan coverage.out to build []*coverageProfile
	profiles, err := loadProfiles(profileFilename, module, basePath)
	if err != nil {
		return nil, err
	}

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

	// code in lines
	targetBranches := make(map[string][]*branch)

	// loop map
	for filename, coverageProfiles := range maps {
		codeInLines, err := fileparser.ReadLines(filename)
		if err != nil {
			return nil, err
		}

		for _, p := range coverageProfiles {
			target, b := convertProfileToBranch(p, codeInLines)
			if values, ok := targetBranches[target]; ok {
				targetBranches[target] = append(values, b)
			} else {
				targetBranches[target] = []*branch{b}
			}
		}
	}

	var results []*Result

	for target, branches := range targetBranches {
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
