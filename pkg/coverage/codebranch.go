package coverage

import (
	"strings"
)

// convertProfileToBranch adjusts starting and finishing line number from coverage profile,
// which is to create Branch for calculation later.
func convertProfileToBranch(profile *coverageProfile, codeInLines []string) (string, []*branch) {
	adjustedStartLine := profile.StartLine
	adjustedFinishLine := profile.FinishLine

	adjustedStartLine += startLineAdjustment(codeInLines, profile.StartLine, profile.StartPosition)
	adjustedFinishLine -= finishLineAdjustment(codeInLines, profile.FinishLine, profile.FinishPosition)

	// to check if there are any new lines
	branches := splitByNewLine(
		codeInLines,
		adjustedStartLine,
		adjustedFinishLine,
		profile.Statements >= 0 && profile.Executions > 0,
	)

	return profile.Target, branches
}

// if startLine and endLine contains new lines, then need to break it into sub branches
func splitByNewLine(codeInLines []string, startLine int, endLine int, covered bool) []*branch {
	var branches []*branch

	lines := codeInLines[startLine-1 : endLine]
	start := startLine
	end := startLine
	for i, l := range lines {
		end = startLine + i
		// if current line is '\n', then create a branch
		if l == "\n" {
			branches = append(branches, &branch{
				Start:   start,
				Finish:  end - 1,
				Covered: covered,
			})
			// reset start to next line
			start = end + 1
			continue
		}
	}

	branches = append(branches, &branch{
		Start:   start,
		Finish:  end,
		Covered: covered,
	})

	return branches
}

// when start position is the end of this line(could contain multiple \n), then skip this line
func startLineAdjustment(codeInLines []string, startLine int, position int) int {
	adjustment := 0

	line := codeInLines[startLine-1]

	// check start char is new line
	switch l := line[position-1]; l {
	case '\n':
		// check if previous char is ':'(case xxxx :\n)
		if position-2 >= 0 && line[position-2] == ':' {
			return adjustment
		}

		adjustment = startEmptyLineAdjustment(codeInLines, startLine+1, 1)
	case '{':
		// check if next character is '\n' (func () {})
		if len(line) >= position+1 && line[position] == '\n' {
			adjustment = startEmptyLineAdjustment(codeInLines, startLine+1, 1)
		}
	}

	return adjustment
}

func startEmptyLineAdjustment(codeInLines []string, startLine int, adjustment int) int {
	line := codeInLines[startLine-1]
	if len(strings.TrimSpace(strings.ReplaceAll(line, "{", ""))) == 0 {
		return startEmptyLineAdjustment(codeInLines, startLine+1, adjustment+1)
	}

	return adjustment
}

func finishLineAdjustment(codeInLines []string, finishLine int, position int) int {
	statement := strings.TrimSpace(codeInLines[finishLine-1][0:position]) // trim "\t", " ", "\n"
	statement = strings.ReplaceAll(statement, "}", "")

	if statement == "" ||
		statement == "(" ||
		statement == "()" {
		return adjustBackwards(codeInLines, finishLine-1, 1)
	}

	return 0
}

func adjustBackwards(codeInLines []string, finishLine int, adjustment int) int {
	line := strings.TrimSpace(codeInLines[finishLine-1]) // trim "\t", " ", "\n"
	line = strings.ReplaceAll(line, "}", "")             // trim all '}'.

	if line == "" ||
		line == "(" ||
		line == "()" {
		return adjustBackwards(codeInLines, finishLine-1, adjustment+1)
	}

	return adjustment
}
