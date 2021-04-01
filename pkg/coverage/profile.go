package coverage

import (
	"strconv"
	"strings"
)

// line: github.com/yesino/example/test/testdata.go:30.50,32.9 2 4
// module: github.com/yesino
// base: .
func ConvertToCoverageProfile(line string, module string, base string) (*CoverageProfile, error) {
	topLvlParts := strings.Split(line, " ") // github.com/yesino/line-coverage/pkg/identity/id.go:30.50,32.9 2 4
	statements, err := strconv.Atoi(topLvlParts[1])
	if err != nil {
		return nil, err
	}
	execution, err := strconv.Atoi(topLvlParts[2])
	if err != nil {
		return nil, err
	}
	secondLvlParts := strings.Split(topLvlParts[0], ":")   // github.com/yesino/line-coverage/pkg/identity/id.go 30.50,32.9
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

	return &CoverageProfile{
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
