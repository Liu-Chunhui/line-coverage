package coverage

import "strings"

// convertProfileToBranch adjusts starting and finishing line number from coverage profile,
// which is to create Branch for calculation later.
func convertProfileToBranch(profile *coverageProfile, codeInLines []string) (string, *branch) {
	startLine := profile.StartLine
	finishLine := profile.FinishLine

	start := codeInLines[startLine-1]
	// when start position is the end of this line, then skip this line
	if start[profile.StartPosition] == '\n' {
		startLine += 1
	}

	finish := codeInLines[finishLine-1]
	// when ending position is the end of the line, then set finish line to previous line
	if strings.TrimSpace(finish) == "}" {
		finishLine -= 1
	}

	return profile.Target,
		&branch{
			Start:   startLine,
			Finish:  finishLine,
			Covered: profile.Statements > 0 && profile.Executions > 0,
		}
}
