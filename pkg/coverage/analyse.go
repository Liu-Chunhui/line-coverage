package coverage

import (
	"fmt"
)

// ConvertProfileToBranch adjusts starting and finishing line number from coverage profile,
// which is to create Branch for calculation later.
func ConvertProfileToBranch(profile *CoverageProfile, filelines []string) (*Branch, error) {
	startLine := profile.StartLine
	finishLine := profile.FinishLine

	start := filelines[startLine-1]
	// when start position is the end of this line, then skip this line
	if start[profile.StartPosition] == '\n' {
		startLine += 1
	}

	finish := filelines[finishLine-1]
	// when ending position is the end of the line, then set finish line to previous line
	if finish[profile.FinishPosition-1] == '\n' {
		finishLine -= 1
	}

	return &Branch{
		Start:   startLine,
		Finish:  finishLine,
		Covered: profile.Statements > 0 && profile.Executions > 0,
	}, nil
}

// AnalyseTarget
func AnalyseTargetFile(target string, branches []*Branch) (*Result, error) {
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
