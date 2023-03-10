// Package slice provides functions for manipulating slices.
package slice

import (
	"math/rand"
	"sort"
	"time"

	"golang.org/x/exp/constraints"
)

// Number is a type constraint that allows only complex, float, and integer types.
type Number interface {
	constraints.Complex | constraints.Float | constraints.Integer
}

// Dedup returns a new slice containing only the unique elements of the input slice.
func Dedup[S ~[]E, E comparable](s S) S {
	if len(s) == 0 {
		return nil
	}

	n := make(S, 0, len(s))
	m := make(map[E]bool)

	for _, v := range s {
		if !m[v] {
			n = append(n, v)
			m[v] = true
		}
	}

	return n
}

// Fill fills the entire slice with a given value.
func Fill[S ~[]E, E any](s S, v E) {
	for i := range s {
		s[i] = v
	}
}

// FillRange fills a range [start:end) of the slice with a given value.
func FillRange[S ~[]E, E any](s S, v E, start, end int) {
	s = s[start:end]

	for i := range s {
		s[i] = v
	}
}

// Filter returns a new slice containing only the elements from the original slice that satisfy the given function.
func Filter[S ~[]E, E any](s S, f func(E) bool) S {
	if len(s) == 0 {
		return nil
	}

	n := make(S, 0, len(s))

	for _, v := range s {
		if f(v) {
			n = append(n, v)
		}
	}

	return n
}

// Flat flattens a slice of slices into a single slice.
func Flat[SS ~[]S, S ~[]E, E any](s SS) S {
	l := Reduce(s, func(v S, acc int) int { return len(v) + acc }, 0)
	if l == 0 {
		return nil
	}

	n := make(S, 0, l)
	for _, v := range s {
		n = append(n, v...)
	}
	return n
}

// ForEach applies the given function to every element in the slice.
func ForEach[S ~[]E, E any](s S, f func(E)) {
	for _, v := range s {
		f(v)
	}
}

// ForEachIndex applies the given function to every element in the slice along with its index.
func ForEachIndex[S ~[]E, E any](s S, f func(v E, i int)) {
	for i, v := range s {
		f(v, i)
	}
}

// Group groups the elements of a slice by applying a given function to each element and using the result as a key in a map.
func Group[S ~[]E, E any, K comparable](s S, f func(E) K) map[K]S {
	m := make(map[K]S)

	for _, v := range s {
		key := f(v)
		m[key] = append(m[key], v)
	}

	return m
}

// Map applies the given function to every element in the slice and returns a new slice containing the results.
func Map[S ~[]E, E, R any](s S, f func(E) R) []R {
	if len(s) == 0 {
		return nil
	}

	n := make([]R, len(s))

	for i, v := range s {
		n[i] = f(v)
	}

	return n
}

// Random returns a random element from a slice and a new slice with the randomly selected element removed. If the input slice is empty, it returns a zero value and a nil slice.
func Random[S ~[]E, E any](s S) (E, S) {
	if len(s) == 0 {
		var zero E
		return zero, nil
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(s))
	return s[r], append(s[:r], s[r+1:]...)
}

// Reduce applies a function to each element of a slice, accumulating the result into an initial value.
func Reduce[S ~[]E, E any, R any](s S, f func(v E, acc R) R, init R) R {
	for _, v := range s {
		init = f(v, init)
	}

	return init
}

// ReduceRight applies a function to each element of a slice in reverse order,
// accumulating the result into an initial value.
func ReduceRight[S ~[]E, E, R any](s S, f func(v E, acc R) R, init R) R {
	for i := len(s) - 1; i >= 0; i-- {
		init = f(s[i], init)
	}

	return init
}

// Reverse reverses the elements of a slice in place.
func Reverse[S ~[]E, E any](s S) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// ReverseCopy returns a new slice with the elements in reverse order.
func ReverseCopy[S ~[]E, E any](s S) S {
	n := make(S, len(s))

	for i, v := range s {
		n[len(s)-i-1] = v
	}

	return n
}

// Sort sorts the elements of slice s in increasing order, according to the
// order defined by the less function. The less function returns true if the
// first element should be ordered before the second element.
func Sort[S ~[]E, E any](s S, less func(a, b E) bool) {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
}

// Sum returns the sum of all elements in a slice of type T.
func Sum[S ~[]E, E Number](s S) E {
	var sum E
	for _, v := range s {
		sum += v
	}
	return sum
}
