package sugar

func Contains[T comparable](collection []T, element T) bool {
	for _, t := range collection {
		if t == element {
			return true
		}
	}
	return false
}
