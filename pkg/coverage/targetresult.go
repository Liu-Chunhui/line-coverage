package coverage

import (
	"fmt"
)

// target: github.com/yesino/line-coverage/test/testdata.go
// branches: [x]*branch from target
func AnalyseTargetFile(target string, branches []*branch) (*Result, error) {
	if branches == nil {
		return nil, fmt.Errorf("target contains empty branches. Target: %s", target)
	}

	coveredLines := 0
	uncoveredLines := 0

	for _, b := range branches {
		if b.Finish < b.Start {
			return nil, fmt.Errorf("branch starting line is behind ending line. Target file: %s, branch: %+v", target, b)
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
