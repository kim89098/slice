package slice

// Count returns the number of elements in the slice that satisfy the given function.
func Count[S ~[]E, E any](s S, f func(E) bool) int {
	var c int
	for _, v := range s {
		if f(v) {
			c++
		}
	}
	return c
}

// Equals returns true if the two slices contain the same elements in the same order.
func Equals[S ~[]E, E comparable](a, b S) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

// EqualsAnyOrder returns true if the two slices contain the same elements, regardless of order.
func EqualsAnyOrder[S ~[]E, E comparable](a, b S) bool {
	if len(a) != len(b) {
		return false
	}

	ma := make(map[E]int)
	for _, v := range a {
		ma[v]++
	}

	mb := make(map[E]int)
	for _, v := range b {
		mb[v]++
	}

	for k, v := range ma {
		if mb[k] != v {
			return false
		}
	}

	return true
}

// Every returns true if the given function returns true for every element in the slice.
func Every[S ~[]E, E any](s S, f func(E) bool) bool {
	for _, v := range s {
		if !f(v) {
			return false
		}
	}

	return true
}

// Find returns the first element in the slice that satisfies the given function, along with a boolean indicating whether such an element was found.
func Find[S ~[]E, E any](s S, f func(E) bool) (E, bool) {
	for _, v := range s {
		if f(v) {
			return v, true
		}
	}

	var zero E
	return zero, false
}

// FindDefault returns the first element in s that satisfies the predicate f.
// If no such element is found, it returns defaultValue.
func FindDefault[S ~[]E, E any](s S, f func(E) bool, defaultValue E) E {
	if v, ok := Find(s, f); ok {
		return v
	}

	return defaultValue
}

// FindIndex returns the index of the first element in the slice that satisfies the given function, or -1 if no such element was found.
func FindIndex[S ~[]E, E any](s S, f func(E) bool) int {
	for i, v := range s {
		if f(v) {
			return i
		}
	}

	return -1
}

// FindLast returns the last element in the slice that satisfies the given function, along with a boolean indicating whether such an element was found.
func FindLast[S ~[]E, E any](s S, f func(E) bool) (E, bool) {
	for i := len(s) - 1; i >= 0; i-- {
		if v := s[i]; f(v) {
			return v, true
		}
	}

	var zero E
	return zero, false
}

// FindLastDefault returns the last element in s that satisfies the predicate f.
// If no such element is found, it returns defaultValue.
func FindLastDefault[S ~[]E, E any](s S, f func(E) bool, defaultValue E) E {
	if v, ok := FindLast(s, f); ok {
		return v
	}

	return defaultValue
}

// FindLastIndex returns the index of the last element in the slice that satisfies the given function, or -1 if no such element was found.
func FindLastIndex[S ~[]E, E any](s S, f func(E) bool) int {
	for i := len(s) - 1; i >= 0; i-- {
		if v := s[i]; f(v) {
			return i
		}
	}

	return -1
}

// Includes returns true if the given value is found in the slice, otherwise false.
func Includes[S ~[]E, E comparable](s S, v E) bool {
	for _, e := range s {
		if e == v {
			return true
		}
	}

	return false
}

// IndexOf returns the index of the first occurrence of the given value in the slice, or -1 if not found.
func IndexOf[S ~[]E, E comparable](s S, v E) int {
	for i, e := range s {
		if e == v {
			return i
		}
	}

	return -1
}

// IndexOfFrom returns the index of the first occurrence of the given value in the slice, starting from the given index, or -1 if not found.
func IndexOfFrom[S ~[]E, E comparable](s S, v E, from int) int {
	for i, n := from, len(s); i < n; i++ {
		if e := s[i]; e == v {
			return i
		}
	}

	return -1
}

// LastIndexOf returns the index of the last occurrence of the given value in the slice, or -1 if not found.
func LastIndexOf[S ~[]E, E comparable](s S, v E) int {
	for i := len(s) - 1; i >= 0; i-- {
		if e := s[i]; e == v {
			return i
		}
	}

	return -1
}

// Most returns the element in the slice that satisfies the given function and is "most" according to that function. The definition of "most" is left up to the caller of this function.
func Most[S ~[]E, E any](s S, f func(v, most E) bool) E {
	if len(s) == 0 {
		var zero E
		return zero
	}

	return Reduce(s[1:], func(v, most E) E {
		if f(v, most) {
			return v
		}
		return most
	}, s[0])
}

// Some returns true if at least one element in the slice satisfies a predicate function.
func Some[S ~[]E, E any](s S, f func(E) bool) bool {
	for _, v := range s {
		if f(v) {
			return true
		}
	}

	return false
}
