package hashtable

import (
	"errors"
	"hash/fnv"
	"reflect"
)

// Hashable is an interface for types that can be used as keys in the hash table
type Hashable interface {
	~int | ~string
}

// calculateHash calculates the hash value for the given key
func calculateHash[K Hashable](key K) (uint32, error) {
	keyBytes, err := keyToBytes(key)
	if err != nil {
		return 0, err
	}

	hasher := fnv.New32a()
	_, err = hasher.Write(keyBytes)
	if err != nil {
		return 0, err
	}

	return hasher.Sum32(), nil
}

// keyToBytes converts a key to bytes using reflection
func keyToBytes(key any) ([]byte, error) {
	value := reflect.ValueOf(key)

	if value.Kind() == reflect.String {
		return []byte(value.String()), nil
	}

	if isInteger(value.Kind()) {
		intValue := value.Interface().(int)
		return intToBytes(intValue), nil
	}

	return nil, errors.New("unsupported key type")
}

// isInteger checks if the given kind is an integer type
func isInteger(kind reflect.Kind) bool {
	return kind >= reflect.Int && kind <= reflect.Int64
}

// intToBytes converts an integer to bytes
func intToBytes(i int) []byte {
	result := make([]byte, 4)
	result[0] = byte(i)
	result[1] = byte(i >> 8)
	result[2] = byte(i >> 16)
	result[3] = byte(i >> 24)

	return result
}
