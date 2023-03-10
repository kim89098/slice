package slice_test

import (
	"testing"

	"github.com/kim89098/slice"
)

func TestEquals(t *testing.T) {
	testCases := []struct {
		a, b []int
		want bool
	}{
		{[]int{1, 2, 3}, []int{1, 2, 3}, true},
		{[]int{1, 2, 3}, []int{3, 2, 1}, false},

		{[]int{}, []int{}, true},
		{[]int{}, []int{1}, false},
		{[]int{1}, []int{}, false},

		{[]int{1}, []int{1, 2}, false},
		{[]int{1, 2}, []int{1}, false},
	}

	for _, c := range testCases {
		if r := slice.Equals(c.a, c.b); r != c.want {
			t.Errorf("got %v, want %v", r, c.want)
		}
	}
}

func TestEqualsAnyOrder(t *testing.T) {
	testCases := []struct {
		a, b []int
		want bool
	}{
		{[]int{1, 2, 3}, []int{1, 2, 3}, true},
		{[]int{1, 2, 3}, []int{3, 2, 1}, true},
		{[]int{1, 2, 3}, []int{1, 2, 4}, false},

		{[]int{}, []int{}, true},
		{[]int{}, []int{1}, false},
		{[]int{1}, []int{}, false},

		{[]int{1}, []int{1, 2}, false},
		{[]int{1, 2}, []int{1}, false},

		{[]int{1, 2, 3, 3}, []int{1, 2, 2, 3}, false},
	}

	for _, c := range testCases {
		if r := slice.EqualsAnyOrder(c.a, c.b); r != c.want {
			t.Errorf("got %v, want %v", r, c.want)
		}
	}
}

func TestEvery(t *testing.T) {
	testCases := []struct {
		s    []int
		f    func(int) bool
		want bool
	}{
		{[]int{1, 2, 3}, func(v int) bool { return v > 0 }, true},
		{[]int{1, 2, 3}, func(v int) bool { return v < 0 }, false},
		{[]int{1, 2, 3}, func(v int) bool { return v > 1 }, false},
		{[]int{}, func(v int) bool { return false }, true},
		{nil, func(v int) bool { return false }, true},
	}

	for _, c := range testCases {
		if r := slice.Every(c.s, c.f); r != c.want {
			t.Errorf("got %v, want %v", r, c.want)
		}
	}
}

func TestFind(t *testing.T) {
	type wantResult struct {
		v  int
		ok bool
	}

	testCases := []struct {
		s    []int
		f    func(int) bool
		want wantResult
	}{
		{[]int{1, 2, 3}, func(v int) bool { return true }, wantResult{1, true}},
		{[]int{1, 2, 3}, func(v int) bool { return false }, wantResult{0, false}},
		{[]int{1, 2, 3}, func(v int) bool { return v > 1 }, wantResult{2, true}},
	}

	for _, c := range testCases {
		if r, ok := slice.Find(c.s, c.f); r != c.want.v || ok != c.want.ok {
			t.Errorf("got %v, %v, want %v, %v", r, ok, c.want.v, c.want.ok)
		}
	}
}

func TestFindDefault(t *testing.T) {
	testCases := []struct {
		s            []int
		f            func(int) bool
		defaultValue int
		want         int
	}{
		{[]int{1, 2, 3}, func(v int) bool { return true }, 4, 1},
		{[]int{1, 2, 3}, func(v int) bool { return false }, 4, 4},
		{[]int{1, 2, 3}, func(v int) bool { return v > 1 }, 4, 2},
	}

	for _, c := range testCases {
		if r := slice.FindDefault(c.s, c.f, c.defaultValue); r != c.want {
			t.Errorf("got %v, want %v", r, c.want)
		}
	}
}

func TestFindIndex(t *testing.T) {
	testCases := []struct {
		s    []int
		f    func(int) bool
		want int
	}{
		{[]int{1, 2, 3}, func(v int) bool { return true }, 0},
		{[]int{1, 2, 3}, func(v int) bool { return false }, -1},
		{[]int{1, 2, 3}, func(v int) bool { return v > 1 }, 1},
	}

	for _, c := range testCases {
		if r := slice.FindIndex(c.s, c.f); r != c.want {
			t.Errorf("got %v, want %v", r, c.want)
		}
	}
}

func TestFindLast(t *testing.T) {
	type wantResult struct {
		v  int
		ok bool
	}

	testCases := []struct {
		s    []int
		f    func(int) bool
		want wantResult
	}{
		{[]int{1, 2, 3}, func(v int) bool { return true }, wantResult{3, true}},
		{[]int{1, 2, 3}, func(v int) bool { return false }, wantResult{0, false}},
		{[]int{1, 2, 3}, func(v int) bool { return v < 3 }, wantResult{2, true}},
		{[]int{1, 2, 3}, func(v int) bool { return v == 1 }, wantResult{1, true}},
	}

	for _, c := range testCases {
		if r, ok := slice.FindLast(c.s, c.f); r != c.want.v || ok != c.want.ok {
			t.Errorf("got %v, %v, want %v, %v", r, ok, c.want.v, c.want.ok)
		}
	}
}

