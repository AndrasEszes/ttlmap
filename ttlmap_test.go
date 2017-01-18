package ttlmap_test

import (
	"sync"
	"testing"
	"time"

	"github.com/AndrasEszes/ttlmap"
)

type s struct {
	a string
	b string
}

func TestNewFunc(t *testing.T) {
	m := ttlmap.New()

	if m == nil {
		t.Error("ttlmap.New() shouldn't return with nil")
	}

	if _, implements := interface{}(m).(ttlmap.TTLMap); !implements {
		t.Error("ttlmap.New() must return with a ttlmap.TTLMap implementation")
	}
}

func TestInsertMethod(t *testing.T) {
	tests := []struct {
		nam string
		key interface{}
		val interface{}
		exp time.Duration
		err error
	}{
		{"nil-key", nil, nil, 0, ttlmap.ErrNilKeyIsNotAcceptable},
		{"int-int", 1, 1, 0, nil},
		{"string-string", "hello", "world", 0, nil},
		{"struct-struct", s{"a", "b"}, s{"c", "d"}, 0, nil},
	}

	m := ttlmap.New()
	w := &sync.WaitGroup{}

	w.Add(len(tests))
	for _, test := range tests {
		test := test
		t.Run(test.nam, func(t *testing.T) {
			defer w.Done()
			if act := m.Insert(test.key, test.val, test.exp); act != test.err {
				t.Errorf("got \"%v\", but it should \"%v\"", act, test.err)
			}
		})
	}
	w.Wait()

	t.Run("key-exists", func(t *testing.T) {
		if err := m.Insert("hello", "world", 0); err != ttlmap.ErrKeyAlreadyExists {
			t.Errorf("got \"%v\", but it should \"%v\"", err, ttlmap.ErrKeyAlreadyExists)
		}
	})
}
