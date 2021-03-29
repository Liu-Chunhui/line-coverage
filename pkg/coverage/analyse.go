package coverage

import (
	"fmt"
)

func Analyse(target Target) (*Result, error) {
	if len(target.Branches) == 0 {
		return nil, fmt.Errorf("target contains empty branches. Target: %s", target.File)
	}

	coveredLines := 0
	uncoveredLines := 0

	for _, b := range target.Branches {
		if b.To < b.From {
			return nil, fmt.Errorf("branch starting line is behind ending line. Target file: %s, Branch: %+v", target.File, b)
		}

		lines := b.To - b.From + 1

		if b.Covered {
			coveredLines += lines
		} else {
			uncoveredLines += lines
		}
	}

	return &Result{
		target.File,
		coveredLines,
		uncoveredLines,
	}, nil
}
