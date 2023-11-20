hash
=======================
Library for work with hash tables.

You can create a hash table and use a list of functions to work with it.

We have an interface which you can use to get index for inserting item to hash table.
```
type Indexer interface {
	Hash(key int) uint
	Index(key int) (uint, error)
}
```

In this library was made the ModuloIndexer realisation of this interface. You can make your realisation or use ModuloIndexer when you create hash table.

## Hash functions
- [Empty hash table creation example](#empty-hash-table-creation-example)
- [Insert key into hash table](#insert-key-into-hash-table)
- [Search element by key](#search-element-by-key)
- [Delete element by key](#delete-element-by-key)

### Empty hash table creation example
```
ht := hash.NewHashTable(3, hash.NewModuloIndexer(3)) // simple hash table
htChain := hash.NewHashTableChain(3, hash.NewModuloIndexer(3)) // create hash table with chain realisation
```

### Insert key into hash table
```
ht := hash.NewHashTable(3, hash.NewModuloIndexer(3)) // simple hash table
ht.Insert(1, "value 1")
ht.Insert(2, "value 2")
```

### Search element by key
```
ht := hash.NewHashTable(3, hash.NewModuloIndexer(3)) // simple hash table
ht.Insert(1, "value 1")
ht.Insert(2, "value 2")

result, err := ht.Search(1) // return "value 1"
resultNil, err := ht.Search(3) // return "not found" error
```

### Delete element by key
```
ht := hash.NewHashTable(3, hash.NewModuloIndexer(3)) // simple hash table
ht.Insert(1, "value 1")
ht.Insert(2, "value 2")

err := ht.Delete(1) // deleting without error
err := ht.Delete(3) // return "not found" error
```
