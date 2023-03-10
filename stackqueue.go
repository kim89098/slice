package slice

// Pop removes and returns the last element of a slice, along with the remaining slice.
// If the slice is empty, Pop returns a zero value and a nil slice.
func Pop[T any](s []T) (T, []T) {
	if len(s) == 0 {
		var zero T
		return zero, nil
	}

	last := len(s) - 1
	return s[last], s[:last]
}

// Push adds an element to the end of a slice and returns the resulting slice.
func Push[T any](s []T, v T) []T {
	return append(s, v)
}

// Shift removes and returns the first element of a slice, along with the remaining slice.
// If the slice is empty, Shift returns a zero value and a nil slice.
func Shift[T any](s []T) (T, []T) {
	if len(s) == 0 {
		var zero T
		return zero, nil
	}

	return s[0], s[1:]
}

// Unshift adds an element to the beginning of a slice and returns the resulting slice.
func Unshift[T any](s []T, v T) []T {
	return append([]T{v}, s...)
}
