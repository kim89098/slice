package slice_test

import (
	"testing"

	"github.com/kim89098/slice"
)

func TestPop(t *testing.T) {
	type wantResult struct {
		v int
		s []int
	}

	testCases := []struct {
		s    []int
		want wantResult
	}{
		{[]int{1, 2, 3}, wantResult{3, []int{1, 2}}},
		{[]int{1, 2, 3, 1}, wantResult{1, []int{1, 2, 3}}},
		{[]int{1}, wantResult{1, []int{}}},
		{[]int{}, wantResult{0, nil}},
		{nil, wantResult{0, nil}},
	}

	for _, c := range testCases {
		if r, s := slice.Pop(c.s); r != c.want.v || !slice.Equals(s, c.want.s) {
			t.Errorf("got %v, %v, want %v, %v", r, s, c.want.v, c.want.s)
		}
	}
}

func TestPush(t *testing.T) {
	testCases := []struct {
		s    []int
		v    int
		want []int
	}{
		{[]int{1, 2, 3}, 1, []int{1, 2, 3, 1}},
		{[]int{1, 2, 3, 1}, 5, []int{1, 2, 3, 1, 5}},
	}

	for _, c := range testCases {
		if r := slice.Push(c.s, c.v); !slice.Equals(r, c.want) {
			t.Errorf("got %v, want %v", r, c.want)
		}
	}
}

func TestShift(t *testing.T) {
	type wantResult struct {
		v int
		s []int
	}

	testCases := []struct {
		s    []int
		want wantResult
	}{
		{[]int{1, 2, 3}, wantResult{1, []int{2, 3}}},
		{[]int{1, 2, 3, 1}, wantResult{1, []int{2, 3, 1}}},
		{[]int{1}, wantResult{1, []int{}}},
		{[]int{}, wantResult{0, nil}},
		{nil, wantResult{0, nil}},
	}

	for _, c := range testCases {
		if r, s := slice.Shift(c.s); r != c.want.v || !slice.Equals(s, c.want.s) {
			t.Errorf("got %v, %v, want %v, %v", r, s, c.want.v, c.want.s)
		}
	}
}

func TestUnshift(t *testing.T) {
	testCases := []struct {
		s    []int
		v    int
		want []int
	}{
		{[]int{1, 2, 3}, 1, []int{1, 1, 2, 3}},
		{[]int{1, 1, 2, 3}, 2, []int{2, 1, 1, 2, 3}},
	}

	for _, c := range testCases {
		if r := slice.Unshift(c.s, c.v); !slice.Equals(r, c.want) {
			t.Errorf("got %v, want %v", r, c.want)
		}
	}
}
