package hashtable

import (
	"errors"
)

// node is a structure item of hash table with chain
type node[K, V Hashable] struct {
	key   K
	value V
	next  *node[K, V]
}

// HashTable is a structure of hash table
// size - is a size of hash table
// items - items of hash table
// indexer - is an indexer which wil be used for getting index by key for hash table
// countElements - current count of elements in hash table
type HashTable[K, V Hashable] struct {
	size          uint
	items         []*node[K, V]
	countElements uint
}

// NewHashTable - function for creating empty hash table
func NewHashTable[K, V Hashable](s uint) *HashTable[K, V] {
	return &HashTable[K, V]{
		size:  s,
		items: make([]*node[K, V], s),
	}
}

// Insert - function for inserting item to hash table using linear probing
func (h *HashTable[K, V]) Insert(key K, value V) error {
	h.countElements++
	if h.countElements > h.size {
		if err := h.rebuild(); err != nil {
			return err
		}
	}

	n := &node[K, V]{
		key:   key,
		value: value,
	}

	index, err := h.index(key)
	if err != nil {
		return err
	}

	v := h.items[index]
	if v == nil {
		h.items[index] = n

		return nil
	}

	for i := 1; i < int(h.size); i++ {
		newIndex := (index + uint(i)) % h.size

		if h.items[newIndex] == nil {
			h.items[newIndex] = n
			return nil
		}
	}

	return nil
}

// Search - function for searching item in hash table by key. Function will return value of item by key
func (h *HashTable[K, V]) Search(key K) (V, error) {
	var zeroValue V
	index, err := h.index(key)
	if err != nil {
		return zeroValue, err
	}

	if h.items[index] == nil {
		return zeroValue, errors.New("item not found")
	}

	return h.items[index].value, nil
}

// Delete - function for Deleting item by key in hash table
func (h *HashTable[K, V]) Delete(key K) error {
	index, err := h.index(key)
	if err != nil {
		return err
	}

	if h.items[index] == nil {
		return errors.New("item not found")
	}

	h.items[index] = nil
	h.countElements--

	return nil
}

// rebuild - internal function for rebuilding hash table
func (h *HashTable[K, V]) rebuild() error {
	newTable := NewHashTable[K, V](h.size * 2)

	for _, item := range h.items {
		if err := newTable.Insert(item.key, item.value); err != nil {
			return err
		}
	}

	h.items = newTable.items
	h.size = h.size * 2

	return nil
}

// Index - get index (by key) where we will insert item to hash table
func (h *HashTable[K, V]) index(key K) (uint, error) {
	if h.size == 0 {
		return 0, errors.New("size cannot be zero")
	}

	hash, err := calculateHash(key)
	if err != nil {
		return 0, err
	}

	return uint(hash % uint32(h.size)), nil
}
