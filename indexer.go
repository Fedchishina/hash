package hash

import "errors"

// Indexer is the interface which you can use to get index for inserting item to hash table
type Indexer interface {
	Hash(key int) uint
	Index(key int) (uint, error)
}

// ModuloIndexer is a realisation Indexer interface: using remainder of the division
type ModuloIndexer struct {
	size uint
}

// NewModuloIndexer - create ModuloIndexer
// size - is a size of hash table which will use this indexer
func NewModuloIndexer(size uint) *ModuloIndexer {
	return &ModuloIndexer{
		size: size,
	}
}

// Hash - get hash by key
func (di *ModuloIndexer) Hash(key int) uint {
	return uint(key)
}

// Index - get index (by key) where we will insert item to hash table
func (di *ModuloIndexer) Index(key int) (uint, error) {
	if di.size == 0 {
		return 0, errors.New("size cannot be zero")
	}

	return di.Hash(key) % di.size, nil
}