func TestFindLastDefault(t *testing.T) {
	testCases := []struct {
		s            []int
		f            func(int) bool
		defaultValue int
		want         int
	}{
		{[]int{1, 2, 3}, func(v int) bool { return true }, 4, 3},
		{[]int{1, 2, 3}, func(v int) bool { return false }, 4, 4},
		{[]int{1, 2, 3}, func(v int) bool { return v < 3 }, 4, 2},
		{[]int{1, 2, 3}, func(v int) bool { return v == 1 }, 4, 1},
	}

	for _, c := range testCases {
		if r := slice.FindLastDefault(c.s, c.f, c.defaultValue); r != c.want {
			t.Errorf("got %v, want %v", r, c.want)
		}
	}
}

func TestFindLastIndex(t *testing.T) {
	testCases := []struct {
		s    []int
		f    func(int) bool
		want int
	}{
		{[]int{1, 2, 3}, func(v int) bool { return true }, 2},
		{[]int{1, 2, 3}, func(v int) bool { return false }, -1},
		{[]int{1, 2, 3}, func(v int) bool { return v < 3 }, 1},
		{[]int{1, 2, 3}, func(v int) bool { return v == 1 }, 0},
	}

	for _, c := range testCases {
		if r := slice.FindLastIndex(c.s, c.f); r != c.want {
			t.Errorf("got %v, want %v", r, c.want)
		}
	}
}

func TestIncludes(t *testing.T) {
	testCases := []struct {
		s    []int
		v    int
		want bool
	}{
		{[]int{1, 2, 3}, 1, true},
		{[]int{1, 2, 3}, 0, false},
	}

	for _, c := range testCases {
		if r := slice.Includes(c.s, c.v); r != c.want {
			t.Errorf("got %v, want %v", r, c.want)
		}
	}
}

func TestIndexOf(t *testing.T) {
	testCases := []struct {
		s    []int
		v    int
		want int
	}{
		{[]int{1, 2, 3}, 1, 0},
		{[]int{1, 2, 3}, 2, 1},
		{[]int{1, 2, 3}, 3, 2},
		{[]int{1, 2, 3}, 0, -1},
	}

	for _, c := range testCases {
		if r := slice.IndexOf(c.s, c.v); r != c.want {
			t.Errorf("got %v, want %v", r, c.want)
		}
	}
}

func TestIndexOfFrom(t *testing.T) {
	testCases := []struct {
		s    []int
		v    int
		from int
		want int
	}{
		{[]int{1, 2, 3}, 1, 1, -1},
		{[]int{1, 2, 3}, 2, 1, 1},
		{[]int{1, 2, 3}, 3, 1, 2},
		{[]int{1, 2, 3}, 0, 0, -1},
	}

	for _, c := range testCases {
		if r := slice.IndexOfFrom(c.s, c.v, c.from); r != c.want {
			t.Errorf("got %v, want %v", r, c.want)
		}
	}
}

func TestLastIndexOf(t *testing.T) {
	testCases := []struct {
		s    []int
		v    int
		want int
	}{
		{[]int{1, 2, 3, 1}, 1, 3},
		{[]int{1, 2, 3, 1}, 2, 1},
		{[]int{1, 2, 3, 1}, 3, 2},
		{[]int{1, 2, 3, 1}, 0, -1},
	}

	for _, c := range testCases {
		if r := slice.LastIndexOf(c.s, c.v); r != c.want {
			t.Errorf("got %v, want %v", r, c.want)
		}
	}
}

func TestMost(t *testing.T) {
	testCases := []struct {
		s    []int
		f    func(v, most int) bool
		want int
	}{
		{[]int{1, 2, 3}, func(v, most int) bool { return v-most > 0 }, 3},
		{[]int{1, 2, 3}, func(v, most int) bool { return v-most < 0 }, 1},
		{[]int{1, 2, 3}, func(v, most int) bool { return v%2 == 0 }, 2},
		{[]int{1, 2, 3}, func(v, most int) bool { return true }, 3},
		{[]int{1, 2, 3}, func(v, most int) bool { return false }, 1},
		{[]int{}, func(v, most int) bool { return true }, 0},
		{nil, func(v, most int) bool { return true }, 0},
	}

	for _, c := range testCases {
		if r := slice.Most(c.s, c.f); r != c.want {
			t.Errorf("got %v, want %v", r, c.want)
		}
	}
}

func TestSome(t *testing.T) {
	testCases := []struct {
		s    []int
		f    func(int) bool
		want bool
	}{
		{[]int{1, 2, 3}, func(v int) bool { return true }, true},
		{[]int{1, 2, 3}, func(v int) bool { return false }, false},

		{[]int{}, func(v int) bool { return true }, false},
		{[]int{}, func(v int) bool { return false }, false},

		{[]int{1, 2, 3}, func(v int) bool { return v > 2 }, true},
		{[]int{1, 2, 3}, func(v int) bool { return v < 0 }, false},
	}

	for _, c := range testCases {
		if r := slice.Some(c.s, c.f); r != c.want {
			t.Errorf("got %v, want %v", r, c.want)
		}
	}
}
