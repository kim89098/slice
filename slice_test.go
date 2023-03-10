package slice_test

import (
	"fmt"
	"testing"

	"github.com/kim89098/slice"
)

func TestDedup(t *testing.T) {
	testCases := []struct {
		s    []int
		want []int
	}{
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{1, 2, 3, 3}, []int{1, 2, 3}},
		{[]int{1, 2, 2, 3}, []int{1, 2, 3}},
		{[]int{1}, []int{1}},
		{[]int{}, nil},
		{nil, nil},
	}

	for _, c := range testCases {
		if r := slice.Dedup(c.s); !slice.Equals(r, c.want) {
			t.Errorf("got %v, want %v", r, c.want)
		}
	}
}

func TestFill(t *testing.T) {
	testCases := []struct {
		s    []int
		v    int
		want []int
	}{
		{[]int{1, 2, 3}, 2, []int{2, 2, 2}},
		{[]int{}, 0, []int{}},
	}

	for _, c := range testCases {
		slice.Fill(c.s, c.v)
		if !slice.Equals(c.s, c.want) {
			t.Errorf("got %v, want %v", c.s, c.want)
		}
	}
}

func TestFillRange(t *testing.T) {
	testCases := []struct {
		s          []int
		v          int
		start, end int
		want       []int
	}{
		{[]int{1, 2, 3}, 4, 1, 1, []int{1, 2, 3}},
		{[]int{1, 2, 3}, 4, 1, 2, []int{1, 4, 3}},
		{[]int{1, 2, 3}, 4, 1, 3, []int{1, 4, 4}},
	}

	for _, c := range testCases {
		slice.FillRange(c.s, c.v, c.start, c.end)
		if !slice.Equals(c.s, c.want) {
			t.Errorf("got %v, want %v", c.s, c.want)
		}
	}
}

func TestFilter(t *testing.T) {
	testCases := []struct {
		s    []int
		f    func(int) bool
		want []int
	}{
		{[]int{1, 2, 3}, func(v int) bool { return true }, []int{1, 2, 3}},
		{[]int{1, 2, 3}, func(v int) bool { return false }, []int{}},
		{[]int{1, 2, 3}, func(v int) bool { return v > 1 }, []int{2, 3}},
		{[]int{}, func(v int) bool { return true }, nil},
		{nil, func(v int) bool { return true }, nil},
	}

	for _, c := range testCases {
		if r := slice.Filter(c.s, c.f); !slice.Equals(r, c.want) {
			t.Errorf("got %v, want %v", r, c.want)
		}
	}
}

func TestFlat(t *testing.T) {
	testCases := []struct {
		s    [][]int
		want []int
	}{
		{[][]int{{1}, {2}, {3}}, []int{1, 2, 3}},
		{[][]int{{1, 2}, {3}}, []int{1, 2, 3}},
		{[][]int{{1, 2, 3}}, []int{1, 2, 3}},
		{nil, nil},
		{[][]int{nil}, nil},
		{[][]int{nil, nil}, nil},
	}

	for _, c := range testCases {
		if r := slice.Flat(c.s); !slice.Equals(r, c.want) {
			t.Errorf("got %v, want %v", r, c.want)
		}
	}
}

func TestForEach(t *testing.T) {
	var buf int
	slice.ForEach([]int{1, 2, 3}, func(v int) { buf += v })
	if buf != 6 {
		t.Errorf("got %v, want 6", buf)
	}
}

func TestForEachIndex(t *testing.T) {
	var buf int
	slice.ForEachIndex([]int{1, 2, 3}, func(v int, i int) { buf += v * i })
	if buf != 8 {
		t.Errorf("got %v, want 8", buf)
	}
}

func TestGroup(t *testing.T) {
	r1 := slice.Group([]int{1, 2, 3}, func(v int) int { return v % 2 })
	if !slice.Equals(r1[0], []int{2}) {
		t.Errorf("got %v, want [2]", r1[0])
	}
	if !slice.Equals(r1[1], []int{1, 3}) {
		t.Errorf("got %v, want [1,3]", r1[1])
	}

	r2 := slice.Group([]int{1, 2, 3}, func(v int) bool { return v%2 == 0 })
	if !slice.Equals(r2[true], []int{2}) {
		t.Errorf("got %v, want [2]", r2[true])
	}
	if !slice.Equals(r2[false], []int{1, 3}) {
		t.Errorf("got %v, want [1,3]", r2[false])
	}
}

