package hash

import "errors"

// HashTable is a structure of hash table
// size - is a size of hash table
// items - items of hash table
// indexer - is an indexer which wil be used for getting index by key for hash table
// countElements - current count of elements in hash table
type HashTable struct {
	size          uint
	items         []interface{}
	indexer       Indexer
	countElements uint
}

// NewHashTable - function for creating empty hash table
func NewHashTable(s uint, indexer Indexer) *HashTable {
	return &HashTable{
		size:    s,
		items:   make([]interface{}, s),
		indexer: indexer,
	}
}

// Insert - function for inserting item to hash table
func (h *HashTable) Insert(key int, value interface{}) error {
	h.countElements++
	if h.countElements > h.size {
		if err := h.rebuild(); err != nil {
			return err
		}
	}

	index, err := h.indexer.Index(key)
	if err != nil {
		return err
	}

	h.items[index] = value

	return nil
}

// InsertLinearProbing - function for inserting item to hash table using linear probing
func (h *HashTable) InsertLinearProbing(key int, value interface{}) error {
	h.countElements++
	if h.countElements > h.size {
		if err := h.rebuild(); err != nil {
			return err
		}
	}

	index, err := h.indexer.Index(key)
	if err != nil {
		return err
	}

	v := h.items[index]
	if v == nil {
		h.items[index] = value

		return nil
	}

	for i := 1; i < int(h.size); i++ {
		newIndex := (int(index) + i) % int(h.size)

		if h.items[newIndex] == nil {
			h.items[newIndex] = value
			return nil
		}
	}

	return nil
}

// Search - function for searching item in hash table by key. Function will return value of item by key
func (h *HashTable) Search(key int) (interface{}, error) {
	index, err := h.indexer.Index(key)
	if err != nil {
		return nil, err
	}

	if h.items[index] == nil {
		return nil, errors.New("item not found")
	}

	return h.items[index], nil
}

// Delete - function for Deleting item by key in hash table
func (h *HashTable) Delete(key int) error {
	index, err := h.indexer.Index(key)
	if err != nil {
		return err
	}

	if h.items[index] == nil {
		return errors.New("item not found")
	}

	h.items[index] = nil

	return nil
}

// rebuild - internal function for rebuilding hash table
func (h *HashTable) rebuild() error {
	newTable := NewHashTable(h.size*2, h.indexer)

	for key, item := range h.items {
		if err := newTable.Insert(key, item); err != nil {
			return err
		}
	}

	h.items = newTable.items
	h.size = h.size * 2

	return nil
}
