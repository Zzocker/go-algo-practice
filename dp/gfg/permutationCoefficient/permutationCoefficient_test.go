package permutationCoefficient

import "testing"

var (
	testCase = [][]int{
		// n , k ,want
		{10, 2, 90},
		{10, 3, 720},
		{10, 0, 1},
		{10, 1, 10},
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

func TestDynamic(t *testing.T) {
	for _, v := range testCase {
		got := dynamic(v[0], v[1])
		if got != v[2] {
			t.Errorf("wanted : %d but got otherwise : %d", v[2], got)
		}
	}
}
