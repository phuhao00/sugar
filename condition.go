package sugar

// Ternary returns ifOutput if condition is true, otherwise returns elseOutput.
func Ternary[T any](condition bool, ifOutput T, elseOutput T) T {
	if condition {
		return ifOutput
	}
	return elseOutput
}

// TernaryF returns the result of ifCallback if condition is true, otherwise returns the result of elseCallback.
func TernaryF[T any](condition bool, ifCallback func() T, elseCallback func() T) T {
	if condition {
		return ifCallback()
	}
	return elseCallback()
}

// IfElse represents a conditional value container.
type IfElse[T any] struct {
	Result T
	Ok     bool
}

// IfValue creates a new IfElse based on a condition and value.
func IfValue[T any](condition bool, result T) *IfElse[T] {
	if condition {
		return &IfElse[T]{result, true}
	}

	var t T
	return &IfElse[T]{t, false}
}

// IfFunc creates a new IfElse based on a condition and callback.
func IfFunc[T any](condition bool, resultFn func() T) *IfElse[T] {
	if condition {
		return &IfElse[T]{resultFn(), true}
	}

	var t T
	return &IfElse[T]{t, false}
}

// ElseIf adds another condition to check if the previous condition was false.
func (i *IfElse[T]) ElseIf(condition bool, result T) *IfElse[T] {
	if condition && !i.Ok {
		i.Result = result
		i.Ok = true
	}
	return i
}

// ElseIfFunc adds another condition with a callback to check if the previous condition was false.
func (i *IfElse[T]) ElseIfFunc(condition bool, resultFn func() T) *IfElse[T] {
	if condition && !i.Ok {
		i.Result = resultFn()
		i.Ok = true
	}
	return i
}

// Else returns the result if any condition was true, otherwise returns the fallback.
func (i *IfElse[T]) Else(result T) T {
	if i.Ok {
		return i.Result
	}
	return result
}

// ElseFunc returns the result if any condition was true, otherwise calls the callback.
func (i *IfElse[T]) ElseFunc(resultFn func() T) T {
	if i.Ok {
		return i.Result
	}
	return resultFn()
}
