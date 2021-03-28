package catalanNumber

import "testing"

// n, ans
var testCase = [][]int{
	{0, 1},
	{1, 1},
	{2, 2},
	{3, 5},
	{4, 14},
	{5, 42},
	{6, 132},
	{7, 429},
	{8, 1430},
	{9, 4862},
}

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
