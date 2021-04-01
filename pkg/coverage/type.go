package coverage

type Profile struct {
	targetfile string
}

type Branch struct {
	Target  string
	Start   int
	Finish  int
	Covered bool
}

type Result struct {
	Target         string
	CoveredLines   int
	UncoveredLines int
}
