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

// Chunk returns a new slice of slices where each slice contains at most the given size number of elements from the original slice.
func Chunk[S ~[]E, E any](slice S, size int) []S {
	n := (len(slice) + size - 1) / size

	r := make([]S, n)
	for i := range r {
		start := i * size
		end := (i + 1) * size
		if end > len(slice) {
			end = len(slice)
		}
		r[i] = slice[start:end]
	}
	return r
}

// Clone returns a copy of the input slice.
func Clone[S ~[]E, E any](s S) S {
	n := make(S, len(s))
	copy(n, s)
	return n
}

// Concat returns a new slice containing all the elements of the input slices in order.
func Concat[S ~[]E, E any](ss ...S) S {
	var totalLen int
	for _, s := range ss {
		totalLen += len(s)
	}

	r := make(S, 0, totalLen)
	for _, s := range ss {
		r = append(r, s...)
	}
	return r
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

// Expand2D expands a 2D slice ss to m rows and n columns, adding elements to any short rows and appending any short columns.
func Expand2D[SS ~[]S, S ~[]E, E any](ss SS, m, n int) SS {
	var zeros S

	for i, s := range ss {
		if l := len(s); l < n {
			if zeros == nil {
				zeros = make(S, n)
			}
			ss[i] = append(s, zeros[:n-l]...)
		}
	}

	for len(ss) < m {
		ss = append(ss, make(S, n))
	}

	return ss
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

// FilterMap returns a new slice containing the elements of the input slice that satisfy filterFunc, transformed by mapFunc.
func FilterMap[S ~[]E, E, R any](s S, filterFunc func(E) bool, mapFunc func(E) R) []R {
	n := make([]R, 0, len(s))
	for _, v := range s {
		if filterFunc(v) {
			n = append(n, mapFunc(v))
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

// GroupMap groups the elements of a slice by applying a key function and a map function to each element,
// using the result of the key function as a key in a map and the result of the map function as a value in a slice.
func GroupMap[S ~[]E, E any, K comparable, V any](s S, keyFunc func(E) K, mapFunc func(E) V) map[K][]V {
	m := make(map[K][]V)

	for _, v := range s {
		key := keyFunc(v)
		m[key] = append(m[key], mapFunc(v))
	}

	return m
}

// Insert inserts the element v into slice s at the given index i.
func Insert[S ~[]E, E any](s S, i int, v E) S {
	if i >= len(s) {
		return append(s, v)
	}

	return append(s[:i], append(S{v}, s[i:]...)...)
}

// Make2D returns a new 2D slice with m rows and n columns.
func Make2D[T any](m, n int) [][]T {
	if m == 0 {
		return nil
	}

	mat := make([][]T, m)

	if n == 0 {
		return mat
	}

	mem := make([]T, m*n)

	start, end := 0, n

	for i := range mat {
		mat[i] = mem[start:end:end]
		start += n
		end += n
	}

	return mat
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

// Move moves the element at index a to index b in slice s.
func Move[S ~[]E, E any](s S, a, b int) {
	if a == b {
		return
	}

	v := s[a]

	if a < b {
		copy(s[a:], s[a+1:b+1])
	} else {
		copy(s[b+1:], s[b:a])
	}

	s[b] = v
}

// NoNil returns the input slice with a non-nil value. If the input slice is nil,
// it returns a new empty slice of the same type.
func NoNil[S ~[]E, E any](s S) S {
	if s == nil {
		return S{}
	}
	return s
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

// Range returns a slice of integers from start (inclusive) to end (exclusive).
// If start is greater than or equal to end, an empty slice is returned.
func Range[T constraints.Integer](start, end T) []T {
	if start > end {
		return nil
	}

	s := make([]T, 0, end-start)
	for i := start; i < end; i++ {
		s = append(s, i)
	}
	return s
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

// Remove returns a new slice with the first occurrence of v removed from s.
// If v is not found in s, Remove returns s unchanged.
func Remove[S ~[]E, E comparable](s S, v E) S {
	return RemoveIndex(s, IndexOf(s, v))
}

// RemoveFunc returns a new slice with the first element e in s for which f(e) is true removed.
// If no such element is found, RemoveFunc returns s unchanged.
func RemoveFunc[S ~[]E, E any](s S, f func(E) bool) S {
	return RemoveIndex(s, FindIndex(s, f))
}

// RemoveIndex returns a new slice with the element at index i removed from s.
// If i is out of bounds for s, RemoveIndex returns s unchanged.
func RemoveIndex[S ~[]E, E any](s S, i int) S {
	if i < 0 || i >= len(s) {
		return s
	}

	return append(s[:i], s[i+1:]...)
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

// Shuffle randomizes the order of elements in the given slice using rand.Shuffle.
// Note that the function modifies the original slice, and does not return a new one.
func Shuffle[S ~[]E, E any](s S) {
	rand.Shuffle(len(s), func(i, j int) {
		s[i], s[j] = s[j], s[i]
	})
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

type Zipped[A, B any] struct {
	A A
	B B
}

// Zip returns a new slice of pairs where the i-th pair contains the i-th elements of each of the input slices.
// If the input slices have different lengths, the resulting slice will have length equal to the length of the shortest input slice.
func Zip[SA ~[]A, SB ~[]B, A, B any](sa SA, sb SB) []Zipped[A, B] {
	n := len(sa)
	if len(sb) < n {
		n = len(sb)
	}

	r := make([]Zipped[A, B], n)
	for i := 0; i < n; i++ {
		r[i] = Zipped[A, B]{sa[i], sb[i]}
	}
	return r
}
