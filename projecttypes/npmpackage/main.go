package npmpackage

func Create() {
	defer func() {
		_ = recover()
	}()

	initPackage()
}
