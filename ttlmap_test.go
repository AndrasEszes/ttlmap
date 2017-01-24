package ttlmap_test

import (
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
		idx string
		key interface{}
		val interface{}
		exp time.Duration
		err error
	}{
		{"nil-key", nil, nil, ttlmap.Never, ttlmap.ErrNilKeyIsNotAcceptable},
		{"int-int", 1, 1, ttlmap.Never, nil},
		{"string-string", "hello", "world", ttlmap.Never, nil},
		{"struct-struct", s{"a", "b"}, s{"c", "d"}, time.Second, nil},
		{"key-exists", s{"a", "b"}, "HELLO", ttlmap.Never, ttlmap.ErrKeyAlreadyExists},
		{"non-expired", "nonex", "pired", ttlmap.Never, nil},
	}

	m := ttlmap.New()

	for _, test := range tests {
		test := test
		t.Run(test.idx, func(t *testing.T) {
			if err := m.Insert(test.key, test.val, test.exp); err != test.err {
				t.Errorf("got \"%v\", but it should \"%v\"", err, test.err)
			}
		})
	}
}

func TestUpdateMethod(t *testing.T) {
	fixtures := []struct {
		key interface{}
		val interface{}
		exp time.Duration
	}{
		{"key", "val", ttlmap.Never},
	}

	m := ttlmap.New()
	for _, f := range fixtures {
		if err := m.Insert(f.key, f.val, f.exp); err != nil {
			t.Error(err.Error())
		}
	}

	tests := []struct {
		idx string
		key interface{}
		val interface{}
		exp time.Duration
		err error
	}{
		{"key-exists", "key", "val2", ttlmap.Never, nil},
		{"key-is-nil", nil, "val2", ttlmap.Never, ttlmap.ErrNilKeyIsNotAcceptable},
		{"key-not-exists", "key2", "val2", ttlmap.Never, ttlmap.ErrItemNotFound},
	}

	for _, test := range tests {
		test := test
		t.Run(test.idx, func(t *testing.T) {
			if err := m.Update(test.key, test.val, test.exp); err != test.err {
				t.Errorf("got \"%v\", but it should \"%v\"", err, test.err)
			}
		})
	}
}

func TestHasMethod(t *testing.T) {
	m := ttlmap.New()
	m.Insert("key", "lel", ttlmap.Never)

	if m.Has("lel") {
		t.Error("got true, when should got false")
	}

	if !m.Has("key") {
		t.Error("got false, when should got true")
	}
}

func TestGetMethod(t *testing.T) {
	tests := []struct {
		idx string
		key interface{}
		val interface{}
		exp time.Duration
		err error
	}{
		{"nil-key", nil, nil, ttlmap.Never, ttlmap.ErrItemNotFound},
		{"never-expired", "key1", "val1", ttlmap.Never, nil},
		{"expired", "key2", nil, -time.Second, ttlmap.ErrItemIsExpired},
		{"not-expired", "key3", "val3", time.Second, nil},
	}

	m := ttlmap.New()
	for _, test := range tests {
		test := test
		t.Run(test.idx, func(t *testing.T) {
			m.Insert(test.key, test.val, test.exp)
			val, err := m.Get(test.key)

			if val != test.val {
				t.Errorf("got \"%v\" value, but it should \"%v\"", val, test.val)
			}

			if err != test.err {
				t.Errorf("got \"%v\" error, but it should \"%v\"", err, test.err)
			}
		})
	}
}

func TestRemoveMethod(t *testing.T) {
	m := ttlmap.New()

	fixtures := []struct {
		key, val interface{}
		ttl      time.Duration
	}{
		{s{"a", "b"}, s{"c", "d"}, 0},
		{s{"e", "f"}, s{"g", "h"}, ttlmap.Never},
	}

	for _, f := range fixtures {
		if err := m.Insert(f.key, f.val, f.ttl); err != nil {
			t.Error(err)
		}

		if !m.Has(f.key) {
			t.Errorf("%v not inserted", f.key)
		}

		if err := m.Remove(f.key); err != nil {
			t.Error(err)
		}

		if m.Has(f.key) {
			t.Errorf("%v not removed", f.key)
		}
	}

	if err := m.Remove(nil); err == nil {
		t.Error("should got", ttlmap.ErrNilKeyIsNotAcceptable, "error")
	}
}
