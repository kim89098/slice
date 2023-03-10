// Package slice provides functions for manipulating slices.
package slice

import (
	"math/rand"
	"time"

	"golang.org/x/exp/constraints"
)

// Number is a type constraint that allows only complex, float, and integer types.
type Number interface {
	constraints.Complex | constraints.Float | constraints.Integer
}

// Dedup returns a new slice containing only the unique elements of the input slice.
func Dedup[T comparable](s []T) []T {
	if len(s) == 0 {
		return nil
	}

	n := make([]T, 0, len(s))
	m := make(map[T]bool)

	for _, v := range s {
		if !m[v] {
			n = append(n, v)
			m[v] = true
		}
	}

	return n
}

// Fill fills the entire slice with a given value.
func Fill[T any](s []T, v T) {
	for i := range s {
		s[i] = v
	}
}

// FillRange fills a range [start:end) of the slice with a given value.
func FillRange[T any](s []T, v T, start, end int) {
	s = s[start:end]

	for i := range s {
		s[i] = v
	}
}

// Filter returns a new slice containing only the elements from the original slice that satisfy the given function.
func Filter[T any](s []T, f func(T) bool) []T {
	if len(s) == 0 {
		return nil
	}

	n := make([]T, 0, len(s))

	for _, v := range s {
		if f(v) {
			n = append(n, v)
		}
	}

	return n
}

// Flat flattens a slice of slices into a single slice.
func Flat[T any](s [][]T) []T {
	l := Reduce(s, func(v []T, acc int) int { return len(v) + acc }, 0)
	if l == 0 {
		return nil
	}

	n := make([]T, 0, l)
	for _, v := range s {
		n = append(n, v...)
	}
	return n
}

// ForEach applies the given function to every element in the slice.
func ForEach[T any](s []T, f func(T)) {
	for _, v := range s {
		f(v)
	}
}

// ForEachIndex applies the given function to every element in the slice along with its index.
func ForEachIndex[T any](s []T, f func(v T, i int)) {
	for i, v := range s {
		f(v, i)
	}
}

// Group groups the elements of a slice by applying a given function to each element and using the result as a key in a map.
func Group[A any, B comparable](s []A, f func(A) B) map[B][]A {
	m := make(map[B][]A)

	for _, v := range s {
		key := f(v)
		m[key] = append(m[key], v)
	}

	return m
}

// Map applies the given function to every element in the slice and returns a new slice containing the results.
func Map[A, B any](s []A, f func(A) B) []B {
	if len(s) == 0 {
		return nil
	}

	n := make([]B, len(s))

	for i, v := range s {
		n[i] = f(v)
	}

	return n
}

// Random returns a random element from a slice and a new slice with the randomly selected element removed. If the input slice is empty, it returns a zero value and a nil slice.
func Random[T any](s []T) (T, []T) {
	if len(s) == 0 {
		var zero T
		return zero, nil
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(s))
	return s[r], append(s[:r], s[r+1:]...)
}

// Reduce applies a function to each element of a slice, accumulating the result into an initial value.
func Reduce[A, B any](s []A, f func(v A, acc B) B, init B) B {
	for _, v := range s {
		init = f(v, init)
	}

	return init
}

// ReduceRight applies a function to each element of a slice in reverse order,
// accumulating the result into an initial value.
func ReduceRight[A, B any](s []A, f func(v A, acc B) B, init B) B {
	for i := len(s) - 1; i >= 0; i-- {
		init = f(s[i], init)
	}

	return init
}

// Reverse reverses the elements of a slice in place.
func Reverse[T any](s []T) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// ReverseCopy returns a new slice with the elements in reverse order.
func ReverseCopy[T any](s []T) []T {
	n := make([]T, len(s))

	for i, v := range s {
		n[len(s)-i-1] = v
	}

	return n
}

// Sum returns the sum of all elements in a slice of type T.
func Sum[T Number](s []T) T {
	var sum T
	for _, v := range s {
		sum += v
	}
	return sum
}
