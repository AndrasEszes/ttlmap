package ttlmap_test

import (
	"testing"

	"github.com/AndrasEszes/ttlmap"
)

func TestNewFunc(t *testing.T) {
	m := ttlmap.New()

	if m == nil {
		t.Error("ttlmap.New() shouldn't return with nil")
	}

	if _, implements := interface{}(m).(ttlmap.TTLMap); !implements {
		t.Error("ttlmap.New() must return with a ttlmap.TTLMap implementation")
	}
}
