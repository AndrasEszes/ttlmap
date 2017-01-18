package ttlmap_test

import (
	"testing"

	"github.com/AndrasEszes/ttlmap"
)

type bt struct {
	a, b string
}

func BenchmarkNewFunc(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		m := ttlmap.New()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			m.Insert(i, i, 0)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		m := ttlmap.New()
		b.ResetTimer()

		b.RunParallel(func(pb *testing.PB) {
			i := 0
			for pb.Next() {
				m.Insert(i, i, 0)
				i++
			}
		})
	})
}
