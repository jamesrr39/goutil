package must

// Must panics on error
func Must(err error) {
	if nil != err {
		panic(err)
	}
}
