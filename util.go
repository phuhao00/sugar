package sugar

import (
	"context"
	"fmt"
	"time"
)

// Must wraps a function call to panics if second argument is error or false, returns the value otherwise.
func Must[T any](val T, err error) T {
	if err != nil {
		panic(err)
	}
	return val
}

// Must0 has the same behavior as Must, but callback returns no variable.
func Must0(err error) {
	if err != nil {
		panic(err)
	}
}

// Must1 is an alias to Must
func Must1[T any](val T, err error) T {
	return Must(val, err)
}

// Must2 has the same behavior as Must, but callback returns 2 variables.
func Must2[T any, R any](val1 T, val2 R, err error) (T, R) {
	if err != nil {
		panic(err)
	}
	return val1, val2
}

// Must3 has the same behavior as Must, but callback returns 3 variables.
func Must3[T any, R any, S any](val1 T, val2 R, val3 S, err error) (T, R, S) {
	if err != nil {
		panic(err)
	}
	return val1, val2, val3
}

// TryFunc calls the function and return false in case of error and panic.
func TryFunc(callback func() error) bool {
	defer func() {
		recover()
	}()

	return callback() == nil
}

// Try1 calls the function and return false in case of error and panic.
func Try1[T any](callback func() (T, error)) bool {
	defer func() {
		recover()
	}()

	_, err := callback()
	return err == nil
}

// Try2 calls the function and return false in case of error and panic.
func Try2[T, R any](callback func() (T, R, error)) bool {
	defer func() {
		recover()
	}()

	_, _, err := callback()
	return err == nil
}

// TryOr calls the function and return a default value in case of error and on panic.
func TryOr[T any](callback func() (T, error), fallback T) T {
	defer func() {
		recover()
	}()

	value, err := callback()
	if err != nil {
		return fallback
	}

	return value
}

// TryOr1 calls the function and return a default value in case of error and on panic.
func TryOr1[T any](callback func() (T, error), fallback T) T {
	return TryOr(callback, fallback)
}

// TryOr2 calls the function and return a default value in case of error and on panic.
func TryOr2[T, R any](callback func() (T, R, error), fallback1 T, fallback2 R) (T, R) {
	defer func() {
		recover()
	}()

	val1, val2, err := callback()
	if err != nil {
		return fallback1, fallback2
	}

	return val1, val2
}

// TryCatch calls the function and calls the catch function in case of error.
func TryCatch(callback func() error, catch func()) bool {
	defer func() {
		if r := recover(); r != nil {
			catch()
		}
	}()

	if err := callback(); err != nil {
		catch()
		return false
	}

	return true
}

// TryCatchWithErrorValue calls the function and calls the catch function in case of error.
func TryCatchWithErrorValue(callback func() error, catch func(any)) bool {
	defer func() {
		if r := recover(); r != nil {
			catch(r)
		}
	}()

	if err := callback(); err != nil {
		catch(err)
		return false
	}

	return true
}

// Validate calls the predicate function and returns an error if false.
func Validate(ok bool, format string, args ...any) error {
	if !ok {
		return fmt.Errorf(format, args...)
	}
	return nil
}

// ToPtr returns a pointer to the value.
func ToPtr[T any](x T) *T {
	return &x
}

// FromPtr returns the value pointed to by the pointer.
// Returns zero value if the pointer is nil.
func FromPtr[T any](x *T) T {
	if x == nil {
		var zero T
		return zero
	}
	return *x
}

// FromPtrOr returns the value pointed to by the pointer or fallback.
func FromPtrOr[T any](x *T, fallback T) T {
	if x == nil {
		return fallback
	}
	return *x
}

// IsNil checks if a value is nil or if it's an interface holding a nil value.
func IsNil(x any) bool {
	return x == nil
}

// IsNotNil checks if a value is not nil.
func IsNotNil(x any) bool {
	return !IsNil(x)
}

// Empty returns a zero value of type T.
func Empty[T any]() T {
	var t T
	return t
}

// CoalesceOrEmpty returns the first non-empty value.
func CoalesceOrEmpty[T comparable](values ...T) T {
	var zero T

	for _, v := range values {
		if v != zero {
			return v
		}
	}

	return zero
}

// Coalesce returns the first non-zero value.
func Coalesce[T comparable](values ...T) (T, bool) {
	var zero T

	for _, v := range values {
		if v != zero {
			return v, true
		}
	}

	return zero, false
}

// IfThen returns the result of the callback if condition is true, otherwise returns the zero value.
func IfThen[T any](condition bool, callback func() T) *Optional[T] {
	if condition {
		return &Optional[T]{value: callback(), present: true}
	}
	return &Optional[T]{present: false}
}

// IfF returns the result of the callback if condition is true, otherwise returns the zero value.
func IfF[T any](condition bool, callback func() T) T {
	if condition {
		return callback()
	}
	var zero T
	return zero
}