func TestMap(t *testing.T) {
	r1 := slice.Map([]int{1, 2, 3}, func(v int) int { return v * 2 })
	if !slice.Equals(r1, []int{2, 4, 6}) {
		t.Errorf("got %v, want []int{2, 4, 6}", r1)
	}

	r2 := slice.Map([]int{1, 2, 3}, func(v int) string { return fmt.Sprint(v) })
	if !slice.Equals(r2, []string{"1", "2", "3"}) {
		t.Errorf("got %v, want []string{2, 4, 6}", r2)
	}

	r3 := slice.Map(nil, func(v int) int { return v * 2 })
	if !slice.Equals(r3, nil) {
		t.Errorf("got %v, want nil", r3)
	}
}

func TestRandom(t *testing.T) {
	r1, s1 := slice.Random[int](nil)
	if r1 != 0 || s1 != nil {
		t.Errorf("got %v, %v, want 0, nil", r1, s1)
	}

	r2, s2 := slice.Random([]int{})
	if r2 != 0 || s2 != nil {
		t.Errorf("got %v, %v, want 0, nil", r1, s1)
	}

	c := []int{1, 2, 3, 4, 5}
	for i := 0; i < 100; i++ {
		r, s := slice.Random(c)
		if !slice.EqualsAnyOrder(append(s, r), c) {
			t.Errorf("got %v, %v", r, s)
		}
	}
}

func TestReduce(t *testing.T) {
	r1 := slice.Reduce([]int{1, 2, 3}, func(v int, acc int) int { return v + acc }, 0)
	if r1 != 6 {
		t.Errorf("got %v, want 6", r1)
	}

	r2 := slice.Reduce([]int{1, 2, 3}, func(v int, acc string) string { return fmt.Sprint(v) + acc }, "")
	if r2 != "321" {
		t.Errorf("got %v, want 321", r2)
	}
}

func TestReduceRight(t *testing.T) {
	r1 := slice.ReduceRight([]int{1, 2, 3}, func(v int, acc int) int { return v + acc }, 0)
	if r1 != 6 {
		t.Errorf("got %v, want 6", r1)
	}

	r2 := slice.ReduceRight([]int{1, 2, 3}, func(v int, acc string) string { return fmt.Sprint(v) + acc }, "")
	if r2 != "123" {
		t.Errorf("got %v, want 123", r2)
	}
}

func TestReverse(t *testing.T) {
	testCases := []struct {
		s    []int
		want []int
	}{
		{[]int{1, 2, 3}, []int{3, 2, 1}},
		{[]int{1, 2}, []int{2, 1}},
		{[]int{1}, []int{1}},
		{[]int{}, []int{}},
	}

	for _, c := range testCases {
		if slice.Reverse(c.s); !slice.Equals(c.s, c.want) {
			t.Errorf("got %v, want %v", c.s, c.want)
		}
	}
}

func TestReverseCopy(t *testing.T) {
	testCases := []struct {
		s    []int
		want []int
	}{
		{[]int{1, 2, 3}, []int{3, 2, 1}},
		{[]int{1, 2}, []int{2, 1}},
		{[]int{1}, []int{1}},
		{[]int{}, []int{}},
	}

	for _, c := range testCases {
		if r := slice.ReverseCopy(c.s); !slice.Equals(r, c.want) {
			t.Errorf("got %v, want %v", r, c.want)
		}
	}
}

func TestSum(t *testing.T) {
	testCases := []struct {
		s    []int
		want int
	}{
		{[]int{1, 2, 3}, 6},
		{[]int{1}, 1},
		{[]int{}, 0},
		{nil, 0},
	}

	for _, c := range testCases {
		if r := slice.Sum(c.s); r != c.want {
			t.Errorf("got %v, want %v", r, c.want)
		}
	}
}
