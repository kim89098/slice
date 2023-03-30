package slice_test

import (
	"fmt"
	"testing"

	"github.com/kim89098/slice"
)

func TestChunk(t *testing.T) {
	testCases := []struct {
		s    []int
		i    int
		want [][]int
	}{
		{[]int{1, 2, 3, 4, 5}, 1, [][]int{{1}, {2}, {3}, {4}, {5}}},
		{[]int{1, 2, 3, 4, 5}, 2, [][]int{{1, 2}, {3, 4}, {5}}},
		{[]int{1, 2, 3, 4, 5}, 5, [][]int{{1, 2, 3, 4, 5}}},
		{[]int{1, 2, 3, 4, 5}, 6, [][]int{{1, 2, 3, 4, 5}}},

		{nil, 1, [][]int{}},
	}

	for _, c := range testCases {
		r := slice.Chunk(c.s, c.i)

		if len(r) != len(c.want) {
			t.Errorf("Chunk(%v, %v) = %v, want %v", c.s, c.i, r, c.want)
			continue
		}

		for i, w := range c.want {
			if !slice.Equals(r[i], w) {
				t.Errorf("Chunk(%v, %v) = %v, want %v", c.s, c.i, r, c.want)
				break
			}
		}
	}
}

func TestConcat(t *testing.T) {
	testCases := []struct {
		ss   [][]int
		want []int
	}{
		{[][]int{{1, 2}, {3, 4}, {5}}, []int{1, 2, 3, 4, 5}},
		{[][]int{nil, nil, nil}, []int{}},
		{[][]int{nil, {1}, nil}, []int{1}},
	}

	for _, c := range testCases {
		if r := slice.Concat(c.ss...); !slice.Equals(r, c.want) {
			t.Errorf("Concat(%v) = %v, want %v", c.ss, r, c.want)
		}
	}
}

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

