package binomialCoefficient

import "testing"

var (
	testCase = [][]int{
		// n,k,want
		{4, 2, 6},
		{5, 2, 10},
	}
)

func TestRecursive(t *testing.T) {
	for _, v := range testCase {
		got := recursive(v[0], v[1])
		if got != v[2] {
			t.Errorf("wanted : %d but got otherwise : %d", v[2], got)
		}
	}
}

func TestMemoized(t *testing.T) {
	for _, v := range testCase {
		got := memoized(v[0], v[1])
		if got != v[2] {
			t.Errorf("wanted : %d but got otherwise : %d", v[2], got)
		}
	}
}

func TestDynamic(t *testing.T) {
	for _, v := range testCase {
		got := dynamic(v[0], v[1])
		if got != v[2] {
			t.Errorf("wanted : %d but got otherwise : %d", v[2], got)
		}
	}
}
