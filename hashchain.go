package hash

import "errors"

// node is a structure item of hash table with chain
type node struct {
	key   int
	value interface{}
	next  *node
}

// HashTableChain - is a structure of hash table
// size - is a size of hash table
// items - items of hash table
// indexer - is an indexer which wil be used for getting index by key for hash table
// countElements - current count of elements in hash table
type HashTableChain struct {
	size          uint
	items         map[uint]*node
	indexer       Indexer
	countElements uint
}

// NewHashTableChain - function for creating empty hash table
func NewHashTableChain(size uint, indexer Indexer) *HashTableChain {
	return &HashTableChain{
		size:    size,
		items:   make(map[uint]*node, size),
		indexer: indexer,
	}
}

// Insert - function for inserting item to hash table
func (ht *HashTableChain) Insert(key int, value interface{}) error {
	ht.countElements++
	if ht.countElements > ht.size {
		if err := ht.rebuild(); err != nil {
			return err
		}
	}

	index, err := ht.indexer.Index(key)
	if err != nil {
		return err
	}

	n := &node{
		key:   key,
		value: value,
		next:  ht.items[index],
	}
	ht.items[index] = n

	return nil
}

// Search - function for searching item in hash table by key. Function will return value of item by key
func (ht *HashTableChain) Search(key int) (interface{}, error) {
	index, err := ht.indexer.Index(key)
	if err != nil {
		return nil, err
	}

	current := ht.items[index]
	for current != nil {
		if current.key == key {
			return current.value, nil
		}
		current = current.next
	}

	return nil, errors.New("item not found")
}

// Delete - function for Deleting item by key in hash table
func (ht *HashTableChain) Delete(key int) error {
	index, err := ht.indexer.Index(key)
	if err != nil {
		return err
	}

	current := ht.items[index]
	if current == nil {
		return errors.New("item not found")
	}

	if current.key == key {
		ht.items[index] = current.next
		return nil
	}

	for current.next != nil && current.next.key != key {
		current = current.next
	}

	if current.next != nil && current.next.key == key {
		current.next = current.next.next
		return nil
	}

	return errors.New("item not found")
}

// rebuild - internal function for rebuilding hash table
func (ht *HashTableChain) rebuild() error {
	newTable := NewHashTableChain(ht.size*2, ht.indexer)

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
