package ttlmap_test

import (
	"testing"

	"github.com/AndrasEszes/ttlmap"
)

func TestNewFunc(t *testing.T) {
	if ttlmap.New() != nil {
		t.Error("ttlmap.New() must return with nil currently")
	}
}
