package fileparser

var (
	CoverageProfileExcludingRules = []string{
		`^\s*$`,     // empty line
		`^mode: .*`, // first line
	}
)
