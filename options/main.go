package options

// hints
type hints struct {
	Hints []string
}

func (hints *hints) Add(hint string) {
	hints.Hints = append(hints.Hints, hint)
}

// ts
type ts struct {
	Files   []string
	Include []string
}

func (ts *ts) AddFile(fileName string) {
	ts.Files = append(ts.Files, fileName)
}

func (ts *ts) AddInclude(include string) {
	ts.Include = append(ts.Include, include)
}

// options
type options struct {
	AreSourcesCopied   bool
	CanInstallPackages bool
	Hints              hints
	ShouldUseDefaults  bool
	TS                 ts
}

var Options = options{CanInstallPackages: true}
