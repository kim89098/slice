package slice_test

import (
	"testing"

	"github.com/kim89098/slice"
)

func TestKeys(t *testing.T) {
	testCases := []struct {
		m    map[int]string
		want []int
	}{
		{map[int]string{1: "1", 2: "2", 3: "3"}, []int{1, 2, 3}},
		{map[int]string{1: "1"}, []int{1}},
		{map[int]string{}, nil},
		{nil, nil},
	}

	for _, c := range testCases {
		if r := slice.Keys(c.m); !slice.EqualsAnyOrder(r, c.want) {
			t.Errorf("got %v, want %v", r, c.want)
		}
	}
}

func TestValues(t *testing.T) {
	testCases := []struct {
		m    map[int]string
		want []string
	}{
		{map[int]string{1: "1", 2: "2", 3: "3"}, []string{"1", "2", "3"}},
		{map[int]string{1: "1"}, []string{"1"}},
		{map[int]string{}, nil},
		{nil, nil},
	}

	for _, c := range testCases {
		if r := slice.Values(c.m); !slice.EqualsAnyOrder(r, c.want) {
			t.Errorf("got %v, want %v", r, c.want)
		}
	}
}
