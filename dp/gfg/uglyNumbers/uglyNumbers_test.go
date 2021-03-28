package uglyNumbers

import (
	"testing"
)

// n,want
var testcase = [][]int{
	{7, 8},
	{10, 12},
	{15, 24},
	{150, 5832},
}

func TestSimple(t *testing.T) {
	for _, v := range testcase {
		got := simple(v[0])
		if got != v[1] {
			t.Errorf("wanted : %d but got otherwise : %d", v[1], got)
		}
	}
}

func TestTabluted(t *testing.T) {
	for _, v := range testcase {
		got := tabulated(v[0])
		if got != v[1] {
			t.Errorf("wanted : %d but got otherwise : %d", v[1], got)
		}
	}
}
