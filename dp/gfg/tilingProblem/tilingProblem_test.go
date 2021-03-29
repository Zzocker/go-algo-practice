package tilingProblem

import "testing"

var (
	testCase = [][]int{
		// n,want
		{4, 3},
		{3, 2},
	}
)

func TestRecursive(t *testing.T) {
	for _, v := range testCase {
		got := recursive(v[0])
		if got != v[1] {
			t.Errorf("wanted : %d but got otherwise : %d", v[1], got)
		}
	}
}

func TestDynamic(t *testing.T) {
	for _, v := range testCase {
		got := dynamic(v[0])
		if got != v[1] {
			t.Errorf("wanted : %d but got otherwise : %d", v[1], got)
		}
	}
}
