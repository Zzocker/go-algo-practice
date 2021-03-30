package minNumCoin

import "testing"

type tCase struct {
	c    []int
	v    int
	want int
}

var (
	testCase = []tCase{
		{[]int{25, 10, 5}, 30, 2},
		{[]int{9, 6, 5, 1}, 11, 2},
	}
)

func TestRecursive(t *testing.T) {
	for _, v := range testCase {
		got := recursive(v.c, v.v, len(v.c))
		if got != v.want {
			t.Errorf("wanted : %d but got otherwise : %d", v.want, got)
		}
	}
}

func TestDynamic(t *testing.T) {
	for _, v := range testCase {
		got := dynamic(v.c, v.v)
		if got != v.want {
			t.Errorf("wanted : %d but got otherwise : %d", v.want, got)
		}
	}
}
