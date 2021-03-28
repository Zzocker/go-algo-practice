package fibonacci

import (
	"testing"
)

const (
	testNumber = 40
)

/**

	BenchmarkRecursive-4           2         667730850 ns/op               4 B/op          0 allocs/op

	BenchmarkMemoized-4      3249326               366 ns/op             352 B/op          1 allocs/op

	BenchmarkTabulated-4     8012648               156 ns/op             352 B/op          1 allocs/op

**/

func BenchmarkRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		recursive(testNumber)
	}
}

func BenchmarkMemoized(b *testing.B) {
	for i := 0; i < b.N; i++ {
		memoized(testNumber)
	}
}

func BenchmarkTabulated(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tabulated(testNumber)
	}
}
