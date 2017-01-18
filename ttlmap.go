package ttlmap

import "time"

var (
	// ErrKeyExists happens, when try to add an item with a key that already
	// exists in the map.
	ErrKeyExists error

	// ErrItemNotFound happens, when try to update an item, but it not
	// found by key.
	ErrItemNotFound error

	// ErrItemIsExpired happens, when the item is exists in the map, but already
	// expired (gc is not removed yet)
	ErrItemIsExpired error
)

// TTLMap is the main public interface type of package. A TTLMap contains a map
// with elements which has expiration time.
type TTLMap interface {
	// Insert a new element to the map. If the key is exists, return with an
	// "ErrKeyExists" error.
	// If the expiration is "nil" the item is never expired.
	Insert(key, value interface{}, expiration time.Duration) error

	// Update is an existing item's value and expiration. If the item is not
	// found by key, then return with an "ErrItemNotFound" error.
	// If the expiration is "nil" remove the expiration from the element.
	Update(key, value interface{}, expiration time.Duration) error

	// Has is just checking the given key is exists in the current map or not.
	Has(key interface{}) bool

	// Get an item from the map by key. If the item is not found, then return
	// with an "ErrItemNotFound" error, otherwise the item is exists, but
	// already expired, then return an "ErrItemIsExpired" error.
	// But when everythig is ok, return with the requested item of course.
	Get(key interface{}) (interface{}, error)

	// Remove an item from the map by key. If the item is not found, then return
	// with an "ErrItemNotFound" error, otherwise the item is exists, but
	// already expired, then return an "ErrItemIsExpired" error.
	Remove(key interface{}) error
}

type ttlMap struct{}

// New is instantiate a TTLMap. Every new TTLMap is fully empty, so not
// contains items.
func New() TTLMap {
	return &ttlMap{}
}

func (m *ttlMap) Insert(key, value interface{}, expiration time.Duration) error {
	return nil
}

func (m *ttlMap) Update(key, value interface{}, expiration time.Duration) error {
	return nil
}

func (m *ttlMap) Has(key interface{}) bool {
	return false
}

func (m *ttlMap) Get(key interface{}) (interface{}, error) {
	return nil, nil
}

func (m *ttlMap) Remove(key interface{}) error {
	return nil
}
