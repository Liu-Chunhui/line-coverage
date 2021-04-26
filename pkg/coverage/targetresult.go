package coverage

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	log "github.com/sirupsen/logrus"
)

// target: github.com/yesino/line-coverage/test/testdata.go
// branches: [x]*branch from target
func calculateTargetResult(target string, branches []*branch) (*Result, error) {
	if branches == nil {
		return nil, fmt.Errorf("target contains empty branches. Target: %s", target)
	}

	log.Debug(spew.Sprintln("targetBranches: %s, %+v", target, branches))

	coveredLines := make(map[int]struct{})
	uncoveredLines := make(map[int]struct{})

	for _, b := range branches {
		if b.Finish < b.Start {
			return nil, fmt.Errorf("branch starting line is behind ending line. Target file: %s, branch: %+v", target, b)
		}

		for i := b.Start; i <= b.Finish; i++ {
			if b.Covered {
				if _, ok := coveredLines[i]; !ok {
					coveredLines[i] = struct{}{}
				}
			} else {
				if _, ok := uncoveredLines[i]; !ok {
					uncoveredLines[i] = struct{}{}
				}
			}
		}
	}

	// if one line is marked as covered and uncovered, then adjusted it as covered
	for key := range coveredLines {
		delete(uncoveredLines, key)
	}

	log.Debug(spew.Sprintln("coveredLines: %s, %+v", target, coveredLines))
	log.Debug(spew.Sprintln("uncoveredLines: %s, %+v", target, uncoveredLines))

	return &Result{
		target,
		len(coveredLines),
		len(uncoveredLines),
	}, nil
}
