package atcoder

import (
	"XCPCBoard/spiders/model"
	"testing"
)

func BenchmarkFlush(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Flush(model.TestAtcIdLQY)
	}
}
