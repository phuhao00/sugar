package sugar

func DecorateFn(before, fn, after func()) func() {
	return func() {
		before()
		fn()
		after()
	}
}
