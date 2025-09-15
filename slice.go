package sugar

// Contains returns true if an element is present in a collection.
func Contains[T comparable](collection []T, element T) bool {
	for _, item := range collection {
		if item == element {
			return true
		}
	}
	return false
}

// ContainsBy returns true if predicate function returns true for any element in collection.
func ContainsBy[T any](collection []T, predicate func(T) bool) bool {
	for _, item := range collection {
		if predicate(item) {
			return true
		}
	}
	return false
}

// Remove removes the first occurrence of element from slice.
func Remove[T comparable](collection []T, element T) []T {
	for i, val := range collection {
		if element == val {
			return append(collection[:i], collection[i+1:]...)
		}
	}
	return collection
}

// RemoveAll removes all occurrences of element from slice.
func RemoveAll[T comparable](collection []T, element T) []T {
	result := make([]T, 0, len(collection))
	for _, val := range collection {
		if val != element {
			result = append(result, val)
		}
	}
	return result
}

// Drop creates a slice with n elements dropped from the beginning.
func Drop[T any](collection []T, n int) []T {
	if n <= 0 {
		return append([]T(nil), collection...)
	}

	if n >= len(collection) {
		return []T{}
	}

	return append([]T(nil), collection[n:]...)
}

// DropRight creates a slice with n elements dropped from the end.
func DropRight[T any](collection []T, n int) []T {
	if n <= 0 {
		return append([]T(nil), collection...)
	}

	if n >= len(collection) {
		return []T{}
	}

	return append([]T(nil), collection[:len(collection)-n]...)
}

// DropWhile creates a slice excluding elements dropped from the beginning.
func DropWhile[T any](collection []T, predicate func(T) bool) []T {
	i := 0
	for ; i < len(collection); i++ {
		if !predicate(collection[i]) {
			break
		}
	}

	result := make([]T, 0, len(collection)-i)
	return append(result, collection[i:]...)
}

// DropRightWhile creates a slice excluding elements dropped from the end.
func DropRightWhile[T any](collection []T, predicate func(T) bool) []T {
	i := len(collection) - 1
	for ; i >= 0; i-- {
		if !predicate(collection[i]) {
			break
		}
	}

	result := make([]T, 0, i+1)
	return append(result, collection[:i+1]...)
}

// Compact creates a slice with all zero values removed.
func Compact[T comparable](collection []T) []T {
	var zero T
	result := make([]T, 0, len(collection))

	for _, item := range collection {
		if item != zero {
			result = append(result, item)
		}
	}

	return result
}

// Union creates an array of unique values from all given arrays.
func Union[T comparable](collections ...[]T) []T {
	seen := make(map[T]struct{})
	result := make([]T, 0)

	for _, collection := range collections {
		for _, item := range collection {
			if _, ok := seen[item]; !ok {
				seen[item] = struct{}{}
				result = append(result, item)
			}
		}
	}

	return result
}

// Intersection creates an array of unique values that are included in all given arrays.
func Intersection[T comparable](collections ...[]T) []T {
	if len(collections) == 0 {
		return []T{}
	}

	if len(collections) == 1 {
		return Uniq(collections[0])
	}

	// Count occurrences
	counts := make(map[T]int)
	for _, collection := range collections {
		seen := make(map[T]struct{})
		for _, item := range collection {
			if _, ok := seen[item]; !ok {
				seen[item] = struct{}{}
				counts[item]++
			}
		}
	}

	// Find items that appear in all collections
	result := make([]T, 0)
	for item, count := range counts {
		if count == len(collections) {
			result = append(result, item)
		}
	}

	return result
}

// Difference creates an array of values from the first array that are not included in the other arrays.
func Difference[T comparable](collection []T, others ...[]T) []T {
	excluded := make(map[T]struct{})

	for _, other := range others {
		for _, item := range other {
			excluded[item] = struct{}{}
		}
	}

	result := make([]T, 0, len(collection))
	for _, item := range collection {
		if _, ok := excluded[item]; !ok {
			result = append(result, item)
		}
	}

	return result
}

// Concat creates a new array concatenating additional arrays.
func Concat[T any](collections ...[]T) []T {
	totalLen := 0
	for _, collection := range collections {
		totalLen += len(collection)
	}

	result := make([]T, 0, totalLen)
	for _, collection := range collections {
		result = append(result, collection...)
	}

	return result
}

// Without creates an array excluding all given values.
func Without[T comparable](collection []T, exclude ...T) []T {
	excludeSet := make(map[T]struct{}, len(exclude))
	for _, item := range exclude {
		excludeSet[item] = struct{}{}
	}

	result := make([]T, 0, len(collection))
	for _, item := range collection {
		if _, ok := excludeSet[item]; !ok {
			result = append(result, item)
		}
	}

	return result
}

// Partition creates an array of elements split into two groups.
func Partition[T any](collection []T, predicate func(T) bool) ([]T, []T) {
	truthy := make([]T, 0)
	falsy := make([]T, 0)

	for _, item := range collection {
		if predicate(item) {
			truthy = append(truthy, item)
		} else {
			falsy = append(falsy, item)
		}
	}

	return truthy, falsy
}
