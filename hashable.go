package hash

// Hashable is an interface for types that can be used as keys in the hash table
type Hashable interface {
	~int | ~string
	Hash() uint
}

type IntKey int

func (i IntKey) Hash() uint {
	return uint(i)
}

type StringKey string

func (k StringKey) Hash() uint {
	hash := 0
	for _, char := range k {
		hash += int(char)
	}

	return uint(hash)
}