func TestExpand2D(t *testing.T) {
	testCases := []struct {
		s    [][]int
		m, n int
		want [][]int
	}{
		{
			[][]int{
				{1, 2},
			},
			2, 3,
			[][]int{
				{1, 2, 0},
				{0, 0, 0},
			},
		},
		{
			nil,
			2, 3,
			[][]int{
				{0, 0, 0},
				{0, 0, 0},
			},
		},
		{
			[][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11},
			},
			2, 3,
			[][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11},
			},
		},
	}

	for _, c := range testCases {
		r := slice.Expand2D(c.s, c.m, c.n)

		if len(r) != len(c.want) {
			t.Errorf("len(r) = %v, want %v", len(r), len(c.want))
		}

		for i, s := range r {
			if !slice.Equals(s, c.want[i]) {
				t.Errorf("got %v, want %v", s, c.want[i])
			}
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

func TestFilterMap(t *testing.T) {
	testCases := []struct {
		s          []int
		filterFunc func(int) bool
		mapFunc    func(int) string
		want       []string
	}{
		{
			[]int{1, 2, 3},
			func(v int) bool { return v%2 == 0 },
			func(v int) string { return fmt.Sprint(v) },
			[]string{"2"},
		},
		{
			[]int{1, 2, 3},
			func(v int) bool { return false },
			func(v int) string { return fmt.Sprint(v) },
			[]string{},
		},
		{
			[]int{1, 2, 3},
			func(v int) bool { return true },
			func(v int) string { return fmt.Sprint(v) },
			[]string{"1", "2", "3"},
		},
		{
			[]int{},
			func(v int) bool { return false },
			func(v int) string { return fmt.Sprint(v) },
			[]string{},
		},
		{
			nil,
			func(v int) bool { return false },
			func(v int) string { return fmt.Sprint(v) },
			[]string{},
		},
	}

	for _, c := range testCases {
		if r := slice.FilterMap(c.s, c.filterFunc, c.mapFunc); !slice.Equals(r, c.want) {
			t.Errorf("FilterMap(%v, func, func) = %v, want %v", c.s, r, c.want)
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

func TestGroupMap(t *testing.T) {
	testCases := []struct {
		s       []int
		keyFunc func(int) int
		mapFunc func(int) string
		want    map[int][]string
	}{
		{
			[]int{1, 2, 3, 4},
			func(v int) int { return v % 2 },
			func(v int) string { return fmt.Sprint(v) },
			map[int][]string{0: []string{"2", "4"}, 1: []string{"1", "3"}},
		},
		{
			[]int{},
			func(v int) int { return v % 2 },
			func(v int) string { return fmt.Sprint(v) },
			map[int][]string{},
		},
		{
			nil,
			func(v int) int { return v % 2 },
			func(v int) string { return fmt.Sprint(v) },
			map[int][]string{},
		},
	}

	for _, c := range testCases {
		r := slice.GroupMap(c.s, c.keyFunc, c.mapFunc)
		if len(r) != len(c.want) {
			t.Errorf("GroupMap(%v, func, func) = %v, want %v", c.s, r, c.want)
			continue
		}

		for k, v := range c.want {
			if !slice.Equals(r[k], v) {
				t.Errorf("GroupMap(%v, func, func) = %v, want %v", c.s, r, c.want)
				break
			}
		}
	}
}

func TestInsert(t *testing.T) {
	testCases := []struct {
		s    []int
		i    int
		v    int
		want []int
	}{
		{[]int{1, 2, 3}, 0, 4, []int{4, 1, 2, 3}},
		{[]int{1, 2, 3}, 1, 4, []int{1, 4, 2, 3}},
		{[]int{1, 2, 3}, 2, 4, []int{1, 2, 4, 3}},
		{[]int{1, 2, 3}, 3, 4, []int{1, 2, 3, 4}},
		{[]int{1, 2, 3}, 4, 4, []int{1, 2, 3, 4}},
	}

	for _, c := range testCases {
		if r := slice.Insert(c.s, c.i, c.v); !slice.Equals(r, c.want) {
			t.Errorf("got %v, want %v", r, c.want)
		}
	}
}

func TestMake2D(t *testing.T) {
	testCases := []struct {
		m, n int
	}{
		{0, 0},
		{1, 1},
		{2, 3},
	}

	for _, c := range testCases {
		r := slice.Make2D[int](c.m, c.n)

		if len(r) != c.m {
			t.Errorf("len(r) = %v, want %v", len(r), c.m)
		}

		for _, s := range r {
			if len(s) != c.n {
				t.Errorf("len(s) = %v, want %v", len(s), c.n)
			}
		}
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

	r3 := slice.Map([]int(nil), func(v int) int { return v * 2 })
	if !slice.Equals(r3, nil) {
		t.Errorf("got %v, want nil", r3)
	}
}

func TestMove(t *testing.T) {
	testCases := []struct {
		s    []int
		a, b int
		want []int
	}{
		{[]int{1, 2, 3, 4}, 1, 2, []int{1, 3, 2, 4}},
		{[]int{1, 2, 3, 4}, 2, 1, []int{1, 3, 2, 4}},
		{[]int{1, 2, 3, 4}, 0, 3, []int{2, 3, 4, 1}},
		{[]int{1, 2, 3, 4}, 3, 0, []int{4, 1, 2, 3}},
		{[]int{1, 2, 3, 4}, 1, 1, []int{1, 2, 3, 4}},
	}

	for _, c := range testCases {
		slice.Move(c.s, c.a, c.b)
		if !slice.Equals(c.s, c.want) {
			t.Errorf("got %v, want %v", c.s, c.want)
		}
	}
}

func TestNoNil(t *testing.T) {
	if r := slice.NoNil([]int(nil)); r == nil || len(r) != 0 {
		t.Errorf("got %v, want %v", r, []int{})
	}

	if r := slice.NoNil([]int{}); r == nil || len(r) != 0 {
		t.Errorf("got %v, want %v", r, []int{})
	}

	if r := slice.NoNil([]int{1}); !slice.Equals(r, []int{1}) {
		t.Errorf("got %v, want %v", r, []int{1})
	}
}

func TestRandom(t *testing.T) {
	r1, s1 := slice.Random([]int(nil))
	if r1 != 0 || s1 != nil {
		t.Errorf("got %v, %v, want 0, nil", r1, s1)
	}

	r2, s2 := slice.Random([]int{})
	if r2 != 0 || s2 != nil {
		t.Errorf("got %v, %v, want 0, nil", r2, s2)
	}

	c := []int{1, 2, 3, 4, 5}
	for i := 0; i < 100; i++ {
		r, s := slice.Random(c)
		if !slice.EqualsAnyOrder(append(s, r), c) {
			t.Errorf("got %v, %v", r, s)
		}
	}
}

func TestRange(t *testing.T) {
	testCases := []struct {
		start int
		end   int
		want  []int
	}{
		{1, 4, []int{1, 2, 3}},
		{1, 1, []int{}},
		{1, -1, nil},
	}

	for _, c := range testCases {
		if r := slice.Range(c.start, c.end); !slice.Equals(r, c.want) {
			t.Errorf("got %v, want %v", r, c.want)
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

func TestRemove(t *testing.T) {
	testCases := []struct {
		s    []int
		v    int
		want []int
	}{
		{[]int{1, 2, 3}, 1, []int{2, 3}},
		{[]int{1, 2, 3}, 2, []int{1, 3}},
		{[]int{1, 2, 3}, 3, []int{1, 2}},
		{[]int{1, 2, 3}, 4, []int{1, 2, 3}},

		{[]int{}, 4, []int{}},
		{nil, 4, nil},
	}

	for _, c := range testCases {
		if r := slice.Remove(c.s, c.v); !slice.Equals(r, c.want) {
			t.Errorf("got %v, want %v", r, c.want)
		}
	}
}

func TestRemoveFunc(t *testing.T) {
	testCases := []struct {
		s    []int
		f    func(v int) bool
		want []int
	}{
		{[]int{1, 2, 3}, func(v int) bool { return true }, []int{2, 3}},
		{[]int{1, 2, 3}, func(v int) bool { return false }, []int{1, 2, 3}},
		{[]int{1, 2, 3}, func(v int) bool { return v == 2 }, []int{1, 3}},

		{[]int{}, func(v int) bool { return true }, []int{}},
		{nil, func(v int) bool { return true }, nil},
	}

	for _, c := range testCases {
		if r := slice.RemoveFunc(c.s, c.f); !slice.Equals(r, c.want) {
			t.Errorf("got %v, want %v", r, c.want)
		}
	}
}

func TestRemoveIndex(t *testing.T) {
	testCases := []struct {
		s    []int
		i    int
		want []int
	}{
		{[]int{1, 2, 3}, 0, []int{2, 3}},
		{[]int{1, 2, 3}, 1, []int{1, 3}},
		{[]int{1, 2, 3}, 2, []int{1, 2}},
		{[]int{1, 2, 3}, 3, []int{1, 2, 3}},

		{[]int{}, 4, []int{}},
		{nil, 4, nil},
	}

	for _, c := range testCases {
		if r := slice.RemoveIndex(c.s, c.i); !slice.Equals(r, c.want) {
			t.Errorf("got %v, want %v", r, c.want)
		}
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

func TestShuffle(t *testing.T) {
	testCases := []struct {
		s    []int
		want []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
	}

	for _, c := range testCases {
		slice.Shuffle(c.s)
		if !slice.EqualsAnyOrder(c.s, c.want) {
			t.Errorf("got %v, want %v", c.s, c.want)
		}
	}
}

func TestSort(t *testing.T) {
	c1, w1 := []int{1, 3, 2, 4}, []int{1, 2, 3, 4}
	slice.Sort(c1, func(a, b int) bool { return a < b })
	if !slice.Equals(c1, w1) {
		t.Errorf("got %v, want %v", c1, w1)
	}

	type t2 struct {
		a int
	}

	c2, w2 := []t2{{1}, {3}, {2}, {4}}, []t2{{1}, {2}, {3}, {4}}
	slice.Sort(c2, func(a, b t2) bool { return a.a < b.a })
	if !slice.Equals(c2, w2) {
		t.Errorf("got %v, want %v", c2, w2)
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

func TestZip(t *testing.T) {
	testCases := []struct {
		a, b []int
		want []slice.Zipped[int, int]
	}{
		{
			[]int{1, 2, 3},
			[]int{-1, -2, -3},
			[]slice.Zipped[int, int]{{1, -1}, {2, -2}, {3, -3}},
		},
		{
			[]int{1, 2, 3},
			[]int{-1, -2},
			[]slice.Zipped[int, int]{{1, -1}, {2, -2}},
		},
		{
			[]int{1, 2},
			[]int{-1, -2, -3},
			[]slice.Zipped[int, int]{{1, -1}, {2, -2}},
		},
		{
			nil,
			[]int{-1, -2, -3},
			[]slice.Zipped[int, int]{},
		},
		{
			[]int{1, 2, 3},
			nil,
			[]slice.Zipped[int, int]{},
		},
		{
			nil,
			nil,
			[]slice.Zipped[int, int]{},
		},
	}

	for _, c := range testCases {
		if r := slice.Zip(c.a, c.b); !slice.Equals(r, c.want) {
			t.Errorf("Zip(%v, %v) = %v, want %v", c.a, c.b, r, c.want)
		}
	}
}
