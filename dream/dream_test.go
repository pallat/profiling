package dream

import (
	"testing"
)

func BenchmarkNonsense(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Nonsense(1000)
	}
}