// SwitchCase creates a switch-case like structure.
func SwitchCase[T comparable, R any](predicate T) *SwitchCaseStruct[T, R] {
	var result R
	return &SwitchCaseStruct[T, R]{
		predicate: predicate,
		result:    result,
		found:     false,
	}
}

// AsyncRun runs given function in a goroutine and returns the result in a channel.
func AsyncRun[T any](callback func() T) <-chan T {
	ch := make(chan T, 1)
	go func() {
		defer close(ch)
		ch <- callback()
	}()
	return ch
}

// AsyncErr runs given function in a goroutine and returns the result and error in channels.
func AsyncErr[T any](callback func() (T, error)) (<-chan T, <-chan error) {
	ch := make(chan T, 1)
	errCh := make(chan error, 1)

	go func() {
		defer func() {
			close(ch)
			close(errCh)
		}()

		result, err := callback()
		ch <- result
		errCh <- err
	}()

	return ch, errCh
}

// DebounceFunc creates a debounced function that delays invoking the callback.
func DebounceFunc[T any](callback func(...T), delay time.Duration) func(...T) {
	var timer *time.Timer
	return func(args ...T) {
		if timer != nil {
			timer.Stop()
		}
		timer = time.AfterFunc(delay, func() {
			callback(args...)
		})
	}
}

// Throttle creates a throttled function that only invokes the callback at most once per duration.
func Throttle[T any](callback func(...T), duration time.Duration) func(...T) {
	var lastCall time.Time
	return func(args ...T) {
		now := time.Now()
		if now.Sub(lastCall) >= duration {
			lastCall = now
			callback(args...)
		}
	}
}

// Times invokes the callback n times, returning an array of the results.
func Times[T any](count int, callback func(int) T) []T {
	if count <= 0 {
		return []T{}
	}

	result := make([]T, count)
	for i := 0; i < count; i++ {
		result[i] = callback(i)
	}
	return result
}

// Repeat builds a slice with N copies of initial value.
func Repeat[T any](count int, initial T) []T {
	if count <= 0 {
		return []T{}
	}

	result := make([]T, count)
	for i := range result {
		result[i] = initial
	}
	return result
}

// Optional represents an optional value.
type Optional[T any] struct {
	value   T
	present bool
}

// IsPresent returns true if the optional contains a value.
func (o *Optional[T]) IsPresent() bool {
	return o.present
}

// IsEmpty returns true if the optional is empty.
func (o *Optional[T]) IsEmpty() bool {
	return !o.present
}

// Get returns the value if present, otherwise returns zero value.
func (o *Optional[T]) Get() T {
	return o.value
}

// OrElse returns the value if present, otherwise returns the fallback.
func (o *Optional[T]) OrElse(fallback T) T {
	if o.present {
		return o.value
	}
	return fallback
}

// OrElseF returns the value if present, otherwise calls the fallback function.
func (o *Optional[T]) OrElseF(callback func() T) T {
	if o.present {
		return o.value
	}
	return callback()
}

// SwitchCaseStruct represents a switch-case structure.
type SwitchCaseStruct[T comparable, R any] struct {
	predicate T
	result    R
	found     bool
}

// Case adds a case to the switch.
func (s *SwitchCaseStruct[T, R]) Case(value T, result R) *SwitchCaseStruct[T, R] {
	if !s.found && s.predicate == value {
		s.result = result
		s.found = true
	}
	return s
}

// CaseF adds a case with a function to the switch.
func (s *SwitchCaseStruct[T, R]) CaseF(value T, callback func() R) *SwitchCaseStruct[T, R] {
	if !s.found && s.predicate == value {
		s.result = callback()
		s.found = true
	}
	return s
}

// Default returns the result or the default value.
func (s *SwitchCaseStruct[T, R]) Default(result R) R {
	if s.found {
		return s.result
	}
	return result
}

// DefaultF returns the result or calls the default function.
func (s *SwitchCaseStruct[T, R]) DefaultF(callback func() R) R {
	if s.found {
		return s.result
	}
	return callback()
}

// WaitFor waits for a condition to be true.
func WaitFor(condition func() bool, timeout time.Duration, interval time.Duration) bool {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		if condition() {
			return true
		}

		select {
		case <-ctx.Done():
			return false
		case <-ticker.C:
			continue
		}
	}
}

// WaitForWithContext waits for a condition to be true with context.
func WaitForWithContext(ctx context.Context, condition func() bool, timeout time.Duration, interval time.Duration) (iterations int, duration time.Duration, ok bool) {
	start := time.Now()

	timeoutCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		iterations++
		if condition() {
			return iterations, time.Since(start), true
		}

		select {
		case <-timeoutCtx.Done():
			return iterations, time.Since(start), false
		case <-ticker.C:
			continue
		}
	}
}
