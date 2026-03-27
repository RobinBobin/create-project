package options

type options struct {
	IsESM          bool
	IsProjectReset bool
	TS             struct {
		Files   []string
		Include []string
	}
}

var Options options
