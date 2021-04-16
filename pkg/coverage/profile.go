package coverage

import (
	"fmt"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/Liu-Chunhui/line-coverage/pkg/fileparser"
)

// loadProfiles maps coverage file to coverageProfile objects
func loadProfiles(coverageFile string, module string, base string) ([]*coverageProfile, error) {
	lines, err := fileparser.ReadLines(coverageFile, fileparser.CoverageProfileExcludingRules...)
	if err != nil {
		return nil, err
	}

	var profiles []*coverageProfile

	for _, line := range lines {
		p, err := mapLineToCoverageProfile(line, module, base)
		if err != nil {
			return nil, err
		}
		profiles = append(profiles, p)
	}

	return profiles, nil
}

// line: github.com/yesino/line-coverage/test/testdata.go:30.50,32.9 2 4
// module: github.com/yesino
// base: .
func mapLineToCoverageProfile(line string, module string, base string) (*coverageProfile, error) {
	log.Debug(fmt.Sprintf("line: %s", line))
	topLvlParts := strings.Split(strings.TrimSpace(line), " ") // github.com/yesino/line-coverage/test/testdata.go:30.50,32.9 2 4
	statements, err := strconv.Atoi(topLvlParts[1])
	if err != nil {
		return nil, err
	}
	execution, err := strconv.Atoi(topLvlParts[2])
	if err != nil {
		return nil, err
	}
	secondLvlParts := strings.Split(topLvlParts[0], ":")   // github.com/yesino/line-coverage/test/testdata.go 30.50,32.9
	thirdLvlParts := strings.Split(secondLvlParts[1], ",") // 30.50 32.9
	startParts := strings.Split(thirdLvlParts[0], ".")     // 30 50
	finishParts := strings.Split(thirdLvlParts[1], ".")    // 32 9
	startLine, err := strconv.Atoi(startParts[0])          // 30
	if err != nil {
		return nil, err
	}
	startPosition, err := strconv.Atoi(startParts[1]) // 50
	if err != nil {
		return nil, err
	}
	finishLine, err := strconv.Atoi(finishParts[0]) // 32
	if err != nil {
		return nil, err
	}
	finishPosition, err := strconv.Atoi(finishParts[1]) // 9
	if err != nil {
		return nil, err
	}

	filename := strings.ReplaceAll(secondLvlParts[0], module, base)

	return &coverageProfile{
		Target:         secondLvlParts[0],
		TargetFilename: filename,
		StartLine:      startLine,
		StartPosition:  startPosition,
		FinishLine:     finishLine,
		FinishPosition: finishPosition,
		Statements:     statements,
		Executions:     execution,
	}, nil
}
