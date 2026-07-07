package options

type options struct {
	AreSourcesCopied  bool
	HasESLint         bool
	Hints             []string
	ShouldUseDefaults bool
	TS                struct {
		Files   []string
		Include []string
	}
}

func (options *options) AddFile(fileName string) {
	options.TS.Files = append(options.TS.Files, fileName)
}

func (options *options) AddInclude(include string) {
	options.TS.Include = append(options.TS.Include, include)
}

var Options options
