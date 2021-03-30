package goldMine

import "testing"

type tCase struct {
	mine [][]int
	want int
}

var (
	// testnumber,mine,want
	testCase = []tCase{
		{
			mine: [][]int{
				{1, 3, 3},
				{2, 1, 4},
				{0, 6, 4},
			},
			want: 12,
		},
		{
			mine: [][]int{
				{1, 3, 1, 5},
				{2, 2, 4, 1},
				{5, 0, 2, 3},
				{0, 6, 1, 2},
			},
			want: 16,
		},
		{
			mine: [][]int{
				{10, 33, 13, 15},
				{22, 21, 04, 1},
				{5, 0, 2, 3},
				{0, 6, 14, 2},
			},
			want: 83,
		},
	}
)

func TestRecursive(t *testing.T) {
	for _, v := range testCase {
		got := recursive(v.mine)
		if got != v.want {
			t.Errorf("wanted : %d but got otherwise : %d", v.want, got)
		}
	}
}

func TestDynamic(t *testing.T) {
	for _, v := range testCase {
		got := dynamic(v.mine)
		if got != v.want {
			t.Errorf("wanted : %d but got otherwise : %d", v.want, got)
		}
	}
}
