package sugar

import (
	"math/rand/v2"

	"golang.org/x/exp/constraints"
)

// Map manipulates a slice and transforms it to a slice of another type.
func Map[T any, R any](collection []T, iteratee func(T, int) R) []R {
	result := make([]R, len(collection))

	for i, item := range collection {
		result[i] = iteratee(item, i)
	}

	return result
}

// Filter iterates over elements of collection, returning an array of all elements predicate returns truthy for.
func Filter[T any](collection []T, predicate func(T, int) bool) []T {
	result := make([]T, 0, len(collection))

	for i, item := range collection {
		if predicate(item, i) {
			result = append(result, item)
		}
	}

	return result
}

// Reduce reduces collection to a value which is the accumulated result of running each element in collection
// thru iteratee, where each successive invocation is supplied the return value of the previous.
func Reduce[T any, R any](collection []T, accumulator func(R, T, int) R, initial R) R {
	for i, item := range collection {
		initial = accumulator(initial, item, i)
	}

	return initial
}

// Find searches for the first element of a collection
func Find[T any](collection []T, predicate func(T) bool) (T, bool) {
	for _, item := range collection {
		if predicate(item) {
			return item, true
		}
	}

	var zero T
	return zero, false
}

// FindIndex searches for the first element of a collection and returns its index
func FindIndex[T any](collection []T, predicate func(T) bool) int {
	for i, item := range collection {
		if predicate(item) {
			return i
		}
	}

	return -1
}

// FindLast searches for the last element of a collection
func FindLast[T any](collection []T, predicate func(T) bool) (T, bool) {
	length := len(collection)

	for i := length - 1; i >= 0; i-- {
		if predicate(collection[i]) {
			return collection[i], true
		}
	}

	var zero T
	return zero, false
}

// FindLastIndex searches for the last element of a collection and returns its index
func FindLastIndex[T any](collection []T, predicate func(T) bool) int {
	length := len(collection)

	for i := length - 1; i >= 0; i-- {
		if predicate(collection[i]) {
			return i
		}
	}

	return -1
}

// Uniq returns a duplicate-free version of an array.
func Uniq[T comparable](collection []T) []T {
	result := make([]T, 0, len(collection))
	seen := make(map[T]struct{}, len(collection))

	for _, item := range collection {
		if _, ok := seen[item]; !ok {
			seen[item] = struct{}{}
			result = append(result, item)
		}
	}

	return result
}

// UniqBy returns a duplicate-free version of an array using a transformation function.
func UniqBy[T any, U comparable](collection []T, iteratee func(T) U) []T {
	result := make([]T, 0, len(collection))
	seen := make(map[U]struct{}, len(collection))

	for _, item := range collection {
		key := iteratee(item)
		if _, ok := seen[key]; !ok {
			seen[key] = struct{}{}
			result = append(result, item)
		}
	}

	return result
}

// GroupBy returns an object composed of keys generated from the results of running each element of collection thru iteratee.
func GroupBy[T any, U comparable](collection []T, iteratee func(T) U) map[U][]T {
	result := map[U][]T{}

	for _, item := range collection {
		key := iteratee(item)
		result[key] = append(result[key], item)
	}

	return result
}

// Chunk creates an array of elements split into groups the length of size.
func Chunk[T any](collection []T, size int) [][]T {
	if size <= 0 {
		return [][]T{}
	}

	result := make([][]T, 0, (len(collection)+size-1)/size)

	for i := 0; i < len(collection); i += size {
		end := i + size
		if end > len(collection) {
			end = len(collection)
		}
		result = append(result, collection[i:end])
	}

	return result
}

// Reverse reverses array so that the first element becomes the last, the second element becomes the second to last, and so on.
func Reverse[T any](collection []T) []T {
	length := len(collection)
	half := length / 2

	for i := 0; i < half; i = i + 1 {
		j := length - 1 - i
		collection[i], collection[j] = collection[j], collection[i]
	}

	return collection
}

// Flatten flattens array a single level deep.
func Flatten[T any](collection [][]T) []T {
	totalLen := 0
	for _, subCollection := range collection {
		totalLen += len(subCollection)
	}

	result := make([]T, 0, totalLen)
	for _, subCollection := range collection {
		result = append(result, subCollection...)
	}

	return result
}

// Shuffle creates an array of shuffled values
func Shuffle[T any](collection []T) []T {
	result := make([]T, len(collection))
	copy(result, collection)

	for i := len(result) - 1; i > 0; i-- {
		j := rand.IntN(i + 1)
		result[i], result[j] = result[j], result[i]
	}

	return result
}

// Sample gets a random element from array.
func Sample[T any](collection []T) T {
	size := len(collection)
	if size == 0 {
		var zero T
		return zero
	}

	return collection[rand.IntN(size)]
}

// SampleSize gets n random elements at unique keys from collection up to the size of collection.
func SampleSize[T any](collection []T, count int) []T {
	size := len(collection)

	copy := append([]T(nil), collection...)

	results := []T{}

	for i := 0; i < size && i < count; i++ {
		copyLength := size - i
		index := rand.IntN(copyLength)
		results = append(results, copy[index])

		// Removes the item at the sample index.
		copy[index] = copy[copyLength-1]
	}

	return results
}

// Min returns the minimum value of a collection.
func Min[T constraints.Ordered](collection []T) T {
	var min T

	if len(collection) == 0 {
		return min
	}

	min = collection[0]
	for i := 1; i < len(collection); i++ {
		item := collection[i]

		if item < min {
			min = item
		}
	}

	return min
}

// Max returns the maximum value of a collection.
func Max[T constraints.Ordered](collection []T) T {
	var max T

	if len(collection) == 0 {
		return max
	}

	max = collection[0]
	for i := 1; i < len(collection); i++ {
		item := collection[i]

		if item > max {
			max = item
		}
	}

	return max
}

// MinBy returns the minimum value of a collection.
func MinBy[T any, U constraints.Ordered](collection []T, iteratee func(T) U) T {
	var minItem T

	if len(collection) == 0 {
		return minItem
	}

	minItem = collection[0]
	min := iteratee(minItem)

	for i := 1; i < len(collection); i++ {
		item := collection[i]
		value := iteratee(item)

		if value < min {
			min = value
			minItem = item
		}
	}

	return minItem
}

// MaxBy returns the maximum value of a collection.
func MaxBy[T any, U constraints.Ordered](collection []T, iteratee func(T) U) T {
	var maxItem T

	if len(collection) == 0 {
		return maxItem
	}

	maxItem = collection[0]
	max := iteratee(maxItem)

	for i := 1; i < len(collection); i++ {
		item := collection[i]
		value := iteratee(item)

		if value > max {
			max = value
			maxItem = item
		}
	}

	return maxItem
}

// Every returns true if all elements in collection pass the predicate test.
func Every[T any](collection []T, predicate func(T) bool) bool {
	for _, item := range collection {
		if !predicate(item) {
			return false
		}
	}

	return true
}

// Some returns true if at least one element in collection passes the predicate test.
func Some[T any](collection []T, predicate func(T) bool) bool {
	for _, item := range collection {
		if predicate(item) {
			return true
		}
	}

	return false
}

// Count returns the number of elements in the collection that pass the predicate test.
func Count[T any](collection []T, predicate func(T) bool) int {
	count := 0
	for _, item := range collection {
		if predicate(item) {
			count++
		}
	}

	return count
}
