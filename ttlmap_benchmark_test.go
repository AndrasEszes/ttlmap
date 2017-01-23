package ttlmap_test

import (
	"testing"

	"github.com/AndrasEszes/ttlmap"
)

type bt struct {
	a, b string
}

func BenchmarkInsertFunc(b *testing.B) {
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

func BenchmarkGetFunc(b *testing.B) {
	m := ttlmap.New()
	for i := 0; i < 10; i++ {
		if err := m.Insert(i, i, ttlmap.Never); err != nil {
			b.FailNow()
		}
	}

	b.Run("sequential-hit", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			if _, err := m.Get(5); err != nil {
				b.Error(err)
			}
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				if _, err := m.Get(5); err != nil {
					b.Error(err)
				}
			}
		})
	})
}
