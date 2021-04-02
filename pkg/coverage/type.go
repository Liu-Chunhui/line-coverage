package coverage

type Result struct {
	Target         string
	CoveredLines   int
	UncoveredLines int
}

// github.com/yesino/example/test/testdata.go:41.21,47.3 2 0
type coverageProfile struct {
	Target         string // github.com/yesino/example/test/testdata.go
	TargetFilename string // /Users/yesino/Documents/GitHub/line-coverage/test/data/testcodefile
	StartLine      int    // 41
	StartPosition  int    // 21
	FinishLine     int    // 47
	FinishPosition int    // 3
	Statements     int    // 2
	Executions     int    // 0
}

// github.com/yesino/example/test/testdata.go:41.21,47.3 2 0
type branch struct {
	Start   int  // 42 	(adjusted. was 41)
	Finish  int  // 46 	(adjusted. was 47)
	Covered bool // false	(0 executions)
}
