package coverage

import (
	"fmt"
	"strconv"
	"strings"
)

// ConvertProfileToBranch splits profileLine into parts and use it to extract Branch data from fine lines
// profileLine: github.com/yesino/pkg/identity/id.go:30.50,32.9 2 4
func ConvertProfileToBranch(profileLine string, filelines []string) (*Branch, error) {
	topLvlParts := strings.Split(profileLine, " ") // github.com/yesino/line-coverage/pkg/identity/id.go:30.50,32.9 2 4
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

	start := filelines[startLine-1]
	// when start position is the end of this line, then skip this line
	if start[startPosition] == '\n' {
		startLine += 1
	}

	finish := filelines[finishLine-1]
	// check if finish line is nothing but "}", then skip this line
	if finish[finishPosition-1] == '\n' {
		finishLine -= 1
	}

	return &Branch{
		Target:  secondLvlParts[0],
		Start:   startLine,
		Finish:  finishLine,
		Covered: statements > 0 && execution > 0,
	}, nil
}

func Analyse(target string, branches []*Branch) (*Result, error) {
	if branches == nil {
		return nil, fmt.Errorf("target contains empty branches. Target: %s", target)
	}

	coveredLines := 0
	uncoveredLines := 0

	for _, b := range branches {
		if b.Finish < b.Start {
			return nil, fmt.Errorf("branch starting line is behind ending line. Target file: %s, Branch: %+v", target, b)
		}

		lines := b.Finish - b.Start + 1

		if b.Covered {
			coveredLines += lines
		} else {
			uncoveredLines += lines
		}
	}

	return &Result{
		target,
		coveredLines,
		uncoveredLines,
	}, nil
}
