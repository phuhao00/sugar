package sugar

// Keys creates an array of the map keys.
func Keys[K comparable, V any](in map[K]V) []K {
	result := make([]K, 0, len(in))

	for k := range in {
		result = append(result, k)
	}

	return result
}

// Values creates an array of the map values.
func Values[K comparable, V any](in map[K]V) []V {
	result := make([]V, 0, len(in))

	for _, v := range in {
		result = append(result, v)
	}

	return result
}

// PickBy creates an object composed of the object properties predicate returns truthy for.
func PickBy[K comparable, V any](in map[K]V, predicate func(K, V) bool) map[K]V {
	result := make(map[K]V, len(in))

	for k, v := range in {
		if predicate(k, v) {
			result[k] = v
		}
	}

	return result
}

// PickByKeys returns same map type filtered by given keys.
func PickByKeys[K comparable, V any](in map[K]V, keys []K) map[K]V {
	result := make(map[K]V, len(keys))

	for k, v := range in {
		if Contains(keys, k) {
			result[k] = v
		}
	}

	return result
}

// PickByValues returns same map type filtered by given values.
func PickByValues[K comparable, V comparable](in map[K]V, values []V) map[K]V {
	result := make(map[K]V)

	for k, v := range in {
		if Contains(values, v) {
			result[k] = v
		}
	}

	return result
}

// OmitBy is the opposite of PickBy; this method creates an object composed of the object properties predicate does not return truthy for.
func OmitBy[K comparable, V any](in map[K]V, predicate func(K, V) bool) map[K]V {
	result := make(map[K]V, len(in))

	for k, v := range in {
		if !predicate(k, v) {
			result[k] = v
		}
	}

	return result
}

// OmitByKeys returns same map type filtered by given keys.
func OmitByKeys[K comparable, V any](in map[K]V, keys []K) map[K]V {
	result := make(map[K]V, len(in))

	for k, v := range in {
		if !Contains(keys, k) {
			result[k] = v
		}
	}

	return result
}

// OmitByValues returns same map type filtered by given values.
func OmitByValues[K comparable, V comparable](in map[K]V, values []V) map[K]V {
	result := make(map[K]V, len(in))

	for k, v := range in {
		if !Contains(values, v) {
			result[k] = v
		}
	}

	return result
}

// Entries transforms a map into array of key/value pairs.
func Entries[K comparable, V any](in map[K]V) []Entry[K, V] {
	result := make([]Entry[K, V], 0, len(in))

	for k, v := range in {
		result = append(result, Entry[K, V]{k, v})
	}

	return result
}

// FromEntries transforms an array of key/value pairs into a map.
func FromEntries[K comparable, V any](entries []Entry[K, V]) map[K]V {
	result := make(map[K]V, len(entries))

	for _, entry := range entries {
		result[entry.Key] = entry.Value
	}

	return result
}

// Invert creates an object composed of the inverted keys and values of object.
func Invert[K comparable, V comparable](in map[K]V) map[V]K {
	result := make(map[V]K, len(in))

	for k, v := range in {
		result[v] = k
	}

	return result
}

// Assign merges multiple maps from left to right.
func Assign[K comparable, V any](maps ...map[K]V) map[K]V {
	size := 0
	for _, m := range maps {
		size += len(m)
	}

	result := make(map[K]V, size)

	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}

	return result
}

// MapKeys manipulates a map keys and transforms it to a map of another type.
func MapKeys[K comparable, V any, R comparable](in map[K]V, iteratee func(K, V) R) map[R]V {
	result := make(map[R]V, len(in))

	for k, v := range in {
		result[iteratee(k, v)] = v
	}

	return result
}

// MapValues manipulates a map values and transforms it to a map of another type.
func MapValues[K comparable, V any, R any](in map[K]V, iteratee func(K, V) R) map[K]R {
	result := make(map[K]R, len(in))

	for k, v := range in {
		result[k] = iteratee(k, v)
	}

	return result
}

// HasKey returns whether the given key exists.
func HasKey[K comparable, V any](in map[K]V, key K) bool {
	_, ok := in[key]
	return ok
}

// HasValue returns whether the given value exists.
func HasValue[K comparable, V comparable](in map[K]V, value V) bool {
	for _, v := range in {
		if v == value {
			return true
		}
	}
	return false
}
