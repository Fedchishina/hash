package hash

import (
	"errors"
)

// node is a structure item of hash table with chain
type node[T Hashable] struct {
	key   T
	value any
	next  *nodeChain[T]
}

// HashTable is a structure of hash table
// size - is a size of hash table
// items - items of hash table
// indexer - is an indexer which wil be used for getting index by key for hash table
// countElements - current count of elements in hash table
type HashTable[T Hashable] struct {
	size          uint
	items         map[uint]*node[T]
	countElements uint
}

// NewHashTable - function for creating empty hash table
func NewHashTable[T Hashable](s uint) *HashTable[T] {
	return &HashTable[T]{
		size:  s,
		items: make(map[uint]*node[T], s),
	}
}

// Insert - function for inserting item to hash table using linear probing
func (h *HashTable[T]) Insert(key T, value any) error {
	h.countElements++
	if h.countElements > h.size {
		if err := h.rebuild(); err != nil {
			return err
		}
	}

	n := &node[T]{
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
func (h *HashTable[T]) Search(key T) (any, error) {
	index, err := h.index(key)
	if err != nil {
		return nil, err
	}

	if h.items[index] == nil {
		return nil, errors.New("item not found")
	}

	return h.items[index].value, nil
}

// Delete - function for Deleting item by key in hash table
func (h *HashTable[T]) Delete(key T) error {
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
func (h *HashTable[T]) rebuild() error {
	newTable := NewHashTable[T](h.size * 2)

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
func (h *HashTable[T]) index(key T) (uint, error) {
	if h.size == 0 {
		return 0, errors.New("size cannot be zero")
	}

	return key.Hash() % h.size, nil
}
