package subsetSum

import "testing"

type tCase struct {
	set  []int
	v    int
	want bool
}

var (
	testCase = []tCase{
		{
			set:  []int{3, 34, 4, 12, 5, 2},
			v:    9,
			want: true,
		},
		{
			set:  []int{3, 34, 4, 12, 5, 2},
			v:    30,
			want: false,
		},
	}
)

func TestRecursive(t *testing.T) {
	for _, v := range testCase {
		got := recursive(v.set, v.v, len(v.set))
		if got != v.want {
			t.Errorf("wanted : %v but got otherwise : %v", v.want, got)
		}
	}
}

func TestDynamic(t *testing.T) {
	for _, v := range testCase {
		got := dynamic(v.set, v.v)
		if got != v.want {
			t.Errorf("wanted : %v but got otherwise : %v", v.want, got)
		}
	}
}
