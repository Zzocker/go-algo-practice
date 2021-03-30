package coinChange

import "testing"

type tCase struct {
	n    int
	s    []int
	want int
}

var (
	testCase = []tCase{
		{
			n:    4,
			s:    []int{1, 2, 3},
			want: 4,
		},
		{
			n:    10,
			s:    []int{2, 5, 3, 6},
			want: 5,
		},
	}
)

func TestRecursive(t *testing.T) {
	for _, v := range testCase {
		got := recursive(v.s, v.n, len(v.s))
		if got != v.want {
			t.Errorf("wanted : %d but got otherwise : %d", v.want, got)
		}
	}
}

func TestDynamic(t *testing.T) {
	for _, v := range testCase {
		got := dynamic(v.s, v.n)
		if got != v.want {
			t.Errorf("wanted : %d but got otherwise : %d", v.want, got)
		}
	}
}
