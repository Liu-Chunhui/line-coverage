package fileparser

import "regexp"

var (
	CoverageProfileExcludingRules = []string{
		`^\s*$`,     // empty line
		`^mode: .*`, // first line
	}

	GoModGetModuleNameRule = "^module .*" // first line is the module name
)

func MatchPattern(line string, patterns ...string) (bool, error) {
	for _, p := range patterns {
		match, err := regexp.Match(p, []byte(line))
		if err != nil {
			return true, err
		}

		if match {
			return true, nil
		}
	}

	return false, nil
}
