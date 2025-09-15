package sugar

import (
	"math"

	"golang.org/x/exp/constraints"
)

// Clamp clamps number within the inclusive lower and upper bounds.
func Clamp[T constraints.Ordered](value, min, max T) T {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

// Sum returns the sum of all values.
func Sum[T constraints.Integer | constraints.Float](collection []T) T {
	if len(collection) == 0 {
		return T(0)
	}
	var sum T
	for _, val := range collection {
		sum += val
	}
	return sum
}

// SumBy returns the sum of all values after applying the transformation.
func SumBy[T any, R constraints.Integer | constraints.Float](collection []T, iteratee func(T) R) R {
	var sum R
	for _, val := range collection {
		sum += iteratee(val)
	}
	return sum
}

// Abs returns the absolute value.
func Abs[T constraints.Signed](value T) T {
	if value < 0 {
		return -value
	}
	return value
}

// Range creates an array of numbers progressing from start up to, but not including, stop.
func Range(start, stop int) []int {
	if start >= stop {
		return []int{}
	}

	result := make([]int, stop-start)
	for i := start; i < stop; i++ {
		result[i-start] = i
	}
	return result
}

// RangeWithStep creates an array of numbers progressing from start up to, but not including, stop.
func RangeWithStep(start, stop, step int) []int {
	if step == 0 || (step > 0 && start >= stop) || (step < 0 && start <= stop) {
		return []int{}
	}

	result := []int{}
	if step > 0 {
		for i := start; i < stop; i += step {
			result = append(result, i)
		}
	} else {
		for i := start; i > stop; i += step {
			result = append(result, i)
		}
	}
	return result
}

// Mean returns the average of values.
func Mean[T constraints.Integer | constraints.Float](collection []T) float64 {
	if len(collection) == 0 {
		return 0
	}

	var sum T
	for _, val := range collection {
		sum += val
	}
	return float64(sum) / float64(len(collection))
}

// Median returns the median of values.
func Median[T constraints.Integer | constraints.Float](collection []T) float64 {
	if len(collection) == 0 {
		return 0
	}

	// Convert to float64 slice for sorting
	floats := make([]float64, len(collection))
	for i, v := range collection {
		floats[i] = float64(v)
	}

	// Simple selection sort
	for i := 0; i < len(floats); i++ {
		minIdx := i
		for j := i + 1; j < len(floats); j++ {
			if floats[j] < floats[minIdx] {
				minIdx = j
			}
		}
		if minIdx != i {
			floats[i], floats[minIdx] = floats[minIdx], floats[i]
		}
	}

	length := len(floats)
	if length%2 == 0 {
		return (floats[length/2-1] + floats[length/2]) / 2
	}
	return floats[length/2]
}

// Pow returns x**y, the base-x exponential of y.
func Pow[T constraints.Integer | constraints.Float](x, y T) float64 {
	return math.Pow(float64(x), float64(y))
}

// Sqrt returns the square root of x.
func Sqrt[T constraints.Integer | constraints.Float](x T) float64 {
	return math.Sqrt(float64(x))
}

// Ceil returns the least integer value greater than or equal to x.
func Ceil(x float64) float64 {
	return math.Ceil(x)
}

// Floor returns the greatest integer value less than or equal to x.
func Floor(x float64) float64 {
	return math.Floor(x)
}

// Round returns the nearest integer, rounding half away from zero.
func Round(x float64) float64 {
	return math.Round(x)
}

// Trunc returns the integer value of x.
func Trunc(x float64) float64 {
	return math.Trunc(x)
}

// IsNaN reports whether f is an IEEE 754 "not-a-number" value.
func IsNaN[T constraints.Float](x T) bool {
	return math.IsNaN(float64(x))
}

// IsInf reports whether f is an infinity.
func IsInf[T constraints.Float](x T) bool {
	return math.IsInf(float64(x), 0)
}
