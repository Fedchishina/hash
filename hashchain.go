package hash

import (
	"errors"
)

// nodeChain is a structure item of hash table with chain
type nodeChain[K, V Hashable] struct {
	key   K
	value V
	next  *nodeChain[K, V]
}

// HashTableChain - is a structure of hash table
// size - is a size of hash table
// items - items of hash table
// indexer - is an indexer which wil be used for getting index by key for hash table
// countElements - current count of elements in hash table
type HashTableChain[K, V Hashable] struct {
	size          uint
	items         []*nodeChain[K, V]
	countElements uint
}

// NewHashTableChain - function for creating empty hash table
func NewHashTableChain[K, V Hashable](size uint) *HashTableChain[K, V] {
	return &HashTableChain[K, V]{
		size:  size,
		items: make([]*nodeChain[K, V], size),
	}
}

// Insert - function for inserting item to hash table
func (ht *HashTableChain[K, V]) Insert(key K, value V) error {
	ht.countElements++
	if ht.countElements > ht.size {
		if err := ht.rebuild(); err != nil {
			return err
		}
	}

	index, err := ht.index(key)
	if err != nil {
		return err
	}

	n := &nodeChain[K, V]{
		key:   key,
		value: value,
		next:  ht.items[index],
	}
	ht.items[index] = n

	return nil
}

// Search - function for searching item in hash table by key. Function will return value of item by key
func (ht *HashTableChain[K, V]) Search(key K) (V, error) {
	var zeroValue V
	index, err := ht.index(key)
	if err != nil {
		return zeroValue, err
	}

	current := ht.items[index]
	for current != nil {
		if current.key == key {
			return current.value, nil
		}
		current = current.next
	}

	return zeroValue, errors.New("item not found")
}

// Delete - function for Deleting item by key in hash table
func (ht *HashTableChain[K, V]) Delete(key K) error {
	index, err := ht.index(key)
	if err != nil {
		return err
	}

	current := ht.items[index]
	if current == nil {
		return errors.New("item not found")
	}

	if current.key == key {
		ht.items[index] = current.next
		ht.countElements--

		return nil
	}

	for current.next != nil && !(current.next.key == key) {
		current = current.next
	}

	if current.next != nil && current.next.key == key {
		current.next = current.next.next
		ht.countElements--

		return nil
	}

	return errors.New("item not found")
}

// rebuild - internal function for rebuilding hash table
func (ht *HashTableChain[K, V]) rebuild() error {
	newTable := NewHashTableChain[K, V](ht.size * 2)

	for _, item := range ht.items {
		current := item
		for current != nil {
			err := newTable.Insert(current.key, current.value)
			if err != nil {
				return err
			}
			current = current.next
		}
	}

	ht.items = newTable.items
	ht.size = ht.size * 2

	return nil
}

// index - get index (by key) where we will insert item to hash table
func (ht *HashTableChain[K, V]) index(key K) (uint, error) {
	if ht.size == 0 {
		return 0, errors.New("size cannot be zero")
	}

	hash, err := calculateHash(key)
	if err != nil {
		return 0, err
	}

	return uint(hash % uint32(ht.size)), nil
}
