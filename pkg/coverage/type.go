package coverage

type Branch struct {
	From    int
	To      int
	Covered bool
}

type Target struct {
	File     string
	Branches []Branch
}

type Result struct {
	File           string
	CoveredLines   int
	UncoveredLines int
}
