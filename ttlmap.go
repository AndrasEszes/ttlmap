package ttlmap

import (
	"errors"
	"sync"
	"time"
)

var (
	// ErrKeyAlreadyExists happens, when try to add an item with a key which is
	// already exists in the map.
	ErrKeyAlreadyExists = errors.New("key is already exists")

	// ErrNilKeyIsNotAcceptable happens, when someone try to operate with nil
	// valued key.
	ErrNilKeyIsNotAcceptable = errors.New("nil key is not acceptable")

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
	// "ErrKeyExists" error, and when key is nil then return with
	// "ErrNilKeyIsNotAcceptable" error.
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

type ttlMap struct {
	mutex sync.RWMutex
	items map[interface{}]interface{}
}

// New is instantiate a TTLMap. Every new TTLMap is fully empty, so not
// contains items.
func New() TTLMap {
	return &ttlMap{items: make(map[interface{}]interface{})}
}

func (m *ttlMap) Insert(key, value interface{}, expiration time.Duration) error {
	if isNil(key) {
		return ErrNilKeyIsNotAcceptable
	}

	m.mutex.RLock()
	if _, exists := m.items[key]; exists {
		m.mutex.RUnlock()
		return ErrKeyAlreadyExists
	}
	m.mutex.RUnlock()

	m.mutex.Lock()
	m.items[key] = value
	m.mutex.Unlock()

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

func isNil(v interface{}) bool {
	return v == nil
}
